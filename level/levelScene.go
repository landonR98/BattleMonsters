package level

import (
	"battleMonsters/level/battle"
	"battleMonsters/level/maps"
	"battleMonsters/level/monster"
	"battleMonsters/level/player"
	"battleMonsters/scene"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelScene struct {
	levelMap *maps.LevelMap
	canvas   rl.RenderTexture2D
	camera   Camera
	monsters []monster.Monster
	moves    []monster.Move
	player   player.Player
}

func NewLevelScene(levelMap *maps.LevelMap, pl *player.Player) (LevelScene, error) {
	moves, monsters, err := monster.LoadMonsters("./resources/monsters/creatures.json")
	if err != nil {
		return LevelScene{}, err
	}

	var level_player player.Player
	if pl == nil {
		level_player = player.NewPlayer(levelMap.GetPlayerSpawnPos(), "redChar.png", monsters[:1])
	} else {
		level_player = *pl
	}

	levelScene := LevelScene{
		levelMap: levelMap,
		camera:   NewCamera(),
		player:   level_player,
		monsters: monsters,
		moves:    moves,
	}
	levelScene.canvas = levelScene.levelMap.CopyRenderTexture()
	return levelScene, nil
}

func (levelScene *LevelScene) Update() {

	levelScene.player.Update()

	levelScene.camera.KeepInRect(levelScene.player.GetPos())

	if levelScene.player.IsMoving() && levelScene.levelMap.CheckGrassCollision(levelScene.player.GetHitBox()) {

		// scene.GetManager().Push(battle.NewBattleScene(&levelScene.player, levelScene.monsters))
		if rand.Int31()%100 == 1 {
			scene.GetManager().Push(battle.NewBattleScene(&levelScene.player, levelScene.monsters))
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
