package api

import (
	"context"
	"fmt"
	"story"

	"github.com/machinebox/graphql"
)

type Client struct {
	GQLClient *graphql.Client
	Token     string
}

// No tests for this, just for demonstration purposes

func (c Client) FetchApps() ([]story.App, error) {
	req := graphql.NewRequest(`query {
      allApps(condition: {deleted: false}) {
        nodes{
          name
        }
      }
    }`)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	type allAppsData struct {
		AllApps struct {
			Nodes []story.App `json:"nodes"`
		} `json:"allApps"`
	}

	var appsData allAppsData
	if err := c.GQLClient.Run(context.Background(), req, &appsData); err != nil {
		return nil, err
	}

	return appsData.AllApps.Nodes, nil
}
