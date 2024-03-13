package battle

import (
	"battleMonsters/level/monster"
	"battleMonsters/level/player"
	"battleMonsters/scene"
	"battleMonsters/window"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BattleScene struct {
	player   *player.Player
	monsters []monster.Monster
	menu     rl.Rectangle
	runBtn   rl.Rectangle
}

func NewBattleScene(player *player.Player, monsters []monster.Monster) *BattleScene {
	return &BattleScene{
		player:   player,
		monsters: monsters,
		menu:     rl.NewRectangle(window.GameWidth-400, window.GameHeight-200, 400, 200),
		runBtn:   rl.NewRectangle(window.GameWidth-300, window.GameHeight-100, 100, 30),
	}
}

func (b *BattleScene) Update() {

}

func (b *BattleScene) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.White)

	rl.DrawRectangleLinesEx(b.menu, 5, rl.Gray)

	if gui.Button(b.runBtn, "run") {
		scene.GetManager().Pop()
	}

	rl.EndTextureMode()
}
