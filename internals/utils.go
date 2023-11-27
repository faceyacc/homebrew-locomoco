package internals

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const outOfRange = 99999

// parseFileLinesToSlice parses the content of a given file path.
func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

// joinSlices appends new slice into existing slice if it's not already there.
func joinSlices(new, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

// sliceContains checks if a slice contains a given value.
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// dumpStringSliceToFile writes content to filePath.
func dumpStringSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(filePath, []byte(content), 0755)
}

// fillCommits returns a map of git commits given a repo and email.
func fillCommits(email, path string, commits map[int]int) map[int]int {

	// Create a git repo object.
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	// Get HEAD ref.
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	// Get commit history starting from HEAD ref.
	commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	offset := calcOffset()
	err = commitIter.ForEach(func(c *object.Commit) error {
		daysAgo := daysSinceDate(c.Author.When) + offset

		if c.Author.Email != email {
			return nil
		}

		// Set an arbitrary out of range value.
		if daysAgo != outOfRange {
			commits[daysAgo]++
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return commits
}

// getBeginningOfDay calculates the start time of a given day.
func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay
}

func daysSinceDate(date time.Time) int {
	days := 0
	now := getBeginningOfDay(time.Now())
	for date.Before(now) {
		date = date.Add(time.Hour * 24)
		days++
		if days > daysInLastSixMonths {
			return outOfRange
		}
	}
	return days
}

// calcOffset returns the amount of days missing to fill
// last row of git commit stats graph.
func calcOffset() int {
	var offset int
	weekday := time.Now().Weekday()

	switch weekday {
	case time.Sunday:
		offset = 7
	case time.Monday:
		offset = 6
	case time.Tuesday:
		offset = 5
	case time.Wednesday:
		offset = 4
	case time.Thursday:
		offset = 3
	case time.Friday:
		offset = 2
	case time.Saturday:
		offset = 1
	}
	return offset
}
