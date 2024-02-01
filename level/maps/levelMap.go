package maps

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelMap struct {
	texture    rl.Texture2D
	sourceRect rl.Rectangle
	destRect   rl.Rectangle
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
		}
	}

	rl.EndTextureMode()

	return &LevelMap{
		texture:    texture.Texture,
		sourceRect: rl.NewRectangle(0, 0, float32(texture.Texture.Width), -float32(texture.Texture.Height)),
		destRect:   rl.NewRectangle(0, 0, float32(texture.Texture.Width), float32(texture.Texture.Height)),
	}
}

func (levelMap *LevelMap) Move(x float32, y float32) {
	levelMap.sourceRect.X = x
	levelMap.sourceRect.Y = y + (levelMap.sourceRect.Height)
}

func (levelMap *LevelMap) Redraw() {
	rl.DrawTexturePro(levelMap.texture, levelMap.sourceRect, levelMap.destRect, rl.NewVector2(0, 0), 0, rl.White)
}

func (levelMap *LevelMap) CopyRenderTexture() rl.RenderTexture2D {
	texture := rl.LoadRenderTexture(levelMap.texture.Width, levelMap.texture.Height)
	rl.BeginTextureMode(texture)

	rl.DrawTexture(levelMap.texture, 0, 0, rl.White)

	rl.EndTextureMode()

	return texture
}
