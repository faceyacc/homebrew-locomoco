package internals

import (
	"os"
)

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
