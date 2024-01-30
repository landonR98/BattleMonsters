package scene

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene interface {
	// updates scene
	Update()
	// draws scene to target
	Render(target rl.RenderTexture2D)
}
