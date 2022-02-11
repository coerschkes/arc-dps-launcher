package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
	Creates a temporary directory inside of the specified root folder.

	@param rootDir - path of the root directory
	@return the path of the temporary directory
*/
func CreateTempFolder(rootDir string) string {
	if dir, err := os.MkdirTemp(rootDir, "tmp"); err != nil {
		panic(err)
	} else {
		return dir
	}
}

/*
	Downloads a file from the given url and saves it to the specified filepath.
	Replaces Linux-Path-Ref with Windows-Path-Ref.
	Overrides an existing file.

	@param url - the file url
	@param filePath - the target file path
*/
func DownloadFile(url string, filePath string) {
	operateOnResponseBody(strings.ReplaceAll(url, "\\", "/"), func(body io.ReadCloser) {
		operateOnFile(filePath, func(file *os.File) {
			if _, err := io.Copy(file, body); err != nil {
				panic(err)
			} else {
				log.Println("Downloaded file: " + url)
			}
		})
	})
}

func operateOnResponseBody(url string, f func(body io.ReadCloser)) {
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else {
		f(resp.Body)
		defer resp.Body.Close()
	}
}

func operateOnFile(path string, f func(file *os.File)) {
	if file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666); err != nil {
		panic(err)
	} else {
		f(file)
		defer file.Close()
	}
}
