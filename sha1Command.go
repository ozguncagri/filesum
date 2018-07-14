package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func sha1Calculator(cmd *cobra.Command, args []string) {
	for _, v := range args {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		hash := sha1.New()
		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal(err)
		}

		hashBytes := hash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)

		fmt.Printf("SHA1 (%v) = %v\n", v, hashString)
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha1",
		Short:   "Calculate SHA1 signature of file",
		Long:    "Calculates SHA1 signature of file and prints to stdout",
		Example: "filesum sha1 <filePath>",
		Args:    cobra.MinimumNArgs(1),
		Run:     sha1Calculator,
	})
}
