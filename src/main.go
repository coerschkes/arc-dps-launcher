package main

import (
	"os/exec"

	"github.com/dayc0re/arc-dps-launcher/src/logging"
	"github.com/dayc0re/arc-dps-launcher/src/updater"
)

/*
	Author: https://github.com/dayc0re
*/

var arcUpdater updater.ArcUpdater
var logger logging.Logger

func init() {
	arcUpdater = updater.NewArcUpdaterDx9()
	logger = logging.GetLogger("main.go")
	logger.Log("---- Arc launcher initialized ----")
	arcUpdater.DownloadChecksumFile()
}

func main() {
	routine()
	defer arcUpdater.RemoveChecksumFile()
	defer startGuildWars2()
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
