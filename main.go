package cleanarr

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "sonarr",
				// Aliases:     []string{"t"},
				Usage: "options for interacting with Sonarr",
				Subcommands: []*cli.Command{
					{
						Name:  "status",
						Usage: "Get status of Sonarr",
						Action: func(c *cli.Context) error {
							sonarStatus()
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "List everything in Sonarr",
						Action: func(c *cli.Context) error {
							// fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
