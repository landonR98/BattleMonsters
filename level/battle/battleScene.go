package battle

import (
	"battleMonsters/level/monster"
	"battleMonsters/level/player"
	"battleMonsters/scene"
	"battleMonsters/transition/textDisplay"
	"battleMonsters/window"
	"fmt"
	"math/rand"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	STATE_CHOICE = iota
	STATE_CHOSE_MOVE
	STATE_ATTACKING
	STATE_OPPONENT_MOVE
	STATE_CATCHING
	STATE_RUNNING
	STATE_WIN
	STATE_LOOSE
)

type BattleScene struct {
	player        *player.Player
	monsters      []monster.Monster
	playerMonster *monster.Monster
	opponent      *monster.Monster
	menuBox       rl.Rectangle
	blBtn         rl.Rectangle
	tlBtn         rl.Rectangle
	brBtn         rl.Rectangle
	trBtn         rl.Rectangle
	state         int
	infoText      string
}

func NewBattleScene(player *player.Player, monsters []monster.Monster) *BattleScene {
	monsterSlice := []monster.Monster{monsters[rand.Intn(len(monsters))]}

	monsterSize := float32(250)

	for i := range monsterSlice {
		monsterSlice[i].SetDirection(monster.DIRECTION_DOWN)
		monsterSlice[i].SetPosition(rl.NewVector2(window.GameWidth-(monsterSize+100), 100), monsterSize, monsterSize)
	}

	for i := range player.Monsters {
		player.Monsters[i].SetDirection(monster.DIRECTION_UP)
		player.Monsters[i].SetPosition(rl.NewVector2(100, window.GameHeight-(monsterSize+50)), monsterSize, monsterSize)
	}

	btnWidth := float32(100)
	btnHeight := float32(30)

	menuBox := rl.NewRectangle(window.GameWidth-350, window.GameHeight-150, 350, 150)
	blBtn := rl.NewRectangle(window.GameWidth-300, window.GameHeight-55, btnWidth, btnHeight)
	tlBtn := rl.NewRectangle(window.GameWidth-300, window.GameHeight-125, btnWidth, btnHeight)
	brBtn := rl.NewRectangle(window.GameWidth-150, window.GameHeight-55, btnWidth, btnHeight)
	trBtn := rl.NewRectangle(window.GameWidth-150, window.GameHeight-125, btnWidth, btnHeight)

	return &BattleScene{
		player:        player,
		monsters:      monsterSlice,
		menuBox:       menuBox,
		blBtn:         blBtn,
		tlBtn:         tlBtn,
		brBtn:         brBtn,
		trBtn:         trBtn,
		playerMonster: &player.Monsters[0],
		opponent:      &monsterSlice[0],
		state:         STATE_CHOICE,
	}
}

func (b *BattleScene) Update() {
	switch b.state {
	case STATE_CHOICE:
		b.infoText = ""
	case STATE_ATTACKING:
		b.infoText = "Attacking..."
		hasAnimation := b.playerMonster.Update()
		if !hasAnimation {
			hasAnimation = b.opponent.Update()
			if !hasAnimation {
				if b.opponent.Health <= 0 {
					if len(b.monsters) == 1 {
						b.opponent = nil
						b.state = STATE_WIN
					} else {
						b.monsters = b.monsters[1:len(b.monsters)]
					}
				} else {
					b.state = STATE_OPPONENT_MOVE

					move := b.opponent.Moves[rand.Intn(len(b.opponent.Moves))]
					b.playerMonster.TakeDamage(int(move.Damage) * b.opponent.Attack)

					b.infoText = fmt.Sprintf("%s uses %s", b.opponent.Name, move.Name)

					opponentsAnimation := monster.NewBottomRightAnimator(b.opponent.GetPosition(), -5, 5)
					playerAnimation := monster.NewBottomRightAnimator(b.playerMonster.GetPosition(), -3, 3)
					b.opponent.SetAnimator(&opponentsAnimation)
					b.playerMonster.SetAnimator(&playerAnimation)
				}
			}
		}
	case STATE_OPPONENT_MOVE:
		hasAnimation := b.opponent.Update()
		if !hasAnimation {
			hasAnimation = b.playerMonster.Update()
			if !hasAnimation {
				if b.playerMonster.Health <= 0 {
					for i, monster := range b.player.Monsters {
						if monster.Health > 0 {
							b.playerMonster = &b.player.Monsters[i]
							break
						}
					}
					if b.playerMonster.Health <= 0 {
						b.state = STATE_LOOSE
					} else {
						b.state = STATE_CHOICE
					}
				} else {
					b.state = STATE_CHOICE
				}
			}
		}
	case STATE_CATCHING:
		b.infoText = "Catching..."
	case STATE_LOOSE:
		scene.GetManager().Swap(textDisplay.NewTextDisplayTransition("You Have Lost", 30, 60))
	case STATE_WIN:
		scene.GetManager().Swap(textDisplay.NewTextDisplayTransition("You Have Won", 30, 60))
	}
}

