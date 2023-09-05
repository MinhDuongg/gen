package reader

import (
	"gen/internal/tree"
)

type Reader interface {
	ParseTree(source string) (tree.Leaf, error)
}
