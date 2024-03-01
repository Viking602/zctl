package main

import (
	"github.com/spf13/cobra"
	"log"
	"zctl/internal/add"
)

var rootCmd = &cobra.Command{
	Use:     "zctl",
	Short:   "zctl is a CLI tool for managing Project Zero",
	Long:    `zctl is a CLI tool for managing Project Zero. It is a tool to manage the Project Zero application.`,
	Version: version,
}

func init() {
	rootCmd.AddCommand(add.CmdAdd)
}

// Zctl is a CLI tool for managing Project Zero
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
