package main

import (
	"fmt"
	"gen/config"
	"gen/internal"
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
					&cli.Int64Flag{
						Name:     "mode",
						Aliases:  []string{"m"},
						Value:    0,
						Required: true,
						Usage:    "choose the mode for operation: 0 for template mode",
						Action: func(context *cli.Context, s int64) error {
							if s != 0 {
								return fmt.Errorf("Only support parsing from template mode")
							}
							return nil
						},
					},
					&cli.StringFlag{
						Name:    "destination",
						Aliases: []string{"d"},
						Value:   "./",
						Usage:   "path of the clone project",
					},
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Value:   "./.env",
						Usage:   "load config file",
					},
					&cli.StringFlag{
						Name:     "source",
						Category: "template",
						Aliases:  []string{"t"},
						Value:    "./template",
						Usage:    "path of the source project",
					},
				},
				Action: GenerateCodeFromTemplate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func GenerateCodeFromTemplate(cCtx *cli.Context) error {
	configFile := cCtx.String("config")
	err := config.NewConfig(configFile)

	if err != nil {
		log.Println("error: reading config file")
		return err
	}

	mode := cCtx.Int64("mode")
	source := cCtx.String("source")
	destination := cCtx.String("destination")

	return internal.Init(cCtx.Context, mode, source, destination)
}
