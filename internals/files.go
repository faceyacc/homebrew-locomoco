package internals

import (
	"errors"
	"log"
	"os"
	"os/user"
	"strings"
)

// GetCurrentDir returns path to current direcotry
func GetCurrentDir() string {
	curDir, _ := os.Getwd()
	dir := strings.Split(curDir, "/")
	curDir = "./" + dir[len(dir)-1]
	return curDir
}

func DotFileExist() bool {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.locomocostats"

	// TODO: print out the local path of the dotfile using logger
	fmt.Println(".locomocostats stored: ", dotFile)

	_, err = os.Stat(dotFile)
	return !errors.Is(err, os.ErrNotExist)
}

// GetDotFilePath returns the path of the dotfile containing
// the database of repo paths.
func GetDotFilePath() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.locomocostats"
	return dotFile
}

// GetShowMeDotFilePath returns the path of the dotfile containing
// the username for GitHub.
func GetShowMeDotFilePath() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.locomocoshowme"
	return dotFile
}

// openFile opens the file located at the given filePath.
// If the file dosen't exist it creates one.
func openFile(filePath string) *os.File {
	// Allow to Append and Write Only to file.
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		// Create file if it doesn't exist
		if os.IsNotExist(err) {
			_, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
			// f = file
		} else {

			panic(err)
		}
	}
	return f
}
