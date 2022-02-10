package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/**
	Downloads the file from the fileDownloader#url and saves it to the fileDownloader#filepath.
	Replaces Linux-Path-Ref with Windows-Path-Ref.
	Overrides an existing file.
**/
func DownloadFile(url string, filePath string) {
	url = strings.ReplaceAll(url, "\\", "/")
	readCloser := downloadBody(url)
	defer readCloser.Close()
	outputFile := createOsFile(filePath)
	defer outputFile.Close()

	if _, err := io.Copy(outputFile, readCloser); err != nil {
		panic(err)
	} else {
		log.Println("DOWNLOADED: " + url + "' AS '" + filePath + "'.")
	}
}

func downloadBody(url string) io.ReadCloser {
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else {
		return resp.Body
	}
}

func createOsFile(filePath string) *os.File {
	if out, err := os.Create(filePath); err != nil {
		panic(err)
	} else {
		return out
	}
}
