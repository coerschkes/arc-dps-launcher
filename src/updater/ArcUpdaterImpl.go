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
	logger = logging.GetLogger()
}

func NewArcUpdater(tmpDir string) ArcUpdater {
	return &ArcUpdaterImpl{tmpDir}
}

func (impl *ArcUpdaterImpl) IsInstalled() bool {
	_, err := os.Open(impl.installationPath())
	return err == nil
}

func (impl *ArcUpdaterImpl) IsUpToDate() bool {
	return impl.parseChecksum() == utils.CalculateChecksum(impl.installationPath())
}

func (impl *ArcUpdaterImpl) ChecksumFilePath() string {
	return impl.tmpDir + "\\" + d3d9Md5Name
}

func (impl *ArcUpdaterImpl) installationPath() string {
	return BinFolderPath + "\\" + d3d9Name
}

func (impl *ArcUpdaterImpl) DownloadLatestVersion() {
	logger.Log("Downloading latest arc-dps version")
	utils.DownloadFile(d3d9Url, impl.installationPath())
}

func (impl *ArcUpdaterImpl) DownloadChecksumFile() {
	logger.Log("Downloading arcdps checksum file")
	utils.DownloadFile(d3d9Md5Url, impl.ChecksumFilePath())
}

func (impl *ArcUpdaterImpl) parseChecksum() string {
	content, err := ioutil.ReadFile(impl.ChecksumFilePath())
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), " ")[0]
}
