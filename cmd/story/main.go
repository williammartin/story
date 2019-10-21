package main

import (
	"os"
	"story/api"
	"story/cmd"
	"story/console"

	"github.com/jessevdk/go-flags"
	"github.com/machinebox/graphql"
)

type Defaults struct{}

func main() {
	ui := console.Writer{
		Out: os.Stdout,
	}

	gqlClient := graphql.NewClient("https://api.storyscript.io/graphql")
	apiClient := api.Client{
		GQLClient: gqlClient,
		Token:     os.Getenv("STORYSCRIPT_TOKEN"),
	}

	appsCmd := &cmd.AppsCommand{
		ListCommand: &cmd.ListCommand{
			UI:         ui,
			AppFetcher: apiClient,
		},
	}

	parser := flags.NewParser(&Defaults{}, flags.Default)
	parser.AddCommand("apps", "create, list, and manage apps", "", appsCmd)

	parser.Parse()
}
