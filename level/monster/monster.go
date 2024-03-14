package monster

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type monsterSprites struct {
	Sheet      rl.Texture2D
	SourceRect rl.Rectangle
	TargetRect rl.Rectangle
}

type animationData struct {
	frameCount   int
	startPos     rl.Vector2
	takingDamage bool
}

type Monster struct {
	Name       string
	Sprites    monsterSprites
	Attack     int
	Defense    int
	MaxHealth  int
	Health     int
	MaxStamina int
	Stamina    int
	Moves      []Move
	FrameCount int
	animation  animationData
}

func NewMonster(m monsterModel, moves []Move) Monster {

	monster_moves := make([]Move, 0, len(m.Moves))
	for _, move_id := range m.Moves {
		monster_moves = append(monster_moves, moves[move_id])
	}

	texture := rl.LoadTexture("./resources/monsters/" + m.Sprites.Sheet)
	spriteSheet := rl.LoadRenderTexture(16*4, 16*2)
	sourceRect := rl.NewRectangle(float32(m.Sprites.Position.X*16), float32(m.Sprites.Position.Y*16), float32(spriteSheet.Texture.Width), float32(spriteSheet.Texture.Height))
	rl.BeginTextureMode(spriteSheet)
	rl.DrawTextureRec(texture, sourceRect, rl.NewVector2(0, 0), rl.White)
	rl.EndTextureMode()

	return Monster{
		Name:       m.Name,
		Attack:     m.Attack,
		Defense:    m.Defense,
		MaxHealth:  m.Health,
		Health:     m.Health,
		MaxStamina: m.Stamina,
		Stamina:    m.Stamina,
		Moves:      monster_moves,
		Sprites: monsterSprites{
			Sheet:      spriteSheet.Texture,
			SourceRect: rl.NewRectangle(0, 0, 16, 16),
		},
		animation: animationData{
			frameCount:   0,
			takingDamage: false,
		},
	}
}

func (m Monster) Update() {
	if m.animation.takingDamage {
		change := float32((math.Exp2(float64(m.animation.frameCount-30)) / 22.5) + 40)
		m.Sprites.TargetRect.X = m.animation.startPos.X + change
		m.Sprites.TargetRect.Y = m.animation.startPos.Y + change

		if m.animation.frameCount > 60 {
			m.animation.takingDamage = false
			m.animation.frameCount = 0
		} else {
			m.animation.frameCount++
		}
	}
}

func (m Monster) Draw() {
	rl.DrawTexturePro(m.Sprites.Sheet, m.Sprites.SourceRect, m.Sprites.TargetRect, rl.NewVector2(0, 0), 0, rl.White)
}

func (m Monster) SetPosition(pos rl.Vector2, width float32, height float32) {
	m.Sprites.TargetRect = rl.NewRectangle(pos.X, pos.Y, width, height)
	m.animation.startPos = pos
}
