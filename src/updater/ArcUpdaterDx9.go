package updater

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/dayc0re/arc-dps-launcher/src/logging"
	"github.com/dayc0re/arc-dps-launcher/src/utils"
)

/*
	Author: https://github.com/dayc0re
*/

/*
	The implementation of the ArcUpdater interface.

	Needs a storage location for downloading the checksum file. In best case uses a temporary directory
	that can be cleaned afterwards.

*/

const binFolderPath = "bin64"

type ArcUpdaterDx9 struct{}

var logger logging.Logger

func init() {
	logger = logging.GetLogger("arcUpdaterImpl.go")
}

func NewArcUpdaterDx9() ArcUpdater {
	return &ArcUpdaterDx9{}
}

func (au ArcUpdaterDx9) IsInstalled() bool {
	_, err := os.Open(au.InstallationPath())
	return err == nil
}

func (au ArcUpdaterDx9) IsUpToDate() bool {
	return au.parseChecksum() == utils.CalculateChecksum(au.InstallationPath())
}

func (au ArcUpdaterDx9) DownloadLatestVersion() {
	logger.Log("Downloading latest arc-dps version")
	utils.DownloadFile(d3d9Url, au.InstallationPath())
}

func (au ArcUpdaterDx9) DownloadChecksumFile() {
	logger.Log("Downloading arcdps checksum file")
	utils.DownloadFile(d3d9Md5Url, d3d9Md5Name)
}

func (au ArcUpdaterDx9) InstallationPath() string {
	return binFolderPath + "\\" + d3d9Name
}

func (au ArcUpdaterDx9) RemoveChecksumFile() {
	os.Remove(d3d9Md5Name)
}

func (au ArcUpdaterDx9) parseChecksum() string {
	content, err := ioutil.ReadFile(d3d9Md5Name)
	if err != nil {
		logger.LogError(err)
		panic(err)
	}
	return strings.Split(string(content), " ")[0]
}
