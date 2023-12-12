package internals

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

const daysInLastSixMonths = 183
const weeksInLastSixMonths = 26
const outOfRange = 99999

func RecursiveScanFolder(folder string) []string {
	return ScanGitFolders(make([]string, 0), folder)
}

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

// AddNewSliceElementsToFile stores a given slice of paths to the filesystem.
func AddNewSliceElementsToFile(filePath string, newRepos []string) {

	existingRepos := parseFileLinesToSlice(filePath)

	repos := joinSlices(newRepos, existingRepos)

	dumpStringSliceToFile(repos, filePath)

}

// ProcessRepos returns the commits made in the
// past 6 months given an user email.
func ProcessRepos(email string) (map[int]int, bool) {

	defer func() {
		if err := recover(); err != nil {
			c := color.New(color.FgYellow)
			c.Print("\n\nLooks like you might've call locomoco from outside project directory. Try calling locomoco from your project directory.\n\n\n")
		}
	}()

	// creates file at Users/Ty/.locomocostats
	filepath := GetDotFilePath()

	repos := parseFileLinesToSlice(filepath)

	daysInMap := daysInLastSixMonths

	commits := make(map[int]int, daysInMap)

	for day := daysInMap; day > 0; day-- {
		commits[day] = 0
	}

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}

	return commits, true
}

func PrintCommitStats(commits map[int]int) {
	keys := sortMapIntoSlice(commits)
	cols := buildCols(keys, commits)
	printCells(cols)
}
