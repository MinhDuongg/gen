package options

import "gen/config"

type Options struct {
	SecretKey        string
	GetFileContent   bool
	SimpleCopy       bool
	RepoDirectory    string
	IncludeHiddenDir bool
	IgnoreDirectory  []string
	IgnoreFile       []string
}

func NewOptions(conf config.Config) Options {
	return Options{}
}
