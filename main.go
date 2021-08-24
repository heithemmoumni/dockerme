package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/heithemmoumni/dockerme/cmd"
)

var templateFS embed.FS

func main() {
	if err := cmd.Help(templateFS); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
