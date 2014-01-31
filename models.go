package main

import (
	"container/list"
	"log"
)

const (
	DEFAULT_PLAYER_RATING       = 5.0
	PLAYER_ALREADY_EXISTS_ERROR = "column identifier is not unique"
)

type RootError struct {
	Error JsonError `json:"error"`
}
type JsonError struct {
	UserMessage string `json:"msg"`
	TechMessage string `json:"tech_msg"`
}

type Club struct {
	Name    string  `json:"name"`
	League  string  `json:"league"`
	Country string  `json:"country"`
	Stars   float64 `json:"stars"`
}

func (g Club) FetchAll() ([]Club, error) {
	rows, err := DB.Query("select name, league, country, stars from clubs")
	if err != nil {
		return []Club{}, err
	}

	result := list.New()
	for rows.Next() {
		var club Club
		if err := rows.Scan(&club.Name, &club.League, &club.Country, &club.Stars); err != nil {
			log.Println(err)
			return []Club{}, err
		} else {
			result.PushBack(club)
		}

	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return []Club{}, err
	}

	clubs := make([]Club, result.Len())
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		clubs[i] = e.Value.(Club)
	}
	return clubs, nil
}

type Player struct {
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Rating     float64 `json:"rating"`
	Played     int     `json:"player"`
	Wins       int     `json:"wins"`
	Losses     int     `json:"losses"`
	Ties       int     `json:"ties"`
}

func (p Player) Save() error {
	_, err := DB.Exec(`insert into players (identifier, name, rating, played, wins, losses, ties) values (?, ?, ?, ?, ?, ?, ?);`,
		p.Identifier, p.Name, p.Rating, p.Played, p.Wins, p.Losses, p.Ties)
	return err
}

func (p Player) FetchAll() ([]Player, error) {
	rows, err := DB.Query(`select identifier, name, rating from players;`)
	if err != nil {
		return []Player{}, err
	}

	result := list.New()
	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.Identifier, &p.Name, &p.Rating); err != nil {
			log.Println(err)
			return []Player{}, err
		} else {
			result.PushBack(p)
		}
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return []Player{}, err
	}

	players := make([]Player, result.Len())
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		players[i] = e.Value.(Player)
	}

	return players, nil
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
