package window

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	GameHeight = float32(720.0)
	GameWidth  = float32(1280.0)
)

type Window struct {
	target     rl.RenderTexture2D
	targetRect rl.Rectangle
	resizeRect rl.Rectangle
	targetVec  rl.Vector2
	mouse      rl.Vector2
}

var window Window

func OpenWindow() {
	width := int32(GameWidth)
	height := int32(GameHeight)
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.InitWindow(width, height, "Battle Monsters")
	rl.SetWindowMinSize(320, 240)
	window = Window{
		target:     rl.LoadRenderTexture(width, height),
		targetRect: rl.NewRectangle(0, 0, GameWidth, -GameHeight),
		resizeRect: rl.NewRectangle(0, 0, 0, 0),
		targetVec:  rl.NewVector2(0, 0),
		mouse:      rl.NewVector2(0, 0),
	}
	rl.SetTextureFilter(window.target.Texture, rl.FilterBilinear)
}

func GetWindow() *Window {
	return &window
}

func (window *Window) Update() {
	screenHeight := rl.GetScreenHeight()
	screenWidth := rl.GetScreenWidth()

	scale := minFloat(float32(screenWidth)/GameWidth, float32(screenHeight)/GameHeight)

	// scale mouse
	mouse := rl.GetMousePosition()
	window.mouse.X = (mouse.X - (float32(screenWidth)-(GameWidth*scale))*0.5) / scale
	window.mouse.Y = (mouse.Y - (float32(screenHeight)-(GameHeight*scale))*0.5) / scale
	window.mouse = rl.Vector2Clamp(window.mouse, rl.NewVector2(0, 0), rl.NewVector2(GameWidth, GameHeight))
	// set offset for raygui
	rl.SetMouseOffset(int(-(float32(screenWidth)-(GameWidth*scale))*0.5), int(-(float32(screenHeight)-(GameHeight*scale))*0.5))
	rl.SetMouseScale(1/scale, 1/scale)

	// scale game window
	window.resizeRect.X = (float32(screenWidth) - (GameWidth * scale)) * 0.5
	window.resizeRect.Y = (float32(screenHeight) - (GameHeight * scale)) * 0.5
	window.resizeRect.Width = GameWidth * scale
	window.resizeRect.Height = GameHeight * scale
}

func (window *Window) Draw() {
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(window.target.Texture, window.targetRect, window.resizeRect, window.targetVec, 0.0, rl.White)
}

func (window *Window) GetTarget() rl.RenderTexture2D {
	return window.target
}

func minFloat(a float32, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}
