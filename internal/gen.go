package internal

import (
	"context"
	"gen/config"
	"gen/enums"
	"gen/internal/generator/treeGenerator"
	"gen/internal/reader/reader"
	tree "gen/internal/tree"
	"log"
)

func Init(ctx context.Context, mode int64, source, destination string) error {
	cfg, err := config.GetCfg()

	if err != nil {
		log.Fatal(err)
	}

	reader := reader.NewReader(cfg)
	var tree tree.Leaf

	if enums.OperationMode(mode) == enums.Template {
		tree = reader.ParseTree(source)
	}

	generator := treeGenerator.NewTreeGenerator(tree, cfg)

	err = generator.Generate(ctx, destination)
	if err != nil {
		return err
	}

	return nil
}
