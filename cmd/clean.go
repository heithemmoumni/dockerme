package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

type cleanOption struct {
	Path       string
	TemplateFS embed.FS
}

func CleanCommand(templateFS embed.FS) *cobra.Command {
	o := &cleanOption{
		TemplateFS: templateFS,
	}

	var cmd = &cobra.Command{
		Use:     "clean",
		Short:   "Clean docker",
		Example: "dockerme clean",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Clean()
		},
	}
	return cmd
}

func (o *cleanOption) Clean() error {
	fmt.Printf("clean ...")
	return nil
}
