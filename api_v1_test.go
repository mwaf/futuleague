package main

import (
	"testing"
)

func TestGameRoot(t *testing.T) {
	createTestDB()
	defer removeTestDB()
	var games []string
	getAndUnmarshalJson(t, "/root", &games)

	expected := []string{"FIFA14", "NHL14"}
	assertIntEquals(t, len(expected), len(games))
	for i, g := range games {
		assertEquals(t, expected[i], g)
	}

}

func TestFIFAClubs(t *testing.T) {
	var game Game
	getAndUnmarshalJson(t, "/FIFA14", &game)

	passed := false
	for _, club := range game.Clubs {
		if club.Name == "Paris SG" && club.League == "Ligue 1" && club.Stars == 5 {
			passed = true
			break
		}
	}
	if !passed {
		t.Errorf("Could not find PSG from club list.")
	}
}
