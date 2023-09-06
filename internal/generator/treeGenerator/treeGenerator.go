package treeGenerator

import (
	"context"
	"fmt"
	"gen/config"
	"gen/enums"
	"gen/internal/generator/options"
	"gen/internal/tree"
	"gen/ulti"
	"log"
	"os"
)

type TreeGenerator struct {
	Tree        tree.Leaf
	destination string
	opts        options.Options
}

func NewTreeGenerator(tree tree.Leaf, config config.Config) TreeGenerator {
	opts := options.NewOptions(config)

	return TreeGenerator{
		Tree: tree,
		opts: opts,
	}
}

func (t TreeGenerator) Generate(ctx context.Context, destination string) error {
	leafDestination, err := t.createBoilerPlate(ctx, t.Tree, destination)
	t.destination = leafDestination
	return err
}

func (t TreeGenerator) CleanUp(ctx context.Context) error {
	if t.destination == "" {
		return nil
	}

	err := os.RemoveAll(t.destination)
	if err != nil {
		return err
	}

	return nil
}

func (t TreeGenerator) createBoilerPlate(ctx context.Context, leaf tree.Leaf, destination string) (string, error) {
	var err error

	if leaf.Name == "" {
		if leaf.Type == enums.Directory {
			return "", fmt.Errorf("Directory should have a specific name")
		}

		leaf.Name, err = ulti.GenerateFileName()
		if err != nil {
			return "", err
		}
	}

	leafDestination, err := ulti.PathConcat(destination, leaf.Name)
	if err != nil {
		return "", err
	}

	switch leaf.Type {
	case enums.Directory:
		{
			err = os.Mkdir(leafDestination, t.opts.PermissionFolder)

			if err != nil {
				if err != os.ErrExist {
					return leafDestination, err
				}

				break
			}
		}
	default:
		{
			if !leaf.Content.ContentAvailbility(ctx) {
				return leafDestination, nil
			}
			content, err := leaf.Content.ContentWriter(ctx)
			if err != nil {
				return leafDestination, err
			}

			err = os.WriteFile(leafDestination, content, t.opts.PermissionFile)
			if err != nil {
				log.Println(err)
				return leafDestination, err
			}

			return leafDestination, nil
		}
	}

	for _, subLeaf := range leaf.SubLeafs {
		_, err = t.createBoilerPlate(ctx, subLeaf, leafDestination)
		if err != nil {
			return leafDestination, err
		}
	}

	return leafDestination, nil
}
