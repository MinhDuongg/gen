package reader

import (
	"gen/config"
	"gen/enums"
	"gen/internal/reader/options"
	"gen/internal/tree"
	"gen/internal/tree/rawContent"
	"gen/ulti"
	"os"
	"strings"
)

type Reader struct {
	opts options.Options
}

func NewReader(config config.Config) Reader {
	opts := options.NewOptions(config)
	return Reader{opts}
}

func (r Reader) ParseTree(destination string) (tree.Leaf, error) {
	parsedTree, err := r.parseTemplateToTree(destination)
	if err != nil {
		return tree.Leaf{}, err
	}

	return *parsedTree, nil
}

func (r Reader) parseTemplateToTree(destination string) (*tree.Leaf, error) {
	var err error
	var leafNode tree.Leaf

	leaf, err := os.Open(destination)
	if err != nil {
		return nil, err
	}

	defer leaf.Close()

	leafStat, err := leaf.Stat()
	if err != nil {
		return nil, err
	}

	if leafStat.IsDir() {
		if ulti.InArray(r.opts.IgnoreDirectory, leafNode.Name) {
			return nil, nil
		}

		if !r.opts.IncludeHiddenDir {
			if strings.HasPrefix(leafNode.Name, ".") {
				return nil, nil
			}
		}

		leafNode.Name = leafStat.Name()
		leafNode.Type = enums.Directory

		subLeafs, err := leaf.Readdir(-1)
		if err != nil {
			return nil, err
		}

		for _, subLeaf := range subLeafs {
			pathSubLeaf, err := ulti.PathConcat(destination, subLeaf.Name())
			if err != nil {
				return nil, err
			}

			subLeafNode, err := r.parseTemplateToTree(pathSubLeaf)
			if err != nil {
				return nil, err
			}

			if subLeafNode == nil {
				continue
			}

			leafNode.SubLeafs = append(leafNode.SubLeafs, *subLeafNode)
		}

		return &leafNode, nil
	}

	leafNode.Name = leafStat.Name()
	leafNode.Type = enums.Text

	if strings.HasSuffix(leafNode.Name, "go") {
		leafNode.Type = enums.GoFile
	}

	var leafContent []byte

	leafContent, err = os.ReadFile(destination)
	if err != nil {
		return nil, err
	}

	leafNode.Content = rawContent.NewRawContent(leafContent)

	return &leafNode, nil
}
