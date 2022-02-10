package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func CalculateChecksum(path string) string {
	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()
		algorithm := md5.New()
		if _, err := io.Copy(algorithm, file); err != nil {
			panic(err)
		} else {
			return fmt.Sprintf("%x", algorithm.Sum(nil))
		}
	}
}
