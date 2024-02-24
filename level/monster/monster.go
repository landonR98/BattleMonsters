package monster

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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
		FrameCount: 0,
	}
}

func (m Monster) Update() {

}

func (m Monster) Draw() {

}
