package player

import (
	"battleMonsters/level/monster"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	sprites    rl.Texture2D
	srcRect    rl.Rectangle
	destRect   rl.Rectangle
	pos        rl.Vector2
	frameCount int
	isMoving   bool
	IsDead     bool
	Monsters   []monster.Monster
}

func NewPlayer(pos rl.Vector2, charSprites string, monsters []monster.Monster) (player Player) {
	img := rl.LoadImage("./resources/characters/" + charSprites)
	sprites := rl.LoadTextureFromImage(img)

	player.sprites = sprites
	player.srcRect = rl.NewRectangle(0, 0, 16, 16)
	player.destRect = rl.NewRectangle(0, 0, 16, 16)
	player.pos = pos
	player.frameCount = 0
	player.Monsters = monsters
	player.IsDead = false

	return player
}

func (p *Player) Draw() {
	rl.DrawTexturePro(p.sprites, p.srcRect, p.destRect, rl.NewVector2(0, 0), 0, rl.White)
}

func (p *Player) GetPos() rl.Vector2 {
	return p.pos
}

func (p *Player) GetHitBox() rl.Rectangle {
	return rl.NewRectangle(p.destRect.X+3, p.destRect.Y, p.destRect.Width-6, p.destRect.Height)
}

func (p *Player) IsMoving() bool {
	return p.isMoving
}

func (p *Player) Update() {

	velocity := rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyUp) {
		velocity.Y -= 1
		p.srcRect.X = 0
	} else if rl.IsKeyDown(rl.KeyDown) {
		velocity.Y += 1
		p.srcRect.X = 32
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		velocity.X -= 1
		p.srcRect.X = 48
	} else if rl.IsKeyDown(rl.KeyRight) {
		velocity.X += 1
		p.srcRect.X = 16
	}

	velocity = rl.Vector2ClampValue(velocity, 0, 1)
	p.pos = rl.Vector2Add(velocity, p.pos)

	p.destRect.X = p.pos.X - (p.srcRect.Width / float32(2))
	p.destRect.Y = p.pos.Y + (p.srcRect.Height / float32(2))

	if velocity.X == 0 && velocity.Y == 0 {
		p.frameCount = 0
		p.isMoving = false
	} else {
		p.frameCount++
		p.isMoving = true
	}

	animationFrame := (p.frameCount % 40) / 10

	switch animationFrame {
	case 0, 2:
		p.srcRect.Y = 16
	case 1:
		p.srcRect.Y = 0
	case 3:
		p.srcRect.Y = 32
	}

}

func (p *Player) SetPos(pos rl.Vector2) {
	p.pos = pos
}
