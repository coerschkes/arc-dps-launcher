package main

import (
	"os"
	"os/exec"

	"github.com/coerschkes/arc-dps-launcher/src/logging"
	"github.com/coerschkes/arc-dps-launcher/src/updater"
	"github.com/coerschkes/arc-dps-launcher/src/utils"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/

var tmpDir string
var arcUpdater updater.ArcUpdater
var logger logging.Logger

func init() {
	tmpDir = utils.CreateTempFolder(updater.BinFolderPath)
	arcUpdater = updater.NewArcUpdater(tmpDir)
	logger = logging.GetLogger("main.go")
	logger.Log("---- Arc launcher initialized ----")
	arcUpdater.DownloadChecksumFile()
}

func main() {
	routine()
	defer startGuildWars2()
	defer cleanup()
}

func routine() {
	if !arcUpdater.IsInstalled() || (arcUpdater.IsInstalled() && !arcUpdater.IsUpToDate()) {
		logger.Log("New Version found!")
		arcUpdater.DownloadLatestVersion()
	} else {
		logger.Log("Arc-dps is up to date!")
	}
}

func startGuildWars2() {
	logger.Log("Launching Guild Wars 2..")
	exec.Command(updater.Gw2Exe).Start()
}

func cleanup() {
	os.RemoveAll(tmpDir)
}
