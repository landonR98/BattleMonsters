package main

import (
	"battleMonsters/scenes"
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

	scenes.Init(scenes.NewStartMenu())

	sceneManager := scenes.GetManager()

	for !rl.WindowShouldClose() {
		sceneManager.GetCurrent().Update()
		windowManager.Update()

		target := windowManager.GetTarget()

		sceneManager.GetCurrent().Render(target)

		rl.BeginDrawing()
		windowManager.Draw()
		rl.EndDrawing()
	}
}
