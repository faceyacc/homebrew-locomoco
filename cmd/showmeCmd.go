package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"locomoco/internals"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

// Item represents a single repo data structure
type Item struct {
	ID              int    `json:"id"`
	CreatedAt       string `json:"created_at"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	StargazersCount int    `json:"stargazers_count"`
}

type JSONData struct {
	Items []Item
}

func getUser(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(data)
}

func SetUsername(userName string) {

	// Create a dotfile
	dotFile := internals.GetShowMeDotFilePath()

	f, err := os.Create(dotFile)
	if err != nil {
		fmt.Print(err)
	}

	// append name to dot file
	_, err = f.Write([]byte(userName))
}

// TODO:
// 1. Get dotfil
// 2. If the dotfile exists return username (string)
// 3. If the dotfile does not exist tell the user that
// 4. They have to first set their GH user name using
//  the --user command.

func GetUserName() string {
	var userName string

	dotFile := internals.GetShowMeDotFilePath()

	if _, err := os.Stat(dotFile); err == nil {

		userName = getUser(dotFile)

	} else if errors.Is(err, os.ErrNotExist) {

		fmt.Println("You haven't gave us your GH username...")
	}
	return userName

}

func ShowMeRepos(userName string) JSONData {
	url := fmt.Sprintf("https://api.github.com/users/%v/repos", userName)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	var data JSONData

	err = json.Unmarshal(body, &data.Items)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func printData(data JSONData) {
	log.Printf("Repos found: %d", len(data.Items)-1)

	const format = "%v\t%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Repo", "Stars", "Created at", "Description")
	fmt.Fprintf(tw, format, "----------", "-----", "----------", "----------")

	for _, i := range data.Items {
		desc := i.Description
		if len(desc) > 50 {
			desc = string(desc[:50]) + "..."
		}
		t, err := time.Parse(time.RFC3339, i.CreatedAt)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprintf(tw, format, i.FullName, i.StargazersCount, t.Year(), desc)
	}
	tw.Flush()
}
