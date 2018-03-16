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

func main() {
	rootCmd := &cobra.Command{
		Use:   "filesum",
		Short: "File integrity calculator",
		Long:  "Filesum calculates integrity signatures of file with different algorithms",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:     "md5",
		Short:   "Calculate MD5 signature of file",
		Long:    "Calculates MD5 signature of file and prints to stdout",
		Example: "filesum md5 <filePath>",
		Run: func(cmd *cobra.Command, args []string) {
			//md5 calculation
			file, err := os.Open(args[0])
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			hash := md5.New()
			if _, err := io.Copy(hash, file); err != nil {
				log.Fatal(err)
			}

			hashBytes := hash.Sum(nil)[:16]
			md5String := hex.EncodeToString(hashBytes)

			fmt.Println(md5String)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha1",
		Short:   "Calculate SHA1 signature of file",
		Long:    "Calculates SHA1 signature of file and prints to stdout",
		Example: "filesum sha1 <filePath>",
		Run: func(cmd *cobra.Command, args []string) {
			//sha1 calculation
		},
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
