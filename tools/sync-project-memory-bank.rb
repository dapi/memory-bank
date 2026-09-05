#!/usr/bin/env ruby
# frozen_string_literal: true

# Синхронизирует project-local `memory-bank/` с payload из `template/memory-bank/`.
#
# Зачем нужен отдельный инструмент. Этот репозиторий одновременно является
# upstream-шаблоном и его downstream-потребителем. Корневые ассеты
# (`WORKFLOW.md`, `.codex/agents`, `.start-issue/prompt.md`,
# `bootstrap-symphony.sh`, `run-symphony.sh`) — симлинки в `template/`, чтобы
# рабочие инструменты репозитория всегда совпадали с payload. По контракту из
# `docs/ownership.md` симлинк в любом компоненте пути считается unsafe path,
# поэтому `memory-bank-cli pull` в корне репозитория останавливается и ничего
# не читает.
#
# Скрипт выполняет pull в изолированной песочнице, где корневых ассетов просто
# нет, и переносит обратно только поддерево `memory-bank/**` вместе с его lock.
# Dual-role layout и симлинки при этом не меняются.
#
# Это обходной путь с ограниченным сроком жизни: если CLI научится ограничивать
# область синхронизации (dapi/memory-bank-cli#60), песочница и правка lock здесь
# станут не нужны.
#
# Использование:
#   ruby tools/sync-project-memory-bank.rb            # план, без изменений
#   ruby tools/sync-project-memory-bank.rb --apply    # применить план
#
# После применения прогоните проверки из `AGENTS.md`.

require 'fileutils'
require 'json'
require 'tmpdir'

class SyncError < StandardError; end

def run(*command, capture: false)
  if capture
    output = IO.popen(command, err: [:child, :out], &:read)
    [output, $?.success?]
  else
    system(*command) || raise(SyncError, "команда завершилась с ошибкой: #{command.join(' ')}")
  end
rescue Errno::ENOENT
  raise SyncError, "команда не найдена: #{command.first}"
end

def git(*args)
  output, ok = run('git', *args, capture: true)
  raise SyncError, "git #{args.join(' ')} завершился с ошибкой:\n#{output}" unless ok

  output.strip
end

def repo_root
  @repo_root ||= git('rev-parse', '--show-toplevel')
end

def ensure_preconditions
  lock = File.join(repo_root, 'memory-bank', '.lock')
  raise SyncError, "не найден #{lock}: это не dual-role checkout" unless File.file?(lock)

  template = File.join(repo_root, 'template', 'memory-bank')
  raise SyncError, "не найден #{template}: нечем синхронизировать" unless File.directory?(template)

  dirty = git('status', '--porcelain', '--', 'template')
  return if dirty.empty?

  raise SyncError, <<~MESSAGE
    в `template/` есть незакоммиченные изменения:
    #{dirty}
    Источником служит закреплённый commit, поэтому сначала закоммитьте payload.
  MESSAGE
end

# Предупреждает, если source_ref не достижим из origin/main.
#
# Lock закрепляет источник по commit SHA. Если синхронизироваться с ветки,
# которую потом смержат squash-ом, записанный SHA исчезнет из истории main и
# следующий pull не сможет сверить базу. Обычный порядок: сначала мержим
# изменение шаблона, потом синхронизируем инстанс.
def warn_if_unmerged(source_ref)
  _, ref_exists = run('git', 'rev-parse', '--verify', '--quiet', 'origin/main', capture: true)
  unless ref_exists
    warn "Внимание: origin/main недоступен, достижимость #{source_ref[0, 7]} не проверена.\n\n"
    return
  end

  _, ok = run('git', 'merge-base', '--is-ancestor', source_ref, 'origin/main', capture: true)
  return if ok

  warn <<~MESSAGE
    Внимание: #{source_ref[0, 7]} не достижим из origin/main.
    Lock закрепит commit, которого может не оказаться в истории main после
    squash-мержа. Надёжнее синхронизировать инстанс уже после мержа шаблона.

  MESSAGE
end

# Оставляет в lock только пути `memory-bank/**`.
#
# Корневые managed-пути в этом репозитории являются симлинками, CLI ими не
# управляет, и запись их digest была бы ложной: следующий pull счёл бы payload
# чистым. До появления этого инструмента lock тоже содержал только поддерево.
def strip_non_subtree_entries(lock_path)
  lock = JSON.parse(File.read(lock_path))
  files = lock['files'] || {}
  removed = files.keys.reject { |path| path.start_with?('memory-bank/') }
  removed.each { |path| files.delete(path) }
  lock['files'] = files.sort.to_h
  File.write(lock_path, "#{JSON.pretty_generate(lock)}\n")
  removed
