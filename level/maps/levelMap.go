package maps

import (
	"battleMonsters/window"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelMap struct {
	tileMap    [][]int
	texture    rl.RenderTexture2D
	sourceRect rl.Rectangle
	destRect   rl.Rectangle
	tiles      *TileSet
	cameraZoom float32
}

func (LevelMap) Foo() {
	fmt.Println("bar")
}

func NewLevelMap(tiledMap *TiledMap) *LevelMap {

	if len(tiledMap.Layers) == 0 {
		return nil
	}

	tileMap := make([][]int, tiledMap.Height)
	for i := 0; i < tiledMap.Height; i++ {
		row := make([]int, tiledMap.Width)
		for j := 0; j < tiledMap.Width; j++ {
			row[j] = tiledMap.Layers[0].Data[(tiledMap.Height*i)+j]
			// row[j] = 0
		}
		tileMap[i] = row
	}

	texture := rl.LoadRenderTexture(int32(tiledMap.TileWidth*tiledMap.Width), int32(tiledMap.TileHeight*tiledMap.Height))

	tiles := NewTileSet()

	rl.BeginTextureMode(texture)

	tileWidth, tileHeight := tiles.TileDimensions()

	for i, row := range tileMap {
		for j, tileId := range row {
			tiles.DrawTile(tileId-1, j*tileHeight, i*tileWidth)
			// _ = tileId
			// tiles.DrawTile(0, i*tileWidth, j*tileHeight)
		}
	}

	rl.EndTextureMode()

	cameraZoom := float32(5)

	return &LevelMap{
		tileMap:    tileMap,
		tiles:      tiles,
		texture:    texture,
		sourceRect: rl.NewRectangle(0, 0, window.GameWidth/cameraZoom, -window.GameHeight/cameraZoom),
		destRect:   rl.NewRectangle(0, 0, window.GameWidth, window.GameHeight),
	}
}

func (levelMap *LevelMap) Move(x float32, y float32) {
	levelMap.sourceRect.X = x
	levelMap.sourceRect.Y = y + (levelMap.sourceRect.Height)
}

func (levelMap *LevelMap) Draw() {
	rl.DrawTexturePro(levelMap.texture.Texture, levelMap.sourceRect, levelMap.destRect, rl.NewVector2(0, 0), 0, rl.White)
}
