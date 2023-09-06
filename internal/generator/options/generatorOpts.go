package options

import (
	"gen/config"
	"os"
)

type Options struct {
	CreateSubFolder  bool
	PermissionFolder os.FileMode
	PermissionFile   os.FileMode
}

func defaultGeneratorOption() Options {
	return Options{
		CreateSubFolder:  false,
		PermissionFolder: os.ModePerm,
		PermissionFile:   os.ModePerm,
	}
}

func NewOptions(conf config.Config) Options {
	return defaultGeneratorOption()
}
