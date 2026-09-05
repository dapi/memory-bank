#!/usr/bin/env ruby
# frozen_string_literal: true

# Поддерживает project-local `memory-bank/` как проекцию payload из
# `template/memory-bank/`.
#
# Модель. Этот репозиторий — источник шаблона, а не его downstream-установка.
# Копировать payload в `memory-bank/` незачем: generic-документы представлены
# симлинками, поэтому инстанс всегда равен текущему payload и синхронизировать
# нечего. Реальными файлами остаются только те документы, которые проект
# переопределяет или которых в шаблоне нет: `features/`, `research/`, `adr/`,
# project-specific части `product/`, `domain/`, `ops/` и корневой индекс.
#
# Отсюда следует правило переопределения: чтобы адаптировать generic-документ
# под этот проект, замените симлинк обычным файлом. Обратно — удалите файл и
# прогоните этот скрипт.
#
# Использование:
#   ruby tools/refresh-memory-bank-projection.rb            # план
#   ruby tools/refresh-memory-bank-projection.rb --apply    # применить
#
# Скрипт не трогает реальные файлы: он только создаёт недостающие симлинки,
# чинит неверные цели и убирает те, что указывают на исчезнувший payload.

require 'fileutils'
require 'pathname'

class ProjectionError < StandardError; end

REPO_ROOT = Pathname.new(__dir__).parent.freeze
INSTANCE = (REPO_ROOT / 'memory-bank').freeze
PAYLOAD = (REPO_ROOT / 'template/memory-bank').freeze

# Целевой путь симлинка задаётся относительно каталога, в котором он лежит.
def link_target(relative)
  '../' * relative.each_filename.count + "template/memory-bank/#{relative}"
end

def payload_files
  PAYLOAD.glob('**/*').select(&:file?).map { |p| p.relative_path_from(PAYLOAD) }.sort
end

def instance_symlinks
  INSTANCE.glob('**/*').select(&:symlink?).map { |p| p.relative_path_from(INSTANCE) }.sort
end

def plan
  actions = []

  payload_files.each do |relative|
    installed = INSTANCE / relative
    expected = link_target(relative)

    if installed.symlink?
      actions << [:relink, relative, expected] unless File.readlink(installed) == expected
    elsif installed.exist?
      # Реальный файл — осознанное переопределение проекта; не трогаем.
      actions << [:override, relative, nil] if FileUtils.identical?(installed, PAYLOAD / relative)
    else
      actions << [:link, relative, expected]
    end
  end

  instance_symlinks.each do |relative|
    actions << [:prune, relative, nil] unless (PAYLOAD / relative).file?
  end

  actions
end

def describe(action, relative)
  case action
  when :link then "создать симлинк      memory-bank/#{relative}"
  when :relink then "исправить цель       memory-bank/#{relative}"
  when :prune then "удалить симлинк      memory-bank/#{relative} (в payload файла нет)"
  when :override then "переопределение      memory-bank/#{relative} (файл совпал с payload — можно вернуть в проекцию)"
  end
end

def apply!(actions)
  actions.each do |action, relative, expected|
    path = INSTANCE / relative
    case action
    when :link, :relink
      FileUtils.mkdir_p(path.dirname)
      path.delete if path.symlink? || path.exist?
      File.symlink(expected, path)
    when :prune
      path.delete
    end
  end
end

def main
  apply = ARGV.include?('--apply')
  unknown = ARGV - ['--apply']
  raise ProjectionError, "неизвестные аргументы: #{unknown.join(' ')}" unless unknown.empty?
  raise ProjectionError, "не найден #{PAYLOAD}" unless PAYLOAD.directory?
  raise ProjectionError, "не найден #{INSTANCE}" unless INSTANCE.directory?

  actions = plan
  changes = actions.reject { |action, _, _| action == :override }

  actions.each { |action, relative, _| puts describe(action, relative) }
  puts 'Проекция актуальна.' if actions.empty?

  return if changes.empty?

  unless apply
    puts "\nЭто план. Повторите с --apply, чтобы применить."
    return
  end

  apply!(changes)
  puts "\nПрименено изменений: #{changes.length}."
end

begin
  main
rescue ProjectionError => e
  warn "refresh-memory-bank-projection: #{e.message}"
  exit 1
end
