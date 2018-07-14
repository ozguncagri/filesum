package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func md5Calculator(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	for _, v := range args {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		hash := md5.New()
		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal(err)
		}

		hashBytes := hash.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)

		fmt.Printf("MD5 (%v) = %v\n", v, hashString)
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "md5",
		Short:   "Calculate MD5 signature of file",
		Long:    "Calculates MD5 signature of file and prints to stdout",
		Example: "filesum md5 <filePath>",
		Run:     md5Calculator,
	})
}
