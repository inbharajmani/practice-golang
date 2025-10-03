package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func sha1sum(fileName string) (sha1sum string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("unable to open file %q - %w", fileName, err)
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, "gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("unable to unzip file %q - %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", fmt.Errorf("unable to copy %q - %w", fileName, err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	fileName := "/home/inbmani/Study/Devops/GO/README.md"
	if sha1sum, err := sha1sum(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sha1sum)
	}
}
