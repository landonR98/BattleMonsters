package level

import (
	"battleMonsters/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	source rl.Rectangle
	target rl.Rectangle
	pos    rl.Vector2
}

func NewCamera() Camera {
	cameraScale := float32(4)
	return Camera{
		source: rl.NewRectangle(0, 0, window.GameWidth/cameraScale, -window.GameHeight/cameraScale),
		target: rl.NewRectangle(0, 0, window.GameWidth, window.GameHeight),
		pos:    rl.NewVector2(0, -64),
	}
}

func (c *Camera) Draw(texture rl.Texture2D, pos rl.Vector2) {
	c.source.X = c.pos.X
	c.source.Y = c.pos.Y
	rl.DrawTexturePro(texture, c.source, c.target, rl.NewVector2(0, 0), 0, rl.White)
}

func (c *Camera) SetPos(pos rl.Vector2) {
	c.pos.X = pos.X - (c.source.Width / 2)
	c.pos.Y = (pos.Y - (c.source.Height / 2)) * -1
}

func (c *Camera) KeepInRect(pos rl.Vector2) {
	vec := rl.NewVector2(pos.X, pos.Y*-1)
	borderWidth := c.source.Width / 3
	borderHeight := c.source.Height / 3
	if vec.X < (c.pos.X + borderWidth) {
		c.pos.X = vec.X - borderWidth
	} else if vec.X > ((c.pos.X + c.source.Width) - borderWidth) {
		c.pos.X = vec.X - c.source.Width + borderWidth
	}
	if vec.Y < (c.pos.Y - borderHeight) {
		c.pos.Y = vec.Y + borderHeight
	} else if vec.Y > (c.pos.Y - c.source.Height + borderHeight) {
		c.pos.Y = vec.Y + c.source.Height - borderHeight
	}
}
