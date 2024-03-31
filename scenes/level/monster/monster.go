package monster

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type direction int

const (
	DIRECTION_UP direction = iota
	DIRECTION_DOWN
)

type monsterSprites struct {
	Sheet      rl.Texture2D
	SourceRect rl.Rectangle
	TargetRect rl.Rectangle
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
	animation  MonsterAnimator
}

func NewMonster(m monsterModel, moves []Move) Monster {

	monsterMoves := make([]Move, 0, len(m.Moves))
	for _, move_id := range m.Moves {
		monsterMoves = append(monsterMoves, moves[move_id])
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
		Moves:      monsterMoves,
		Sprites: monsterSprites{
			Sheet:      spriteSheet.Texture,
			SourceRect: rl.NewRectangle(0, 0, 16, -16),
		},
		animation: nil,
	}
}

func (m *Monster) Update() (hasAnimation bool) {
	if m.animation != nil {
		x, y, finished := m.animation.Tick()
		m.Sprites.TargetRect.X = x
		m.Sprites.TargetRect.Y = y

		if finished {
			m.animation = nil
			return false
		} else {
			return true
		}
	}
	return false
}

func (m Monster) Draw() {
	rl.DrawTexturePro(m.Sprites.Sheet, m.Sprites.SourceRect, m.Sprites.TargetRect, rl.NewVector2(0, 0), 0, rl.White)
}

func (m *Monster) SetPosition(pos rl.Vector2, width float32, height float32) {
	m.Sprites.TargetRect = rl.NewRectangle(pos.X, pos.Y, width, height)
}

func (m *Monster) SetDirection(dir direction) {
	switch dir {
	case DIRECTION_UP:
		m.Sprites.SourceRect.Y = 0
	case DIRECTION_DOWN:
		m.Sprites.SourceRect.Y = 16
	}
}

func (m *Monster) SetAnimator(animation MonsterAnimator) {
	m.animation = animation
}

func (m Monster) GetPosition() rl.Vector2 {
	return rl.NewVector2(m.Sprites.TargetRect.X, m.Sprites.TargetRect.Y)
}

func (m *Monster) TakeDamage(damage int) {
	m.Health -= damage
	if m.Health < 0 {
		m.Health = 0
	}

}
