package updater

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/coerschkes/arc-dps-launcher/src/utils"
)

const binFolderPath = "bin"
const d3d9Name = "d3d9.dll"
const d3d9Md5Name = "d3d9.dll.md5sum"
const d3d9Url = "https://www.deltaconnected.com/arcdps/x64/d3d9.dll"
const d3d9Md5Url = "https://www.deltaconnected.com/arcdps/x64/d3d9.dll.md5sum"

type IArcUpdater interface {
	IsInstalled() bool
	IsUpToDate() bool
	ChecksumFilePath() string
	TempDirPath() string
	DownloadLatestVersion()
	DownloadChecksumFile()
}

type ArcUpdater struct{}

var tmpDir string

func init() {
	if dir, err := os.MkdirTemp(binFolderPath, "tmp"); err != nil {
		panic(err)
	} else {
		tmpDir = dir
	}
}

func NewArcUpdater() ArcUpdater {
	return ArcUpdater{}
}

func (impl *ArcUpdater) IsInstalled() bool {
	_, err := os.Open(impl.installationPath()) //TODO: check if necessary to close
	return err == nil
}

func (impl *ArcUpdater) IsUpToDate() bool {
	return impl.parseChecksum() == utils.CalculateChecksum(impl.installationPath())
}

func (impl *ArcUpdater) ChecksumFilePath() string {
	return tmpDir + "\\" + d3d9Md5Name
}

func (impl *ArcUpdater) installationPath() string {
	return binFolderPath + "\\" + d3d9Name
}

func (impl *ArcUpdater) TempDirPath() string {
	return tmpDir
}

func (impl *ArcUpdater) DownloadLatestVersion() {
	log.Println("Downloading latest arc-dps version")
	utils.DownloadFile(d3d9Url, impl.installationPath())
}

func (impl *ArcUpdater) DownloadChecksumFile() {
	log.Println("Downloading arcdps checksum file")
	utils.DownloadFile(d3d9Md5Url, impl.ChecksumFilePath())
}

func (impl *ArcUpdater) parseChecksum() string {
	content, err := ioutil.ReadFile(impl.ChecksumFilePath())
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), " ")[0]
}
