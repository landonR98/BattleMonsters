package maps

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileSet struct {
	tilesWide  int
	tileHeight int
	tileWidth  int
	texture    rl.Texture2D
	sourceRect rl.Rectangle
	destRect   rl.Rectangle
	offset     int
}

func NewTileSet(source TileSetSource, tileWidth int, tileHeight int) *TileSet {
	img := strings.Split(source.Source, ".")[0] + ".png"

	image := rl.LoadImage("./resources/tileSets/" + img)
	texture := rl.LoadTextureFromImage(image)
	return &TileSet{
		tilesWide:  int(texture.Width) / tileWidth,
		tileHeight: tileHeight,
		tileWidth:  tileWidth,
		texture:    texture,
		sourceRect: rl.NewRectangle(0, 0, float32(tileWidth), float32(tileHeight)),
		destRect:   rl.NewRectangle(0, 0, float32(tileWidth), float32(tileHeight)),
		offset:     source.FirstGID,
	}
}

func (tileSet *TileSet) TileDimensions() (width int, height int) {
	width = tileSet.tileWidth
	height = tileSet.tileHeight
	return width, height
}

func (tileSet *TileSet) DrawTile(tileID int, x int, y int) {
	tileID -= tileSet.offset

	tileX := tileID % tileSet.tilesWide * tileSet.tileWidth
	tileY := (tileID / tileSet.tilesWide) * tileSet.tileHeight

	tileSet.sourceRect.X = float32(tileX)
	tileSet.sourceRect.Y = float32(tileY)

	tileSet.destRect.X = float32(x)
	tileSet.destRect.Y = float32(y)

	rl.DrawTexturePro(tileSet.texture, tileSet.sourceRect, tileSet.destRect, rl.NewVector2(0, 0), 0, rl.White)
}
