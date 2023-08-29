package enums

type FileType int64

const (
	Directory FileType = iota
	GoFile
	Text
)
