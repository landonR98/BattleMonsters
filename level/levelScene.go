package level

import (
	"battleMonsters/level/maps"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelScene struct {
	levelMap *maps.LevelMap

	position rl.Vector2
}

func NewLevelScene(levelMap *maps.LevelMap) *LevelScene {
	startMenu := LevelScene{levelMap: levelMap, position: rl.NewVector2(0, 0)}
	return &startMenu
}

func (levelScene *LevelScene) Update() {

	velocity := rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyUp) {
		velocity.Y += 1
	} else if rl.IsKeyDown(rl.KeyDown) {
		velocity.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		velocity.X -= 1
	} else if rl.IsKeyDown(rl.KeyRight) {
		velocity.X += 1
	}

	velocity = rl.Vector2ClampValue(velocity, 0, 1)
	levelScene.position = rl.Vector2Add(velocity, levelScene.position)
	levelScene.levelMap.Move(levelScene.position.X, levelScene.position.Y)
	fmt.Println(levelScene.position)
}

func (levelScene *LevelScene) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.RayWhite)

	levelScene.levelMap.Draw()

	rl.EndTextureMode()
}
