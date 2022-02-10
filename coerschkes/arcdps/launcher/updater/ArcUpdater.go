package updater

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/coerschkes/arcdps/launcher/github.com/coerschkes/arcdps/launcher/utils"
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
	InstallationPath() string
	TempDirPath() string
	DownloadLatestVersion()
	DownloadChecksumFile()
	ParseChecksum() string
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

func NewArcUtility() IArcUpdater {
	return &ArcUpdater{}
}

func (impl *ArcUpdater) IsInstalled() bool {
	_, err := os.Open(impl.InstallationPath())
	return err == nil
}

func (impl *ArcUpdater) IsUpToDate() bool {
	return impl.ParseChecksum() != CalculateChecksum(impl.InstallationPath())
}

func (impl *ArcUpdater) ChecksumFilePath() string {
	return tmpDir + "\\" + d3d9Md5Name
}

func (impl *ArcUpdater) InstallationPath() string {
	return binFolderPath + "\\" + d3d9Name
}

func (impl *ArcUpdater) TempDirPath() string {
	return tmpDir
}

func (impl *ArcUpdater) DownloadLatestVersion() {
	log.Println("Downloading latest d3d9.dll..")
	DownloadFile(d3d9Url, impl.InstallationPath())
}

func (impl *ArcUpdater) DownloadChecksumFile() {
	log.Println("Downloading checksum file..")
	utils.DownloadFile(d3d9Md5Url, impl.ChecksumFilePath())
}

func (impl *ArcUpdater) ParseChecksum() string {
	content, err := ioutil.ReadFile(impl.ChecksumFilePath())
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), " ")[0]
}
