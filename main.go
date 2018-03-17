package main

import (
	"log"

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
			md5Calculator(args[0])
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha1",
		Short:   "Calculate SHA1 signature of file",
		Long:    "Calculates SHA1 signature of file and prints to stdout",
		Example: "filesum sha1 <filePath>",
		Run: func(cmd *cobra.Command, args []string) {
			//sha1 calculation
			sha1Calculator(args[0])
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha256",
		Short:   "Calculate SHA256 signature of file",
		Long:    "Calculates SHA256 signature of file and prints to stdout",
		Example: "filesum sha256 <filePath>",
		Run: func(cmd *cobra.Command, args []string) {
			//sha1 calculation
			sha256Calculator(args[0])
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:     "sha512",
		Short:   "Calculate SHA512 signature of file",
		Long:    "Calculates SHA512 signature of file and prints to stdout",
		Example: "filesum sha512 <filePath>",
		Run: func(cmd *cobra.Command, args []string) {
			//sha1 calculation
			sha512Calculator(args[0])
		},
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
