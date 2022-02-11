package main

import (
	"log"
	"os"

	"github.com/coerschkes/arc-dps-launcher/src/updater"
)

var arcUpdater updater.ArcUpdater

func init() {
	arcUpdater = updater.NewArcUpdater()
	arcUpdater.DownloadChecksumFile()
}

func main() {
	defer os.RemoveAll(arcUpdater.TempDirPath())
	if !arcUpdater.IsInstalled() || (arcUpdater.IsInstalled() && !arcUpdater.IsUpToDate()) {
		log.Println("New Version found!")
		arcUpdater.DownloadLatestVersion()
	} else {
		log.Println("Arc-dps is up to date!")
	}
	//TODO: launch gw2
}
