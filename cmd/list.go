package cmd

import "story"

//go:generate counterfeiter . UI

type UI interface {
	DisplayText(text string, data ...map[string]interface{})
	DisplayTable(headers []string, data [][]string)
}

//go:generate counterfeiter . AppFetcher

type AppFetcher interface {
	FetchApps() ([]story.App, error)
}

type ListCommand struct {
	UI         UI
	AppFetcher AppFetcher
}

func (c *ListCommand) Execute(args []string) error {
	apps, err := c.AppFetcher.FetchApps()
	if err != nil {
		return err
	}

	if len(apps) == 0 {
		c.UI.DisplayText("No apps found.")
		c.UI.DisplayText("Create your first app using `story apps create`")
		return nil
	}

	headers := []string{"NAME"}
	data := [][]string{}
	for _, a := range apps {
		data = append(data, []string{a.Name})
	}

	c.UI.DisplayTable(headers, data)
	return nil
}
