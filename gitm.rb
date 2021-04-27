#! /usr/bin/ruby

require_relative './core'

default_branch = ENV["DEFAULT_BRANCH"] || 'master'
branch = `git rev-parse --abbrev-ref HEAD`.strip

if branch == default_branch
    run_cmd "git pull origin main"
else 
    run_cmds(
        "git co #{default_branch}",
        "git pull origin #{default_branch}",
        "git co #{branch}",
        "git pull origin #{default_branch}"
    )
end