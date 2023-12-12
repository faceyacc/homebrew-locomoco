package cmd

import (
	"fmt"
	"locomoco/internals"
)

func scan(path string) {
	fmt.Printf("Found folders: \n\n")

	repos := internals.RecursiveScanFolder(path)

	filePath := internals.GetDotFilePath()

	internals.AddNewSliceElementsToFile(filePath, repos)

	fmt.Printf("\n\nSuccessfully added\n\n")

}
