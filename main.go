package main

import (
	"battleMonsters/menu"
	"battleMonsters/scene"
	"battleMonsters/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	defaultHeight = 720
	defaultWidth  = 1080
)

func main() {
	window.OpenWindow()
	defer rl.CloseWindow()

	windowManager := window.GetWindow()

	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	scene.Init(menu.NewStartMenu())

	sceneManager := scene.GetManager()

	for !rl.WindowShouldClose() {
		sceneManager.GetCurrent().Update()
		windowManager.Update()

		target := windowManager.GetTarget()
		rl.BeginTextureMode(target)
		sceneManager.GetCurrent().Draw(&target)
		rl.EndTextureMode()

		rl.BeginDrawing()
		windowManager.Draw()
		rl.EndDrawing()
	}
}
