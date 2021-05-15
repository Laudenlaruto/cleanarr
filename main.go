package main

import (
	"cleanarr/sonarr"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	var url = "http://192.168.1.19:8989/api/"
	var key = "e695509add214a7c8f7f20d829ea8d44"
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
							client := &http.Client{}
							sonarr.Status(client, url, key)
							return nil
						},
					},
					{
						Name:  "space",
						Usage: "space what is using space in Sonarr",
						Action: func(c *cli.Context) error {
							client := &http.Client{}
							sonarr.UsedSpace(client, url, key)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list all episodes that have beed added X days ago",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "days",
								Usage: "How long age has the episode been added?",
							},
						},
						Action: func(c *cli.Context) error {
							days, err := strconv.Atoi(c.String("days"))
							if err != nil {
								log.Fatal("Parameter --days should be an number")
							}
							client := &http.Client{}
							sonarr.ListEpisodeWithFilter(client, url, key, days)
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
