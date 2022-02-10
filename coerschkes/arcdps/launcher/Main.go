package main

import (
	"errors"
	"log"
	"os"

	"github.com/coerschkes/arcdps/launcher/utils"
)

func init() {
	downloadArcChecksum()
}

func main() {
	defer cleanUp()
	newMd5 := readMd5FromMd5File()

	if currentMd5 := calculateMd5(binFolderPath + "\\" + d3d9Name); newMd5 != currentMd5 {
		log.Println("New version found!")
		downloadLatestArc()
	}
}

func cleanUp() {
	defer os.RemoveAll(tmpDir)
	if err := recover(); err != nil {
		var pathError *os.PathError
		if errors.As(err.(error), &pathError) {
			log.Println("Arc not found! Starting installation..")
			downloadLatestArc()
		} else {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	log.Println("Everything up to date! Starting GuildWars2..")
}

func downloadLatestArc() {
	log.Println("Downloading latest d3d9.dll..")
	utils.DownloadFile(d3d9Url, binFolderPath+"\\"+d3d9Name)
}

func downloadArcChecksum() {
	log.Println("Downloading checksum file..")
	downloadFile(d3d9Md5Url, tmpDir+"\\"+d3d9Md5Name)
}
