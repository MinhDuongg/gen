package tree

import "gen/enums"

type Leaf struct {
	Name     string
	Type     enums.FileType
	SubLeafs []Leaf
	Content  ContentI
}