end

# Подменяет `memory-bank/` результатом из песочницы.
#
# Транзакционные гарантии CLI заканчиваются на песочнице, поэтому перенос
# обратно делаем сами: сначала готовим полностью собранное дерево рядом с
# целью, затем меняем каталоги местами через rename. Прямая пара rm_rf + cp_r
# оставила бы инстанс отсутствующим или недописанным, если копирование упадёт
# или его прервут, — а песочница к тому моменту уже удалена.
def swap_in_result(result)
  target = File.join(repo_root, 'memory-bank')
  staging_root = File.join(repo_root, 'tmp')
  FileUtils.mkdir_p(staging_root)
  staging = File.join(staging_root, 'memory-bank.incoming')
  previous = File.join(staging_root, 'memory-bank.previous')
  [staging, previous].each { |path| FileUtils.rm_rf(path) }

  FileUtils.cp_r(result, staging)
  removed = strip_non_subtree_entries(File.join(staging, '.lock'))

  File.rename(target, previous)
  begin
    File.rename(staging, target)
  rescue StandardError
    File.rename(previous, target)
    raise
  end

  FileUtils.rm_rf(previous)
  removed
end

def pull_command(cli, sandbox, source, source_ref)
  [
    cli, 'pull',
    '--repo-root', sandbox,
    '--source', source,
    '--template-version', source_ref[0, 7],
    '--source-ref', source_ref
  ]
end

def main
  apply = ARGV.include?('--apply')
  unknown = ARGV - ['--apply']
  raise SyncError, "неизвестные аргументы: #{unknown.join(' ')}" unless unknown.empty?

  cli = ENV.fetch('MEMORY_BANK_CLI', 'memory-bank-cli')
  ensure_preconditions
  source_ref = git('rev-parse', 'HEAD')
  warn_if_unmerged(source_ref)

  Dir.mktmpdir('memory-bank-sync') do |tmp|
    source = File.join(tmp, 'source')
    sandbox = File.join(tmp, 'sandbox')

    git('worktree', 'add', '--detach', '--quiet', source, source_ref)
    begin
      FileUtils.mkdir_p(sandbox)
      FileUtils.cp_r(File.join(repo_root, 'memory-bank'), sandbox)

      command = pull_command(cli, sandbox, source, source_ref)
      plan, planned = run(*command, '--dry-run', capture: true)

      # В песочнице нет корневых ассетов, поэтому CLI планирует создать их
      # заново. Это артефакт изоляции: обратно переносится только поддерево.
      subtree = plan.lines.grep(/\tmemory-bank\//)
      conflicts = subtree.grep(/^conflict\t/)

      # Конфликты — ожидаемый исход с ненулевым кодом возврата; всё остальное
      # ненулевое означает сбой самого CLI и не должно выглядеть как «всё в
      # порядке».
      raise SyncError, "pull --dry-run завершился с ошибкой:\n#{plan}" if !planned && conflicts.empty?

      puts subtree.empty? ? 'memory-bank/ уже соответствует payload.' : subtree.join
      puts 'Корневые пути в плане пропущены: ими управляют симлинки в template/.'

      unless conflicts.empty?
        warn <<~MESSAGE

          Конфликты внутри memory-bank/ требуют явного решения человеком:
          #{conflicts.join}
          Примите incoming payload или осознанно мигрируйте ownership в lock,
          затем повторите запуск. См. docs/ownership.md.
        MESSAGE
        exit 1
      end

      unless apply
        puts "\nЭто план. Повторите с --apply, чтобы применить."
        return
      end

      run(*command)

      removed = swap_in_result(File.join(sandbox, 'memory-bank'))
      puts "\nПрименено. Из lock убрано корневых записей: #{removed.length}."
      puts 'Проверьте результат командами из AGENTS.md.'
    ensure
      begin
        git('worktree', 'remove', '--force', source)
      rescue SyncError => e
        # Не подменяем исходную ошибку сбоем уборки.
        warn "не удалось убрать временный worktree #{source}: #{e.message}"
      end
    end
  end
end

begin
  main
rescue SyncError => e
  warn "sync-project-memory-bank: #{e.message}"
  exit 1
end
