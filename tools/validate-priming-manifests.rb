# frozen_string_literal: true

require "pathname"
require "yaml"

class PrimingManifestValidator
  TOP_LEVEL_KEYS = %w[process stages version].freeze
  IDENTIFIER = /\A[a-z][a-z0-9_]*\z/
  PLACEHOLDER = /<[A-Z][A-Z0-9-]*>/
  EXTERNAL_SOURCE = %r{\A[a-z][a-z0-9+.-]*://}i

  def initialize(scope_root)
    @scope_root = Pathname(scope_root).cleanpath
    @errors = []
    @process_owners = {}
    @validated = 0
  end

  def run
    unless @scope_root.directory?
      @errors << "#{@scope_root}: scope root does not exist"
      return finish
    end

    manifests = Dir[@scope_root.join("flows/priming/*.yaml").to_s].sort
    if manifests.empty?
      @errors << "#{@scope_root}/flows/priming: no manifests found"
    else
      manifests.each { |path| validate(path) }
    end

    process_template = @scope_root.join("flows/templates/process/priming.yaml")
    validate(process_template, template: true) if process_template.file?

    finish
  end

  private

  def validate(path, template: false)
    document = YAML.safe_load(
      File.read(path),
      permitted_classes: [],
      permitted_symbols: [],
      aliases: false,
      filename: path.to_s
    )
    @validated += 1

    unless document.is_a?(Hash)
      add_error(path, "document must be a mapping")
      return
    end

    keys = document.keys
    unless keys.all? { |key| key.is_a?(String) } && keys.sort == TOP_LEVEL_KEYS
      add_error(path, "expected only keys #{TOP_LEVEL_KEYS.join(', ')}")
    end
    add_error(path, "version must equal 1") unless document["version"] == 1

    process = document["process"]
    unless process.is_a?(String) && (template || process.match?(IDENTIFIER))
      add_error(path, "process must be a lowercase snake_case identifier")
    else
      register_process(path, process) unless template
    end

    stages = document["stages"]
    unless stages.is_a?(Hash) && !stages.empty?
      add_error(path, "stages must be a non-empty mapping")
      return
    end

    stages.each do |stage, inputs|
      unless stage.is_a?(String) && stage.match?(IDENTIFIER)
        add_error(path, "stage #{stage.inspect} must be a lowercase snake_case identifier")
      end
      validate_inputs(path, stage, inputs, template: template)
    end
  rescue Psych::Exception => error
    add_error(path, "invalid YAML: #{error.message.lines.first.strip}")
  end

  def validate_inputs(path, stage, inputs, template:)
    unless inputs.is_a?(Array) && !inputs.empty?
      add_error(path, "stage #{stage.inspect} must contain a non-empty input list")
      return
    end

    duplicates = inputs.group_by(&:itself).select { |_input, occurrences| occurrences.length > 1 }.keys
    add_error(path, "stage #{stage.inspect} contains duplicate inputs: #{duplicates.join(', ')}") unless duplicates.empty?

    inputs.each do |input|
      unless input.is_a?(String) && !input.empty?
        add_error(path, "stage #{stage.inspect} contains a non-string or empty input")
        next
      end

      validate_input(path, stage, input, template: template)
    end
  end

  def validate_input(path, stage, input, template:)
    return if input.match?(EXTERNAL_SOURCE)

    unless input.start_with?("memory-bank/")
      add_error(path, "stage #{stage.inspect}: #{input.inspect} must be repo-relative under memory-bank/")
      return
    end

    if input.match?(/TODO/i) || input.include?("**") || input.match?(/[?\[\]{}]/)
      add_error(path, "stage #{stage.inspect}: #{input.inspect} contains an unsupported placeholder or glob")
      return
    end

    relative = input.delete_prefix("memory-bank/")
    relative_path = Pathname(relative)
    if relative_path.absolute?
      add_error(path, "stage #{stage.inspect}: #{input.inspect} must remain relative under memory-bank/")
      return
    end

    if relative_path.each_filename.include?("..")
      add_error(path, "stage #{stage.inspect}: #{input.inspect} escapes the scope root")
      return
    end

    return if template

    placeholders = input.scan(/<[^>]*>/)
    invalid_placeholder = (input.include?("<") || input.include?(">")) &&
      (placeholders.empty? || input.gsub(PLACEHOLDER, "").match?(/[<>]/))
    if invalid_placeholder || !placeholders.all? { |placeholder| placeholder.match?(PLACEHOLDER) }
      add_error(path, "stage #{stage.inspect}: #{input.inspect} contains an invalid placeholder")
      return
    end

    if placeholders.any?
      prefix = relative.split("<", 2).first
      prefix_path = @scope_root.join(prefix)
      add_error(path, "stage #{stage.inspect}: fixed prefix for #{input.inspect} does not exist") unless prefix_path.directory?
      return
    end

    resolved = @scope_root.join(relative).to_s
    if input.include?("*")
      matches = Dir.glob(resolved).select { |match| File.file?(match) }
      add_error(path, "stage #{stage.inspect}: #{input.inspect} matches no files") if matches.empty?
    elsif !File.file?(resolved)
      add_error(path, "stage #{stage.inspect}: #{input.inspect} does not exist")
    end
  end

  def add_error(path, message)
    @errors << "#{path}: #{message}"
  end

  def register_process(path, process)
    owner = @process_owners[process]
    if owner
      add_error(path, "process #{process.inspect} is already owned by #{owner}")
    else
      @process_owners[process] = path
    end
  end

  def finish
    if @errors.empty?
      puts "Validated #{@validated} priming manifests."
      return true
    end

    warn @errors.join("\n")
    false
  end
end

if $PROGRAM_NAME == __FILE__
  scope_root = ARGV.fetch(0, "template/memory-bank")
  exit 1 unless PrimingManifestValidator.new(scope_root).run
end
