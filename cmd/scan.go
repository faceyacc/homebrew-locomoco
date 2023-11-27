package cmd

import (
	"fmt"
	. "loco-moco/internals"
)

func recursiveScanFolder(folder string) []string {
	return ScanGitFolders(make([]string, 0), folder)
}

func scan(path string) {
	fmt.Printf("Found folders: \n\n")
	repos := recursiveScanFolder(path)
	filePath := GetDotFilePath()
	AddNewSliceElementsToFile(filePath, repos)
	fmt.Printf("\n\nSuccessfully added\n\n")

}
