package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestShowMe(t *testing.T) {

	t.Run("Get user repos", func(t *testing.T) {
		res, _ := http.Get("https://api.github.com/users/faceyacc/repos")
		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		var data JSONData
		json.Unmarshal(body, &data.Items)

		got := ShowMeRepos()
		expected := data

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, but expected, %v", got, expected)
		}
	})

}
