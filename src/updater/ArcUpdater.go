package updater

/*
	Author: https://github.com/dayc0re
*/

const Gw2Exe = "Gw2-64.exe"
const d3d9Name = "d3d9.dll"
const d3d9Md5Name = "d3d9.dll.md5sum"
const d3d9Url = "https://www.deltaconnected.com/arcdps/x64/d3d9.dll"
const d3d9Md5Url = "https://www.deltaconnected.com/arcdps/x64/d3d9.dll.md5sum"

/*
	Interface for specifying an API to encapsulate the functionality regarding arc-dps.
*/
type ArcUpdater interface {
	/*
		Checks wether arc-dps is installed or not.

		@return true if the d3d9.dll file is found in the bin64 directory
	*/
	IsInstalled() bool

	/*
		Checks wether arc-dps is up to date or not.

		@return true if the d3d9.dll file has the same checksum as the d3d9.dll.md5sum file
	*/
	IsUpToDate() bool

	/*
		Downloads the latest version of the d3d9.dll file.
	*/
	DownloadLatestVersion()

	/*
		Downloads the latest version of the d3d9.dll.md5sum file.
	*/
	DownloadChecksumFile()

	/*
		Removes the checksum file.
	*/
	RemoveChecksumFile()

	/*
		Returns the installtion path of arc-dps.
	*/
	InstallationPath() string
}
