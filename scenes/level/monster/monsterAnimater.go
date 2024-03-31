package monster

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MonsterAnimator interface {
	Tick() (x float32, y float32, finished bool)
}

type topRightAnimation struct {
	frameCount     int
	startPos       rl.Vector2
	xScale, yScale float32
}

func (a *topRightAnimation) Tick() (x float32, y float32, finished bool) {
	change := float32((math.Pow(float64(a.frameCount-30), 2) / -22.5) + 40)

	if a.frameCount > 60 {
		a.frameCount = 0
		return a.startPos.X, a.startPos.Y, true
	} else {
		a.frameCount++
		return a.startPos.X + (change * a.xScale), a.startPos.Y + (change * a.yScale), false
	}
}

func NewBottomRightAnimator(startPos rl.Vector2, xScale float32, yScale float32) topRightAnimation {
	return topRightAnimation{frameCount: 0, startPos: startPos, xScale: xScale, yScale: yScale}
}
