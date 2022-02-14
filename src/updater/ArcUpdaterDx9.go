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
	The implementation for the Dx9 installation of the ArcUpdater interface.
*/

const dx9ArcDpsName = "d3d9.dll"

type ArcUpdaterDx9 struct {
	binFolderPath string
	logger        logging.Logger
}

func NewArcUpdaterDx9() ArcUpdater {
	return &ArcUpdaterDx9{"bin64", logging.GetLogger("arcUpdaterDx9.go")}
}

func (au ArcUpdaterDx9) IsInstalled() bool {
	_, err := os.Open(au.InstallationPath())
	return err == nil
}

func (au ArcUpdaterDx9) IsUpToDate() bool {
	return au.parseChecksum() == utils.CalculateChecksum(au.InstallationPath())
}

func (au ArcUpdaterDx9) DownloadLatestVersion() {
	au.logger.Log("Downloading latest arc-dps version")
	utils.DownloadFile(arcDpsUrl, au.InstallationPath())
}

func (au ArcUpdaterDx9) DownloadChecksumFile() {
	au.logger.Log("Downloading arcdps checksum file")
	utils.DownloadFile(checksumFileUrl, checksumFile)
}

func (au ArcUpdaterDx9) InstallationPath() string {
	return au.binFolderPath + "\\" + dx9ArcDpsName
}

func (au ArcUpdaterDx9) RemoveChecksumFile() {
	os.Remove(checksumFile)
}

func (au ArcUpdaterDx9) GetVersion() string {
	return "DX9"
}

func (au ArcUpdaterDx9) parseChecksum() string {
	content, err := ioutil.ReadFile(checksumFile)
	if err != nil {
		au.logger.LogError(err)
		panic(err)
	}
	return strings.Split(string(content), " ")[0]
}
