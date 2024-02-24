package monster

import (
	"encoding/json"
	"io"
	"os"
)

type moveModel struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Target       string  `json:"target"`
	Damage       float32 `json:"damage"`
	Effect       string  `json:"effect"`
	EffectChance float32 `json:"effectChance"`
	Cost         int     `json:"cost"`
}

type monsterModel struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sprites struct {
		Sheet    string `json:"sheet"`
		Position struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"position"`
	} `json:"sprites"`
	Attack  int   `json:"attack"`
	Health  int   `json:"health"`
	Defense int   `json:"defense"`
	Stamina int   `json:"stamina"`
	Moves   []int `json:"moves"`
}

func LoadMonsters(filepath string) ([]Move, []Monster, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	monsterData := struct {
		Monsters []monsterModel `json:"monsters"`
		Moves    []moveModel    `json:"moves"`
	}{}
	err = json.Unmarshal(bytes, &monsterData)
	if err != nil {
		return nil, nil, err
	}

	moves := make([]Move, 0, len(monsterData.Moves))
	monsters := make([]Monster, 0, len(monsterData.Monsters))

	for _, move := range monsterData.Moves {
		moves = append(moves, NewMove(move))
	}

	for _, monster := range monsterData.Monsters {
		monsters = append(monsters, NewMonster(monster, moves))
	}

	return moves, monsters, nil
}
