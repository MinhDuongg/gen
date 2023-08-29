package options

import "os"

type Options struct {
	CreateSubFolder  bool
	PermissionFolder os.FileMode
	PermissionFile   os.FileMode
}

func DefaultGeneratorOption() Options {
	return Options{
		CreateSubFolder:  false,
		PermissionFolder: os.ModeDir,
		PermissionFile:   os.ModePerm,
	}
}
