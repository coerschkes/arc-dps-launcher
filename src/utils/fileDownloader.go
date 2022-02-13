package utils

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/coerschkes/arc-dps-launcher/src/logging"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/

var fdLogger logging.Logger

func init() {
	fdLogger = logging.GetLogger("fileDownloader.go")
}

/*
	Downloads a file from the given url and saves it to the specified filepath.
	Replaces Linux-Path-Ref with Windows-Path-Ref.
	Overrides an existing file.

	@param url - the file url
	@param filePath - the target file path
*/
func DownloadFile(url string, path string) {
	operateOnResponseBody(strings.ReplaceAll(url, "\\", "/"), func(body io.ReadCloser) {
		OperateOnFile(path, func(file *os.File) {
			if _, err := io.Copy(file, body); err != nil {
				fdLogger.LogError(err)
				panic(err)
			} else {
				fdLogger.Log("Downloaded file: " + url)
			}
		})
	})
}

/*
	Downloads the content of a given url. Executes a defined function on the response body.
	Closes the response body afterwards.

	@param url - the url of the content to be downloaded
	@param f - the function that will be executed on the response body
*/
func operateOnResponseBody(url string, f func(body io.ReadCloser)) {
	if resp, err := http.Get(url); err != nil {
		fdLogger.LogError(err)
		panic(err)
	} else {
		f(resp.Body)
		defer resp.Body.Close()
	}
}
