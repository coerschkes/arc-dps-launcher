package updater

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/coerschkes/arc-dps-launcher/src/logging"
	"github.com/coerschkes/arc-dps-launcher/src/utils"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/

/*
	The implementation of the ArcUpdater interface.

	Needs a storage location for downloading the checksum file. In best case uses a temporary directory
	that can be cleaned afterwards.

*/

type ArcUpdaterImpl struct {
	tmpDir string
}

var logger logging.Logger

func init() {
	logger = logging.GetLogger("arcUpdaterImpl.go")
}

func NewArcUpdater(tmpDir string) ArcUpdater {
	return &ArcUpdaterImpl{tmpDir}
}

func (au ArcUpdaterImpl) IsInstalled() bool {
	_, err := os.Open(au.installationPath())
	return err == nil
}

func (au ArcUpdaterImpl) IsUpToDate() bool {
	return au.parseChecksum() == utils.CalculateChecksum(au.installationPath())
}

func (au ArcUpdaterImpl) ChecksumFilePath() string {
	return au.tmpDir + "\\" + d3d9Md5Name
}

func (au ArcUpdaterImpl) DownloadLatestVersion() {
	logger.Log("Downloading latest arc-dps version")
	utils.DownloadFile(d3d9Url, au.installationPath())
}

func (au ArcUpdaterImpl) DownloadChecksumFile() {
	logger.Log("Downloading arcdps checksum file")
	utils.DownloadFile(d3d9Md5Url, au.ChecksumFilePath())
}

func (au ArcUpdaterImpl) installationPath() string {
	return BinFolderPath + "\\" + d3d9Name
}

func (au ArcUpdaterImpl) parseChecksum() string {
	content, err := ioutil.ReadFile(au.ChecksumFilePath())
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), " ")[0] //TODO: check if string can be acessed as array
}
