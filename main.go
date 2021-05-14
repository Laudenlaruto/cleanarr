package main

import (
	"log"
	"net/http"
	"os"

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
							sonarStatus(client, url, key)
							return nil
						},
					},
					{
						Name:  "space",
						Usage: "space what is using space in Sonarr",
						Action: func(c *cli.Context) error {
							client := &http.Client{}
							getSeries(client,url,key)
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