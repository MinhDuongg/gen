package treeGenerator

import (
	"context"
	"fmt"
	"gen/enums"
	"gen/internal/generator/options"
	"gen/internal/tree"
	"gen/ulti"
	"os"
)

type TreeGenerator struct {
	Tree tree.Leaf
	Opts options.Options
}

func NewTreeGenerator(tree tree.Leaf, opts ...options.Options) TreeGenerator {
	options := options.DefaultGeneratorOption()

	if len(opts) != 0 {
		options = opts[0]
	}

	return TreeGenerator{
		Tree: tree,
		Opts: options,
	}
}

func (t TreeGenerator) Generate(ctx context.Context, destination string) error {
	return t.CreateBoilerPlate(ctx, t.Tree, destination)
}

func (t TreeGenerator) CreateBoilerPlate(ctx context.Context, leaf tree.Leaf, destination string) error {
	var err error

	if leaf.Name == "" {
		if leaf.Type == enums.Directory {
			return fmt.Errorf("Directory should have a specific name")
		}

		leaf.Name, err = ulti.GenerateFileName()
		if err != nil {
			return err
		}
	}

	leafDestination, err := ulti.PathConcat(destination, leaf.Name)

	switch leaf.Type {
	case enums.Directory:
		{
			err = os.Mkdir(leafDestination, t.Opts.PermissionFolder)

			if err != nil {
				if err != os.ErrExist {
					return err
				}

				break
			}
		}
	case enums.GoFile:
		{
			var f *os.File
			f, err = os.Create(leafDestination)

			if err != nil {
				if err != os.ErrExist {
					return err
				}

				f, err = os.Open(leafDestination)
				if err != nil {
					return err
				}
			}

			if !leaf.Content.ContentAvailbility(ctx) {
				return nil
			}

			content, err := leaf.Content.ContentWriter(ctx)
			if err != nil {
				return err
			}

			_, err = f.Write(content)
			if err != nil {
				return err
			}

			err = os.Chmod(leafDestination, t.Opts.PermissionFile)
			if err != nil {
				return err
			}

			return nil
		}
	}

	for _, subLeaf := range leaf.SubLeafs {
		err = t.CreateBoilerPlate(ctx, subLeaf, destination)
		if err != nil {
			return err
		}
	}

	return nil
}
