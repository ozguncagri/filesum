package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func sha512Calculator(filePath string) {
	file, err := os.Open(filePath)
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
