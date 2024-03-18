package textDisplay

import (
	"battleMonsters/scene"
	"battleMonsters/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type textDisplayTransition struct {
	frameCount           int
	length               int
	displayText          string
	positionX, positionY int32
	textSize             int32
}

func (t *textDisplayTransition) Update() {
	t.frameCount++

	if t.frameCount >= t.length {
		scene.GetManager().Pop()
	}
}

func (t *textDisplayTransition) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.White)

	rl.DrawText(t.displayText, t.positionX, t.positionY, t.textSize, rl.Black)

	rl.EndTextureMode()
}

func NewTextDisplayTransition(text string, size int32, length int) *textDisplayTransition {
	textSize := rl.MeasureText(text, size)
	posX := (int32(window.GameWidth) - textSize) / 2
	posY := (int32(window.GameHeight) - size) / 2

	return &textDisplayTransition{
		frameCount:  0,
		length:      length,
		displayText: text,
		positionX:   posX,
		positionY:   posY,
		textSize:    size,
	}
}
