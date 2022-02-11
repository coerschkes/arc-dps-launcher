package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/

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
		OperateOnFile(filePath, func(file *os.File) {
			if _, err := io.Copy(file, body); err != nil {
				panic(err)
			} else {
				log.Println("Downloaded file: " + url)
			}
		})
	})
}

/*
	Opens or creates and opens a file at the specified path. Executes a defined function on the opened file.
	Closes the file afterwards.

	@param path - the path to the file
	@param f - the function that will be executed on the file
*/
func OperateOnFile(path string, f func(file *os.File)) {
	if file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666); err != nil {
		panic(err)
	} else {
		f(file)
		defer file.Close()
	}
}

/*
	Downloads the content of a given url. Executes a defined function on the response body.
	Closes the response body afterwards.

	@param url - the url of the content to be downloaded
	@param f - the function that will be executed on the response body
*/
func operateOnResponseBody(url string, f func(body io.ReadCloser)) {
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else {
		f(resp.Body)
		defer resp.Body.Close()
	}
}
