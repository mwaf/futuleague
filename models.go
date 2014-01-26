package main

import (
	"container/list"
	"log"
)

type Game struct {
	Name  string `json:"name"`
	Clubs []Club `json:"clubs"`
}

func (g Game) FetchAll() []interface{} {
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

	games := make([]interface{}, result.Len())
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		games[i] = e.Value
	}
	return games
}

func (g Game) fetchByName(name string) Game {
	//rows, err = db.Query("select c.* from clubs c join games g where g.name=?", name)
	return Game{}
}

type Club struct {
	Name   string  `json:"name"`
	League string  `json:"name"`
	Stars  float64 `json:"stars"`
}
