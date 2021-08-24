package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

type genOption struct {
	Path       string
	TemplateFS embed.FS
}

func GenCommand(templateFS embed.FS) *cobra.Command {

	o := &genOption{
		TemplateFS: templateFS,
	}

	var cmd = &cobra.Command{
		Use:     "gen",
		Short:   "Generate a Dockerfile",
		Example: "dockerme gen",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}
	/*cmd.Flags().StringArrayP(
		&o.
	)*/
	return cmd
}

func (o *genOption) Run() error {
	fmt.Println("ok")
	return nil
}
