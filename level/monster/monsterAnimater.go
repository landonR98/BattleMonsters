package monster

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MonsterAnimator interface {
	Tick() (x float32, y float32, finished bool)
}

type damageAnimation struct {
	frameCount int
	startPos   rl.Vector2
}

func (a *damageAnimation) Tick() (x float32, y float32, finished bool) {
	change := float32((math.Exp2(float64(a.frameCount-30)) / 22.5) + 40)

	if a.frameCount > 60 {
		a.frameCount = 0
		return a.startPos.X, a.startPos.Y, true
	} else {
		a.frameCount++
		return a.startPos.X + change, a.startPos.Y + change, false
	}
}

func NewDamageAnimator(startPos rl.Vector2) damageAnimation {
	return damageAnimation{frameCount: 0, startPos: startPos}
}
