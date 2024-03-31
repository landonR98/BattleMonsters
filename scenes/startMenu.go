package scenes

import (
	"battleMonsters/window"
	"fmt"
	"net/http"
	"net/url"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)

type StartMenu struct {
	startRect       rl.Rectangle
	multiplayerRect rl.Rectangle
	title           rl.RenderTexture2D
	titleRect       rl.Rectangle
	titleDest       rl.Rectangle
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
		startRect:       rl.NewRectangle(btnX, 300, btnWidth, btnHeight),
		multiplayerRect: rl.NewRectangle(btnX, 400, btnWidth, btnHeight),
		title:           title,
		titleRect:       titleRect,
		titleDest:       titleDest,
	}
	return &startMenu
}

func (startMenu *StartMenu) Update() {

}

func (startMenu *StartMenu) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.RayWhite)

	rl.DrawTexturePro(startMenu.title.Texture, startMenu.titleRect, startMenu.titleDest, rl.NewVector2(0, 0), 0, rl.White)

	if gui.Button(startMenu.startRect, "Start Game") {
		start()
	} else if gui.Button(startMenu.multiplayerRect, "Multiplayer") {
		startMultiplayer()
	}

	rl.EndTextureMode()
}

func start() {
	levelScene, err := NewLevelScene("resources/levels/level1.json", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	GetManager().Push(&levelScene)
}

func startMultiplayer() {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusSwitchingProtocols {
		fmt.Println("Connection Error: Bad Status", resp.StatusCode)
	}

	levelScene, err := NewMultiplayerLevelScene("resources/levels/level1.json", nil, conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	GetManager().Push(&levelScene)
}
