package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func sha512Calculator(cmd *cobra.Command, args []string) {
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha512.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println(hashString)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha512",
		Aliases: []string{"sha3"},
		Short:   "Calculate SHA512 (SHA3) signature of file",
		Long:    "Calculates SHA512 (SHA3) signature of file and prints to stdout",
		Example: "filesum sha512 <filePath>",
		Run:     sha512Calculator,
	})
}
