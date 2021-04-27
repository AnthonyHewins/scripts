require_relative './core'

n = ARGV.length

error "Not enough args" if n < 1

bug = 'bugfix'
feature = 'feature'
hotfix = 'hotfix'

modes = {
	"b" => bug,
	"bug" => bug,
	"f" => feature,
	"feature" => feature,
	'hotfix' => hotfix,
	'h' => hotfix,
}

mode = modes[ARGV[0]]
error "Invalid mode #{ARGV[0]}. Use bug, feature, hotfix" if mode.nil?

ticketName = ARGV[1..-1].join(' ').strip

default_branch = ENV['DEFAULT_BRANCH'] || 'master'
run_cmds(
	"git co #{default_branch}",
	"git checkout -b #{mode.downcase}/#{ticketName.gsub(/\s+/, '-')}"
)
