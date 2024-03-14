package battle

import (
	"battleMonsters/level/monster"
	"battleMonsters/level/player"
	"battleMonsters/scene"
	"battleMonsters/window"
	"math/rand"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BattleScene struct {
	player      *player.Player
	monsters    []monster.Monster
	menuBox     rl.Rectangle
	runBtn      rl.Rectangle
	catchBtn    rl.Rectangle
	fightBtn    rl.Rectangle
	monstersBtn rl.Rectangle
}

func NewBattleScene(player *player.Player, monsters []monster.Monster) *BattleScene {
	monsterSlice := []monster.Monster{monsters[rand.Intn(len(monsters))]}

	btnWidth := float32(100)
	btnHeight := float32(30)

	menuBox := rl.NewRectangle(window.GameWidth-350, window.GameHeight-150, 350, 150)
	runBtn := rl.NewRectangle(window.GameWidth-300, window.GameHeight-55, btnWidth, btnHeight)
	catchBtn := rl.NewRectangle(window.GameWidth-300, window.GameHeight-125, btnWidth, btnHeight)
	fightBtn := rl.NewRectangle(window.GameWidth-150, window.GameHeight-55, btnWidth, btnHeight)
	monstersBtn := rl.NewRectangle(window.GameWidth-150, window.GameHeight-125, btnWidth, btnHeight)

	return &BattleScene{
		player:      player,
		monsters:    monsterSlice,
		menuBox:     menuBox,
		runBtn:      runBtn,
		catchBtn:    catchBtn,
		fightBtn:    fightBtn,
		monstersBtn: monstersBtn,
	}
}

func (b *BattleScene) Update() {

}

func (b *BattleScene) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.White)

	rl.DrawRectangleLinesEx(b.menuBox, 5, rl.Gray)

	if gui.Button(b.runBtn, "run") {
		scene.GetManager().Pop()
	}

	if gui.Button(b.catchBtn, "catch") {
		scene.GetManager().Pop()
	}

	if gui.Button(b.fightBtn, "fight") {
		scene.GetManager().Pop()
	}

	if gui.Button(b.monstersBtn, "monsters") {
		scene.GetManager().Pop()
	}

	rl.EndTextureMode()
}
