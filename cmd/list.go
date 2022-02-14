package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

type listOption struct {
	Path       string
	TemplateFS embed.FS
}

func ListCommand(
	templateFS embed.FS) *cobra.Command {

	o := &listOption{
		TemplateFS: templateFS,
	}

	var cmd = &cobra.Command{
		Use:     "list",
		Short:   "List all support Dockerfile",
		Example: "dockerme list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.List()
		},
	}
	/*cmd.Flags().StringArrayP(
		&o.
	)*/
	return cmd
}

func (o *listOption) List() error {
	fmt.Println("ok")
	return nil
}
