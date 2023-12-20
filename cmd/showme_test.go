package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/user"
	"reflect"
	"strings"
	"testing"

	"github.com/Flaque/filet"
)

func TestShowMe(t *testing.T) {

	checkAssert := func(t testing.TB, got, expected JSONData) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, but expected, %v", got, expected)
		}
	}

	t.Run("Get user repos", func(t *testing.T) {
		res, _ := http.Get("https://api.github.com/users/faceyacc/repos")
		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		var data JSONData
		json.Unmarshal(body, &data.Items)

		got := ShowMeRepos("faceyacc")
		expected := data

		checkAssert(t, got, expected)
	})
}

func TestGetUser(t *testing.T) {

	checkAssert := func(t testing.TB, got, expected string) {
		t.Helper()
		if strings.TrimSpace(got) != strings.TrimSpace(expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var (
		fileName = usr.HomeDir + "/.test"
		email    = "test@gmail.com" + "\n"
		userName = "testuser" + "\n"
	)

	t.Run("file exists", func(t *testing.T) {
		fileName = usr.HomeDir + "/.test"
		email = "test@gmail.com" + "\n"
		userName = "testuser" + "\n"

		filet.File(t, fileName, email+userName)

		exist := filet.Exists(t, fileName)

		if exist == false {
			t.Fatal("File does not exist")
		}

	})

	t.Run("get email", func(t *testing.T) {
		defer filet.CleanUp(t)

		filet.File(t, fileName, email+userName)

		gotEmail, _ := getUser(fileName)
		expectedEmail := email

		checkAssert(t, gotEmail, expectedEmail)
	})

	t.Run("get username", func(t *testing.T) {
		defer filet.CleanUp(t)

		filet.File(t, fileName, email+userName)

		_, gotUsername := getUser(fileName)
		expectedUsername := userName

		checkAssert(t, gotUsername, expectedUsername)
	})

}
