package main

import (
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
	createTestDB()
	defer removeTestDB()
	ts := httptest.NewServer(defineRoutes())

	form := url.Values{}
	form.Set("name", "Jon")
	form.Set("identifier", "jkro")

	res, err := http.PostForm(ts.URL+"/players", form)
	if err != nil {
		t.Errorf("Posting failed ", err)
	}
	assertIntEquals(t, http.StatusCreated, res.StatusCode)

	loc, err := res.Location()
	if err != nil {
		t.Errorf("No location redirect found")
		return
	}
	var postPlayer Player
	unmarshalJsonFromResponse(t, res, &postPlayer)

	// Next fetch the redirect content
	res, err = http.Get(loc.String())
	if err != nil {
		t.Errorf("Fetching from redirect failed", loc.String())
	}
	var getPlayer Player
	unmarshalJsonFromResponse(t, res, &getPlayer)

	// Now do another post with same ID
	form.Set("name", "JonJon")
	res, err = http.PostForm(ts.URL+"/players", form)
	if err != nil {
		t.Errorf("Posting failed ", err)
	}
	assertIntEquals(t, http.StatusSeeOther, res.StatusCode)
	var repostPlayer Player
	unmarshalJsonFromResponse(t, res, &repostPlayer)

	expected := Player{"jkro", "Jon", 5}
	assertEquals(t, expected, postPlayer)
	assertEquals(t, expected, getPlayer)
	assertEquals(t, expected, repostPlayer)

}
