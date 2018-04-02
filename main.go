package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filesum",
	Short: "File integrity calculator",
	Long:  "Filesum calculates integrity signatures of file with different algorithms",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
