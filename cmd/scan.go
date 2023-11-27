package cmd

import (
	"fmt"
	"loco-moco/internals"
)

func recursiveScanFolder(folder string) []string {
	return internals.ScanGitFolders(make([]string, 0), folder)
}

func scan(path string) {
	fmt.Printf("Found folders: \n\n")

	// repos := ScanGitFolders(make([]string, 0), path)

	repos := recursiveScanFolder(path)
	filePath := internals.GetDotFilePath()
	internals.AddNewSliceElementsToFile(filePath, repos)
	fmt.Printf("\n\nSuccessfully added\n\n")

}
