package main

import (
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
