package main

import (
	"container/list"
	"errors"
	"log"
)

const (
	DEFAULT_PLAYER_RATING       = 5
	PLAYER_ALREADY_EXISTS_ERROR = "column identifier is not unique"
)

type Root struct {
	Games []string `json:"games"`
}

type Game struct {
	Name  string `json:"name"`
	Clubs []Club `json:"clubs"`
}

func (g Game) FetchAll() []Game {
	rows, err := DB.Query("select name from games")
	if err != nil {
		log.Fatal(err)
	}

	result := list.New()
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.Name); err != nil {
			log.Fatal(err)
		} else {
			result.PushBack(game)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	games := make([]Game, result.Len())
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		games[i] = e.Value.(Game)
	}
	return games
}

func (g Game) FetchByName(name string) (Game, error) {
	rows, err := DB.Query("select c.name, c.league, c.country, c.stars from clubs c join games g where c.game = g.id and g.name=?", name)
	if err != nil {
		log.Fatal(err)
	}

	result := list.New()
	count := 0
	for ; rows.Next(); count++ {
		var club Club
		if err := rows.Scan(&club.Name, &club.League, &club.Country, &club.Stars); err != nil {
			log.Fatal(err)
		} else {
			result.PushBack(club)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return Game{}, errors.New("Could not find game")
	}

	clubs := make([]Club, count)
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		clubs[i] = e.Value.(Club)
	}

	return Game{Name: name, Clubs: clubs}, nil
}

type Club struct {
	Name    string  `json:"name"`
	League  string  `json:"league"`
	Country string  `json:"country"`
	Stars   float64 `json:"stars"`
}

type Player struct {
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Rating     float64 `json:"rating"`
}

func (p Player) Save() error {
	_, err := DB.Exec(`insert into players (identifier, name, rating) values (?, ?, ?);`, p.Identifier, p.Name, p.Rating)
	return err
}

func (p Player) FetchAll() ([]Player, error) {
	// todo get stuff from DB
	return []Player{}, nil
}

func (p Player) FetchByIdentifier(id string) (Player, error) {
	row := DB.QueryRow(`select identifier, name, rating from players where identifier=?;`, id)
	err := row.Scan(&p.Identifier, &p.Name, &p.Rating)

	// sorry, being lazy and not handling empty result and DB
	// failure differently (assuming empty result)
	if err != nil {
		return p, err
	}

	return p, nil
}
