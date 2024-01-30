package maps

import (
	"encoding/json"
	"io"
	"os"
)

type TiledMap struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Layers []struct {
		Data []int `json:"data"`
	} `json:"layers"`
	TileHeight int `json:"tileheight"`
	TileWidth  int `json:"tilewidth"`
	TileSets   []struct {
		FirstGID int    `json:"firstgid"`
		Source   string `json:"source"`
	} `json:"tilesets"`
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
