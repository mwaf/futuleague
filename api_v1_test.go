package main

import (
	"testing"
)

func TestGameRoot(t *testing.T) {
	var games []Game
	getAndUnmarshalJson(t, "/root", &games)

	expected := []Game{Game{Name: "FIFA14"}, Game{Name: "NHL14"}}
	assertIntEquals(t, len(expected), len(games))
	for i, g := range games {
		assertEquals(t, expected[i].Name, g.Name)
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
