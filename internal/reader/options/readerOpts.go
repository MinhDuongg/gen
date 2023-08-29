package options

type Options struct {
	SecretKey        string
	GetFileContent   bool
	SimpleCopy       bool
	RepoDirectory    string
	IncludeHiddenDir bool
	IgnoreDirectory  []string
	IgnoreFile       []string
}
