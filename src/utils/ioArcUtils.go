package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/coerschkes/arc-dps-launcher/src/logging"
)

/*
	Author: https://github.com/dayc0re
*/

var auLogger logging.Logger

func init() {
	auLogger = logging.GetLogger("ioArcUtils.go")
}

/*
	Creates a temporary directory inside of the specified root folder.

	@param rootDir - path of the root directory
	@return the path of the temporary directory
*/
func CreateTempFolder(rootDir string) string {
	if dir, err := os.MkdirTemp(rootDir, "tmp"); err != nil {
		auLogger.LogError(err)
		panic(err)
	} else {
		return dir
	}
}

/*
	Opens or creates and opens a file at the specified path. Executes a defined function on the opened file.
	Closes the file afterwards.

	@param path - the path to the file
	@param f - the function that will be executed on the file
*/
func OperateOnFile(path string, f func(file *os.File)) {
	if file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666); err != nil {
		auLogger.LogError(err)
		panic(err)
	} else {
		f(file)
		defer file.Close()
	}
}

/*
	Calculates the checksum of a given file.

	@param - the path to the file
	@return the checksum of the file
*/
func CalculateChecksum(path string) string {
	algorithm := md5.New()
	OperateOnFile(path, func(file *os.File) {
		if _, err := io.Copy(algorithm, file); err != nil {
			auLogger.LogError(err)
			panic(err)
		}
	})
	return fmt.Sprintf("%x", algorithm.Sum(nil))
}
