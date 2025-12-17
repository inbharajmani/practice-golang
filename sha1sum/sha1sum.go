package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"
)

func sha1sum(file string) (sha1sum string, err error) {
	if _, err := os.Stat(file); err != nil {
		return "", err
	}
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h = sha1.New()
	if _, err 
}

func main() {
	file := "/home/inbmani/old_files/Study/Devops/GO/README.md"
	sha, err := sha1sum(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sha)
	// fmt.Printf("SHA1SUM of file %v is %v", file, sha1sum(file))
}
