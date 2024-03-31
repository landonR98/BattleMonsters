package maps

import (
	"encoding/json"
	"io"
	"os"
)

type MapLayer struct {
	Data    []int `json:"data"`
	Objects []struct {
		Height float32 `json:"height"`
		Width  float32 `json:"width"`
		X      float32 `json:"x"`
		Y      float32 `json:"y"`
	} `json:"objects"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type TileSetSource struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

type TiledMap struct {
	Height     int             `json:"height"`
	Width      int             `json:"width"`
	Layers     []MapLayer      `json:"layers"`
	TileHeight int             `json:"tileheight"`
	TileWidth  int             `json:"tilewidth"`
	TileSets   []TileSetSource `json:"tilesets"`
}

func LoadMapFromFile(filepath string) (*LevelMap, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var mapInfo TiledMap

	err = json.Unmarshal(bytes, &mapInfo)
	if err != nil {
		return nil, err
	}

	levelMap := NewLevelMap(&mapInfo)

	// for _, row := range levelMap.tileMap {
	// 	fmt.Println(row)
	// }

	return levelMap, nil
}
