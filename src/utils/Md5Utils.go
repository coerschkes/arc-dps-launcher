package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/

/*
	Calculates the checksum of a given file.

	@param - the path to the file
	@return the checksum of the file
*/
func CalculateChecksum(path string) string {
	algorithm := md5.New()
	OperateOnFile(path, func(file *os.File) {
		if _, err := io.Copy(algorithm, file); err != nil {
			panic(err)
		}
	})
	return fmt.Sprintf("%x", algorithm.Sum(nil))
}
