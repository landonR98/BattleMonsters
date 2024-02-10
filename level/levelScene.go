package level

import (
	"battleMonsters/level/battle"
	"battleMonsters/level/maps"
	"battleMonsters/scene"
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelScene struct {
	levelMap *maps.LevelMap
	canvas   rl.RenderTexture2D
	camera   Camera

	player Player
}

func NewLevelScene(levelMap *maps.LevelMap) *LevelScene {
	levelScene := LevelScene{levelMap: levelMap, camera: NewCamera(), player: NewPlayer(levelMap.GetPlayerSpawnPos(), "redChar.png")}
	levelScene.canvas = levelScene.levelMap.CopyRenderTexture()
	return &levelScene
}

func (levelScene *LevelScene) Update() {

	levelScene.player.Update()

	levelScene.camera.KeepInRect(levelScene.player.GetPos())

	if levelScene.player.IsMoving() && levelScene.levelMap.CheckGrassCollision(levelScene.player.GetHitBox()) {

		if rand.Int31()%100 == 1 {
			scene.GetManager().Push(battle.NewBattleScene())
			fmt.Println("encounter")
		}
	}

}

func (levelScene *LevelScene) Render(target rl.RenderTexture2D) {

	// draw scene to canvas
	rl.BeginTextureMode(levelScene.canvas)

	levelScene.levelMap.Redraw()

	levelScene.player.Draw()

	rl.EndTextureMode()

	// draw part of canvas to screen
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.RayWhite)

	levelScene.camera.Draw(levelScene.canvas.Texture, levelScene.camera.pos)

	rl.EndTextureMode()
}
