package treeGenerator

import (
	"context"
	"fmt"
	"gen/config"
	"gen/enums"
	"gen/internal/generator/options"
	"gen/internal/tree"
	"gen/ulti"
	"os"
)

type TreeGenerator struct {
	Tree tree.Leaf
	opts options.Options
}

func NewTreeGenerator(tree tree.Leaf, config config.Config) TreeGenerator {
	opts := options.NewOptions(config)

	return TreeGenerator{
		Tree: tree,
		opts: opts,
	}
}

func (t TreeGenerator) Generate(ctx context.Context, destination string) error {
	return t.createBoilerPlate(ctx, t.Tree, destination)
}

func (t TreeGenerator) createBoilerPlate(ctx context.Context, leaf tree.Leaf, destination string) error {
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
	if err != nil {
		return err
	}

	switch leaf.Type {
	case enums.Directory:
		{
			err = os.Mkdir(leafDestination, t.opts.PermissionFolder)

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

			err = os.Chmod(leafDestination, t.opts.PermissionFile)
			if err != nil {
				return err
			}

			return nil
		}
	}

	for _, subLeaf := range leaf.SubLeafs {
		err = t.createBoilerPlate(ctx, subLeaf, destination)
		if err != nil {
			return err
		}
	}

	return nil
}
