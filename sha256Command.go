package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func sha256Calculator(cmd *cobra.Command, args []string) {
	for _, v := range args {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		hash := sha256.New()
		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal(err)
		}

		hashBytes := hash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)

		fmt.Printf("SHA256 (%v) = %v\n", v, hashString)
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha256",
		Aliases: []string{"sha2"},
		Short:   "Calculate SHA256 (SHA2) signature of file",
		Long:    "Calculates SHA256 (SHA2) signature of file and prints to stdout",
		Example: "filesum sha256 <filePath>",
		Args:    cobra.MinimumNArgs(1),
		Run:     sha256Calculator,
	})
}
