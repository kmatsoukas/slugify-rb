# frozen_string_literal: true
require 'ffi'
require_relative "slugify/version"

module Slugify
  extend FFI::Library
  class Error < StandardError; end

  ffi_lib "#{File.dirname(__FILE__)}/../ext/slugify.so"

  attach_function :slugify, [:string], :string
end
