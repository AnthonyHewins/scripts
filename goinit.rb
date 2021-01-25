#!/usr/bin/env ruby

require 'erb'
require_relative './core'

err("Not enough args, exiting") if ARGV.length < 1

pwd = Dir.pwd
repo_dir = pwd + '/' + ARGV[0]

run_cmds(
  "mkdir -p #{repo_dir}/cmd/cli",
  "git init #{repo_dir}",
)

def build_template(filename)
  script_path = File.expand_path File.dirname(__FILE__)
  buf = File.read "#{script_path}/templates/go/#{filename}"

  template = ERB.new(buf)
  script_name = File.basename repo_dir
  template.result_with_hash(script_name: script_name)
end

mainFile = build_template("main.go.erb")
makeFile = build_template("Makefile.erb")

fileWrite("#{repo_dir}/cmd/cli/main.go", mainFile)
fileWrite("#{repo_dir}/Makefile", makeFile)
