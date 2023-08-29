package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Generator",
		Version: "v0.1",
		Commands: []*cli.Command{
			{
				Name:  "gen",
				Usage: "generate code",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "Opts",
						Aliases: []string{"rdo"},
						Value:   "./",
						Usage:   "Load Options for reader and generator",
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
