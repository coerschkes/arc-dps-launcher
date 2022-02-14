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
	The implementation for the Dx11 installation of the ArcUpdater interface.
*/

const dx11ArcDpsName = "d3d11.dll"

//directx11: save d3d9.dll into 'gw2 install dir/' and rename to d3d11.dll or dxgi.dll while the game is not running.
type ArcUpdaterDx11 struct {
	logger logging.Logger
}

func NewArcUpdaterDx11() ArcUpdater {
	return &ArcUpdaterDx11{logging.GetLogger("arcUpdaterDx11.go")}
}

func (au ArcUpdaterDx11) IsInstalled() bool {
	_, err := os.Open(au.InstallationPath())
	return err == nil
}

func (au ArcUpdaterDx11) IsUpToDate() bool {
	return au.parseChecksum() == utils.CalculateChecksum(au.InstallationPath())
}

func (au ArcUpdaterDx11) DownloadLatestVersion() {
	au.logger.Log("Downloading latest arc-dps version")
	utils.DownloadFile(arcDpsUrl, au.InstallationPath())
}

func (au ArcUpdaterDx11) DownloadChecksumFile() {
	au.logger.Log("Downloading arcdps checksum file")
	utils.DownloadFile(checksumFileUrl, checksumFile)
}

func (au ArcUpdaterDx11) InstallationPath() string {
	return dx11ArcDpsName
}

func (au ArcUpdaterDx11) RemoveChecksumFile() {
	os.Remove(checksumFile)
}

func (au ArcUpdaterDx11) GetVersion() string {
	return "DX11"
}

func (au ArcUpdaterDx11) parseChecksum() string {
	content, err := ioutil.ReadFile(checksumFile)
	if err != nil {
		au.logger.LogError(err)
		panic(err)
	}
	return strings.Split(string(content), " ")[0]
}
