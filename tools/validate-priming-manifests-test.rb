# frozen_string_literal: true

require "fileutils"
require "minitest/autorun"
require "pathname"
require "tmpdir"
require "yaml"

require_relative "validate-priming-manifests"

class PrimingManifestValidatorTest < Minitest::Test
  def test_rejects_absolute_remainder_that_points_outside_scope
    Dir.mktmpdir do |directory|
      root = Pathname(directory)
      scope_root = root.join("memory-bank")
      manifest_directory = scope_root.join("flows/priming")
      FileUtils.mkdir_p(manifest_directory)

      outside_file = root.join("outside.md")
      outside_file.write("outside scope")
      manifest_directory.join("sample.yaml").write(
        {
          "version" => 1,
          "process" => "sample",
          "stages" => {
            "entry" => ["memory-bank/#{outside_file}"]
          }
        }.to_yaml
      )

      result = nil
      _stdout, stderr = capture_io do
        result = PrimingManifestValidator.new(scope_root).run
      end

      refute result
      assert_includes stderr, "must remain relative under memory-bank/"
    end
  end
end
