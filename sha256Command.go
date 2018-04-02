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
	file, err := os.Open(args[0])
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

	fmt.Println(hashString)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha256",
		Aliases: []string{"sha2"},
		Short:   "Calculate SHA256 (SHA2) signature of file",
		Long:    "Calculates SHA256 (SHA2) signature of file and prints to stdout",
		Example: "filesum sha256 <filePath>",
		Run:     sha256Calculator,
	})
}
