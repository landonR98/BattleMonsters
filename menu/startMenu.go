package menu

import (
	"battleMonsters/level"
	"battleMonsters/level/maps"
	"battleMonsters/scene"
	"battleMonsters/window"
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type StartMenu struct {
	startRect    rl.Rectangle
	settingsRect rl.Rectangle
	title        rl.RenderTexture2D
	titleRect    rl.Rectangle
	titleDest    rl.Rectangle
}

func NewStartMenu() *StartMenu {
	halfScreen := float32((window.GameWidth / 2))

	titleWidth := int32(250)
	title := rl.LoadRenderTexture(titleWidth, 40)

	rl.BeginTextureMode(title)
	// rl.ClearBackground(rl.LightGray)
	// rl.DrawRectangleLines(0, 0, titleWidth, 40, rl.Gray)
	rl.DrawText("Battle Monsters", 10, 0, 29, rl.Black)
	rl.EndTextureMode()

	titleRect := rl.NewRectangle(0, 0, float32(title.Texture.Width), -float32(title.Texture.Height))
	titleDest := rl.NewRectangle(halfScreen-float32(title.Texture.Width/2), 50, float32(titleWidth), 40)

	var btnWidth float32 = 200
	var btnHeight float32 = 80
	btnX := halfScreen - (btnWidth / 2)

	startMenu := StartMenu{
		startRect:    rl.NewRectangle(btnX, 300, btnWidth, btnHeight),
		settingsRect: rl.NewRectangle(btnX, 400, btnWidth, btnHeight),
		title:        title,
		titleRect:    titleRect,
		titleDest:    titleDest,
	}
	return &startMenu
}

func (startMenu *StartMenu) Update() {

}

func (startMenu *StartMenu) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.RayWhite)

	rl.DrawTexturePro(startMenu.title.Texture, startMenu.titleRect, startMenu.titleDest, rl.NewVector2(0, 0), 0, rl.White)

	if gui.Button(startMenu.startRect, "Start") {
		fmt.Println("btn press")
		// scene.GetManager().Swap(level.NewLevelScene())
		start()
	} else if gui.Button(startMenu.settingsRect, "Settings") {
		fmt.Println("settings")
	}

	rl.EndTextureMode()
}

func start() {
	levelMap, err := maps.LoadMapFromFile("./resources/maps/map2.tmj")
	if err != nil {
		fmt.Println(err)
		return
	}
	levelScene := level.NewLevelScene(levelMap)
	scene.GetManager().Swap(levelScene)
}
