package cmd

type AppsCommand struct {
	ListCommand *ListCommand `command:"list" description:"list apps"`
}
