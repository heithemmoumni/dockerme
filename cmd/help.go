package cmd

import (
	"embed"

	"github.com/spf13/cobra"
)

func Help(templateFS embed.FS) error {
	cmd := &cobra.Command{
		Use:   "dockerme",
		Short: "A tool help you to generate a Dockerfile",
	}

	cmd.AddCommand(
		GenCommand(templateFS),
		ListCommand(templateFS),
		UpdateCommand(templateFS),
		CleanCommand(templateFS),
	)

	return cmd.Execute()
}
