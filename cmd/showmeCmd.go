package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Item represents a single repo data structure
type Item struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
}

type JSONData struct {
	Items []Item
}

func ShowMeRepos() JSONData {
	res, err := http.Get("https://api.github.com/users/faceyacc/repos")
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

	// var data []Item
	var data JSONData

	err = json.Unmarshal(body, &data.Items)
	// test := JSONData{data}

	// fmt.Printf("Printing body %v", string(body))
	if err != nil {
		log.Fatal(err)
	}

	return data
}
