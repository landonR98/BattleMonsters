package level

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	texture  rl.Texture2D
	srcRect  rl.Rectangle
	destRect rl.Rectangle
	pos      rl.Vector2
}

func NewPlayer() (player Player) {
	renderTexture := rl.LoadRenderTexture(64, 16)

	rl.BeginTextureMode(renderTexture)

	rl.DrawRectangle(0, 0, 64, 16, rl.Red)

	rl.EndTextureMode()

	player.texture = renderTexture.Texture
	player.srcRect = rl.NewRectangle(0, 0, 16, -16)
	player.destRect = rl.NewRectangle(0, 0, 16, 16)
	player.pos = rl.NewVector2(150, 150)

	return player
}

func (p *Player) Draw() {
	rl.DrawTexturePro(p.texture, p.srcRect, p.destRect, rl.NewVector2(0, 0), 0, rl.White)
}

func (player *Player) GetPos() rl.Vector2 {
	return player.pos
}

func (player *Player) Update() {

	velocity := rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyUp) {
		velocity.Y -= 1
	} else if rl.IsKeyDown(rl.KeyDown) {
		velocity.Y += 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		velocity.X -= 1
	} else if rl.IsKeyDown(rl.KeyRight) {
		velocity.X += 1
	}

	velocity = rl.Vector2ClampValue(velocity, 0, 1)
	player.pos = rl.Vector2Add(velocity, player.pos)

	player.destRect.X = player.pos.X - (player.srcRect.Width / float32(2))
	player.destRect.Y = player.pos.Y + (player.srcRect.Height / float32(2))

}
