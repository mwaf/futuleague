package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGameRoot(t *testing.T) {
	createTestDB()
	defer removeTestDB()
	var root Root
	getAndUnmarshalJson(t, "/root", &root)

	expected := []string{"FIFA14", "NHL14"}
	assertIntEquals(t, len(expected), len(root.Games))
	for i, g := range root.Games {
		assertEquals(t, expected[i], g)
	}

}

func TestFIFAClubs(t *testing.T) {
	var game Game
	getAndUnmarshalJson(t, "/FIFA14", &game)

	passed := false
	for _, club := range game.Clubs {
		if club.Name == "PSG" && club.League == "Ligue 1" && club.Country == "France" && club.Stars == 5 {
			passed = true
			break
		}
	}
	if !passed {
		t.Errorf("Could not find PSG from club list.")
	}
}

func TestCreateGetPlayers(t *testing.T) {
	form := url.Values{}
	form.Set("name", "Jon")
	form.Set("identifier", "jkro")

	ts := httptest.NewServer(defineRoutes())
	res, err := http.PostForm(ts.URL+"/players", form)
	if err != nil {
		t.Errorf("Posting failed ", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Wrong response status")
	}
	loc, err := res.Location()
	if err != nil {
		t.Errorf("No location redirect found")
		return
	}
	res, err = http.Get(loc.String())
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Unable to read body contets")
	}
	var player Player
	unmarshalJson(t, content, player)

	expected := Player{"jkro", "Jon", 5}
	assertEquals(t, expected, player)
}
