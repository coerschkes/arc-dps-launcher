package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/coerschkes/arc-dps-launcher/src/updater"
	"github.com/coerschkes/arc-dps-launcher/src/utils"
)

var tmpDir string
var arcUpdater updater.ArcUpdater

func init() {
	tmpDir = utils.CreateTempFolder(updater.BinFolderPath)
	arcUpdater = updater.NewArcUpdater(tmpDir)
	arcUpdater.DownloadChecksumFile()
}

func main() {
	routine()
	defer os.RemoveAll(tmpDir)
	defer startGuildWars2()
}

func routine() {
	if !arcUpdater.IsInstalled() || (arcUpdater.IsInstalled() && !arcUpdater.IsUpToDate()) {
		log.Println("New Version found!")
		arcUpdater.DownloadLatestVersion()
	} else {
		log.Println("Arc-dps is up to date!")
	}
}

func startGuildWars2() {
	exec.Command(updater.Gw2Exe).Start()
}
