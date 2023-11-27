package internals

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

// scanGitFolders recursively looks for .git folders.
func ScanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}

			// Ignore vendor and node_models files.
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			// Recursively call scanGitFolders to look for .git files.
			folders = ScanGitFolders(folders, path)
		}
	}

	return folders
}

// GetDotFilePath returns the path of the dotfile containing
// the database of repo paths.
func GetDotFilePath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := user.HomeDir + "/.loco-moco-stats"
	return dotFile
}

// AddNewSliceElementsToFile stores a given slice of paths to the filesystem.
func AddNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSliceToFile(repos, filePath)
}
