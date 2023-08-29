package reader

import (
	"gen/enums"
	"gen/internal/generator"
	"gen/internal/reader/options"
	"gen/internal/tree"
	"gen/internal/tree/rawContent"
	"gen/ulti"
	"os"
	"strings"
)

type Reader struct {
	Opts options.Options
}

func NewReader(opts options.Options) Reader {
	return Reader{opts}
}

func (r Reader) ParseTreeGenerator(destination string) generator.Generator {
	//TODO implement me
	panic("implement me")
}

func (r Reader) ParseStructGenerator() generator.Generator {
	//TODO implement me
	panic("implement me")
}

func (r Reader) ParseTree(destination string) (*tree.Leaf, error) {
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
		if ulti.InArray(r.Opts.IgnoreDirectory, leafNode.Name) {
			return nil, nil
		}

		if !r.Opts.IncludeHiddenDir {
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

			subLeafNode, err := r.ParseTree(pathSubLeaf)
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
	var leafContent []byte

	_, err = leaf.Read(leafContent)
	if err != nil {
		return nil, err
	}

	leafNode.Content = rawContent.NewRawContent(leafContent)

	return &leafNode, nil
}
