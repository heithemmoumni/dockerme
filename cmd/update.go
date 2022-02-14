package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

type updateOption struct {
	Path       string
	TemplateFS embed.FS
}

func UpdateCommand(
	templateFS embed.FS,
) *cobra.Command {

	o := &updateOption{
		TemplateFS: templateFS,
	}

	var cmd = &cobra.Command{
		Use:     "update",
		Short:   "Update docker container",
		Example: "dockerme update",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Update()
		},
	}
	return cmd
}

func (o *updateOption) Update() error {
	fmt.Printf("update ...")
	return nil
}