func (b *BattleScene) Render(target rl.RenderTexture2D) {
	rl.BeginTextureMode(target)

	rl.ClearBackground(rl.White)

	rl.DrawRectangleLinesEx(b.menuBox, 5, rl.Gray)

	switch b.state {
	case STATE_CHOICE:
		if gui.Button(b.blBtn, "run") {
			scene.GetManager().Pop()
		}

		if gui.Button(b.tlBtn, "catch") {
			b.state = STATE_CATCHING
		}

		if gui.Button(b.brBtn, "fight") {
			b.state = STATE_CHOSE_MOVE
		}

		if gui.Button(b.trBtn, "monsters") {
			scene.GetManager().Pop()
		}
	case STATE_CHOSE_MOVE:
		btns := []rl.Rectangle{b.tlBtn, b.blBtn, b.trBtn, b.brBtn}

		for i, move := range b.playerMonster.Moves {
			if gui.Button(btns[i], move.Name) {
				b.state = STATE_ATTACKING

				animator := monster.NewBottomRightAnimator(b.playerMonster.GetPosition(), 5, -5)
				b.playerMonster.SetAnimator(&animator)
				opponentAnimator := monster.NewBottomRightAnimator(b.opponent.GetPosition(), 3, -3)
				b.opponent.SetAnimator(&opponentAnimator)

				b.opponent.TakeDamage(int(move.Damage) * b.playerMonster.Attack)
			}
		}

	case STATE_WIN, STATE_LOOSE:
		rl.EndTextureMode()
		return
	default:
		fontSize := int32(18)
		textLength := rl.MeasureText(b.infoText, fontSize)
		textX := ((b.menuBox.ToInt32().Width - textLength) / 2) + b.menuBox.ToInt32().X
		textY := ((b.menuBox.ToInt32().Height - fontSize) / 2) + b.menuBox.ToInt32().Y
		rl.DrawText(b.infoText, textX, textY, fontSize, rl.Black)
	}

	b.opponent.Draw()
	b.playerMonster.Draw()

	drawMonsterStats(0, 0, b.opponent)
	drawMonsterStats(0, int32(window.GameHeight)-50, b.playerMonster)

	rl.EndTextureMode()
}

func drawMonsterStats(x, y int32, mons *monster.Monster) {
	rl.DrawText(fmt.Sprintf("%s   (%d / %d)", mons.Name, mons.Health, mons.MaxHealth), x, y, 20, rl.Black)

	healthPercent := float64(mons.Health) / float64(mons.MaxHealth)
	remainingHealth := int(healthPercent * 50)
	damageTaken := 50 - remainingHealth
	healthStr := ""
	damageStr := ""
	for i := 0; i < remainingHealth; i++ {
		healthStr += "-"
	}

	for j := 0; j < damageTaken; j++ {
		damageStr += " "
	}
	rl.DrawText(fmt.Sprintf("<(%s%s)>", healthStr, damageStr), x, y+25, 20, rl.Black)
}
