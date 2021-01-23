require_relative './core'

commit_msg = ARGV.join(' ').strip

err("Empty commit message, exiting") if commit_msg.empty?

branch_name = `git rev-parse --abbrev-ref HEAD`.strip

run_cmds(
	"git status",
	"git add -A",
	"git commit -m '#{commit_msg}'",
	"git push origin #{branch_name}"
)
