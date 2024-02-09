package maps

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelMap struct {
	texture    rl.Texture2D
	sourceRect rl.Rectangle
	destRect   rl.Rectangle
	grass      []rl.Rectangle
	player     rl.Vector2
}

func makeTileLayer(tileLayer MapLayer, source TileSetSource, tileWidth int, tileHeight int) rl.Texture2D {
	tileMap := make([][]int, tileLayer.Height)
	for i := 0; i < tileLayer.Height; i++ {
		row := make([]int, tileLayer.Width)
		for j := 0; j < tileLayer.Width; j++ {
			row[j] = tileLayer.Data[(tileLayer.Height*i)+j]
		}
		tileMap[i] = row
	}

	texture := rl.LoadRenderTexture(int32(tileWidth*tileLayer.Width), int32(tileHeight*tileLayer.Height))

	tiles := NewTileSet(source, tileWidth, tileHeight)

	rl.BeginTextureMode(texture)

	for i, row := range tileMap {
		for j, tileId := range row {
			tiles.DrawTile(tileId, j*tileHeight, i*tileWidth)
		}
	}

	rl.EndTextureMode()

	return texture.Texture
}

func makeObjectLayer(layer MapLayer) []rl.Rectangle {
	objects := make([]rl.Rectangle, 0, len(layer.Objects))

	for _, object := range layer.Objects {
		objects = append(objects, rl.NewRectangle(object.X, object.Y, float32(object.Width), float32(object.Height)))
	}

	return objects
}

func NewLevelMap(tiledMap *TiledMap) *LevelMap {

	if len(tiledMap.Layers) == 0 {
		return nil
	}
	var texture rl.Texture2D
	var grass []rl.Rectangle
	var player rl.Vector2

	for _, layer := range tiledMap.Layers {
		switch layer.Name {
		case "Tile Layer 1":
			{
				texture = makeTileLayer(layer, tiledMap.TileSets[0], tiledMap.TileWidth, tiledMap.TileHeight)
			}
		case "grass":
			{
				grass = makeObjectLayer(layer)
			}
		case "player":
			{
				player = rl.NewVector2(layer.Objects[0].X, layer.Objects[0].Y)
				fmt.Println("player", player)
			}
		}
	}

	return &LevelMap{
		texture:    texture,
		sourceRect: rl.NewRectangle(0, 0, float32(texture.Width), -float32(texture.Height)),
		destRect:   rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
		grass:      grass,
		player:     player,
	}
}

func (levelMap *LevelMap) Move(x float32, y float32) {
	levelMap.sourceRect.X = x
	levelMap.sourceRect.Y = y + (levelMap.sourceRect.Height)
}

func (levelMap *LevelMap) Redraw() {
	rl.DrawTexturePro(levelMap.texture, levelMap.sourceRect, levelMap.destRect, rl.NewVector2(0, 0), 0, rl.White)

	// for _, patch := range levelMap.grass {
	// 	rl.DrawRectangleRec(patch, rl.NewColor(255, 0, 0, 100))
	// }
}

func (levelMap *LevelMap) CopyRenderTexture() rl.RenderTexture2D {
	texture := rl.LoadRenderTexture(levelMap.texture.Width, levelMap.texture.Height)
	rl.BeginTextureMode(texture)

	rl.DrawTexture(levelMap.texture, 0, 0, rl.White)

	rl.EndTextureMode()

	return texture
}

func (lm *LevelMap) GetPlayerSpawnPos() rl.Vector2 {
	return lm.player
}

func (lm *LevelMap) CheckGrassCollision(player rl.Rectangle) bool {
	for _, patch := range lm.grass {
		if rl.CheckCollisionRecs(player, patch) {
			return true
		}
	}
	return false
}
