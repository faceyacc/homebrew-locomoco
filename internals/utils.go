package internals

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type column []int

// parseFileLinesToSlice parses the content of a given file path
// and returns a list of repos.
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

// sortMapIntoSlice takes a map and returns a slice with the map
// keys ordered by their integer value.
func sortMapIntoSlice(m map[int]int) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	return keys
}

// buildCols returns a map with rows and columns to print to cli
// given an slice of key and commits.
func buildCols(keys []int, commits map[int]int) map[int]column {
	cols := make(map[int]column)
	col := column{}

	for _, k := range keys {
		week := int(k / 7)
		dayInWeek := k % 7

		// Create a new column on Sunday.
		if dayInWeek == 0 {
			col = column{}
		}
		col = append(col, commits[k])

		if dayInWeek == 6 {
			cols[week] = col
		}

	}
	return cols
}

// printCells prints cells of the graph.
func printCells(cols map[int]column) {
	printMonths()
	for j := 6; j >= 0; j-- {
		for i := weeksInLastSixMonths + 1; i >= 0; i-- {
			if i == weeksInLastSixMonths+1 {
				printDayCol(j)
			}
			if col, ok := cols[i-1]; ok {
				if i == 0 && j == calcOffset()-1 {
					printCell(col[j], true)
					continue
				} else {
					if len(col) > j {
						printCell(col[j], false)
						continue
					}
				}
			}
			printCell(0, false)
		}
		fmt.Printf("\n")
	}
}

// printMonths print the month names based on weeks.
func printMonths() {
	week := getBeginningOfDay(time.Now()).Add(-(daysInLastSixMonths * time.Hour * 24))
	month := week.Month()
	fmt.Printf("      ")
	for {
		if week.Month() != month {
			fmt.Printf("%s", week.Month().String()[:3])
			month = week.Month()
		} else {
			fmt.Printf("   ")
		}

		week = week.Add(7 * time.Hour * 24)
		if week.After(time.Now()) {
			break
		}
	}
	fmt.Printf("\n")
}

// printDayCol prints the day name, given the day number.
func printDayCol(day int) {
	out := "     "
	switch day {
	case 1:
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}
	fmt.Printf(out)
}

func printCell(val int, today bool) {
	escape := "\033[0;37;30m"

	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	if today {
		escape = "\033[1;37;45m"
	}

	if val == 0 {
		fmt.Printf(escape + " - " + "\033[0m")
		return
	}

	str := " %d "
	switch {
	case val >= 10:
		str = " %d "
	case val >= 100:
		str = "%d"
	}
	fmt.Printf(escape+str+"\033[0m", val)
}
