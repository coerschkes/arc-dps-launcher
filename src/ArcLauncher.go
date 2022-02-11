package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/coerschkes/arc-dps-launcher/src/updater"
)

const gw2Exe = "Gw2-64.exe"

var arcUpdater updater.ArcUpdater

func init() {
	arcUpdater = updater.NewArcUpdater()
	arcUpdater.DownloadChecksumFile()
}

func main() {
	if !arcUpdater.IsInstalled() || (arcUpdater.IsInstalled() && !arcUpdater.IsUpToDate()) {
		log.Println("New Version found!")
		arcUpdater.DownloadLatestVersion()
	} else {
		log.Println("Arc-dps is up to date!")
	}
	defer os.RemoveAll(arcUpdater.TempDirPath())
	defer exec.Command(gw2Exe).Start()
}
