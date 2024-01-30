package maps

import rl "github.com/gen2brain/raylib-go/raylib"

type TileSet struct {
	tilesWide  int
	tileHeight int
	tileWidth  int
	texture    rl.Texture2D
	sourceRect rl.Rectangle
	destRect   rl.Rectangle
}

func NewTileSet() *TileSet {
	image := rl.LoadImage("./resources/tileSets/TileSet.png")
	texture := rl.LoadTextureFromImage(image)
	return &TileSet{
		tilesWide:  int(texture.Width) / 16,
		tileHeight: 16,
		tileWidth:  16,
		texture:    texture,
		sourceRect: rl.NewRectangle(0, 0, 16, 16),
		destRect:   rl.NewRectangle(0, 0, 16, 16),
	}
}

func (tileSet *TileSet) TileDimensions() (width int, height int) {
	width = tileSet.tileWidth
	height = tileSet.tileHeight
	return width, height
}

func (tileSet *TileSet) DrawTile(tileID int, x int, y int) {
	tileX := tileID % tileSet.tilesWide * tileSet.tileWidth
	tileY := (tileID / tileSet.tilesWide) * tileSet.tileHeight

	tileSet.sourceRect.X = float32(tileX)
	tileSet.sourceRect.Y = float32(tileY)

	tileSet.destRect.X = float32(x)
	tileSet.destRect.Y = float32(y)

	rl.DrawTexturePro(tileSet.texture, tileSet.sourceRect, tileSet.destRect, rl.NewVector2(0, 0), 0, rl.White)
}
