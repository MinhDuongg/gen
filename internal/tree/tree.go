package tree

import "gen/enums"

type Leaf struct {
	Name     string         `json:"name"`
	Type     enums.FileType `json:"type"`
	SubLeafs []Leaf         `json:"subLeafs"`
	Content  ContentI       `json:"-"`
}
