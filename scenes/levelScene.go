package scenes

import (
	"battleMonsters/scenes/level"
	"battleMonsters/scenes/level/maps"
	"battleMonsters/scenes/level/monster"
	"battleMonsters/scenes/level/player"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)

type LevelScene struct {
	levelMap *maps.LevelMap
	canvas   rl.RenderTexture2D
	camera   level.Camera
	monsters []monster.Monster
	moves    []monster.Move
	player   player.Player
	conn     *websocket.Conn
}

func NewLevelScene(levelPath string, pl *player.Player) (LevelScene, error) {
	var levelData map[string]string

	data, err := os.ReadFile("resources/levels/level1.json")
	if err != nil {
		fmt.Println(err)
		return LevelScene{}, err
	}

	err = json.Unmarshal(data, &levelData)
	if err != nil {
		fmt.Println(err)
		return LevelScene{}, err
	}

	mapPath := levelData["map"]
	monsterPath := levelData["monsters"]

	levelMap, err := maps.LoadMapFromFile(mapPath)
	if err != nil {
		fmt.Println(err)
		return LevelScene{}, err
	}

	moves, monsters, err := monster.LoadMonsters(monsterPath)
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
		camera:   level.NewCamera(),
		player:   level_player,
		monsters: monsters,
		moves:    moves,
		conn:     nil,
	}
	levelScene.canvas = levelScene.levelMap.CopyRenderTexture()
	return levelScene, nil
}

func NewMultiplayerLevelScene(levelPath string, pl *player.Player, connection *websocket.Conn) (LevelScene, error) {
	var levelData map[string]string

	data, err := os.ReadFile("resources/levels/level1.json")
	if err != nil {
		fmt.Println(err)
		return LevelScene{}, err
	}

	err = json.Unmarshal(data, &levelData)
	if err != nil {
		fmt.Println(err)
		return LevelScene{}, err
	}

	mapPath := levelData["map"]
	monsterPath := levelData["monsters"]

	levelMap, err := maps.LoadMapFromFile(mapPath)
	if err != nil {
		fmt.Println(err)
		return LevelScene{}, err
	}

	moves, monsters, err := monster.LoadMonsters(monsterPath)
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
		camera:   level.NewCamera(),
		player:   level_player,
		monsters: monsters,
		moves:    moves,
		conn:     connection,
	}
	levelScene.canvas = levelScene.levelMap.CopyRenderTexture()
	return levelScene, nil
}

func (ls *LevelScene) Update() {
	if ls.player.IsDead {
		GetManager().Pop()
	}

	ls.player.Update()

	if ls.conn != nil {
		ls.sendLocation()
	}

	ls.camera.KeepInRect(ls.player.GetPos())

	if ls.player.IsMoving() && ls.levelMap.CheckGrassCollision(ls.player.GetHitBox()) {

		// GetManager().Push(NewBattleScene(&levelScene.player, levelScene.monsters))
		if rand.Int31()%100 == 1 {
			GetManager().Push(NewBattleScene(&ls.player, ls.monsters))
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

	levelScene.camera.Draw(levelScene.canvas.Texture, levelScene.camera.GetPos())

	rl.EndTextureMode()
}

type Event struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}

type pos struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type LocationEvent struct {
	Pos pos `json:"position"`
}

func (ls *LevelScene) sendLocation() error {
	pp := ls.player.GetPos()
	position := LocationEvent{Pos: pos{X: pp.X, Y: pp.Y}}
	eventJson, err := json.Marshal(position)
	if err != nil {
		fmt.Println(err)
	}
	event := Event{Type: "location", Payload: eventJson}
	fmt.Println(event)

	return ls.conn.WriteJSON(event)
}
