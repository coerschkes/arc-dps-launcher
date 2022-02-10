package main

import (
	"os"

	"github.com/coerschkes/arc-dps-launcher/src/updater"
)

var arcUpdater updater.IArcUpdater

func init() {
	arcUpdater = updater.NewArcUpdater()
	arcUpdater.DownloadChecksumFile()
}

func main() {
	defer os.RemoveAll(arcUpdater.TempDirPath())
	if !arcUpdater.IsInstalled() && (arcUpdater.IsInstalled() && !arcUpdater.IsUpToDate()) {
		arcUpdater.DownloadLatestVersion()
	}
	//launch gw2
}
