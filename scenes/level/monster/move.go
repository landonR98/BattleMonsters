package monster

const (
	Effect_Sleep int = 1 + iota
	Effect_Burn
	Effect_Poison
	Effect_Coma
	Effect_Defense
	Effect_Health
	Effect_Wrap
)

const (
	Target_Self int = iota
	Target_Opponent
)

type Move struct {
	Id           int
	Name         string
	Target       int
	Damage       float32
	Effect       int
	EffectChance float32
	Cost         int
}

func NewMove(m moveModel) Move {

	var move_effect int
	switch m.Effect {
	case "sleep", "Sleep":
		move_effect = Effect_Sleep
	case "burn", "Burn":
		move_effect = Effect_Burn
	case "poison", "Poison":
		move_effect = Effect_Poison
	case "coma", "Coma":
		move_effect = Effect_Coma
	case "defense", "Defense":
		move_effect = Effect_Defense
	case "health", "Health":
		move_effect = Effect_Health
	case "wrap", "Wrap":
		move_effect = Effect_Wrap
	}

	var move_target int
	switch m.Target {
	case "Opponent", "opponent":
		move_target = Target_Opponent
	case "self", "Self":
		move_target = Target_Self
	}

	return Move{
		Id:           m.Id,
		Name:         m.Name,
		Target:       move_target,
		Damage:       m.Damage,
		Effect:       move_effect,
		EffectChance: m.EffectChance,
		Cost:         m.Cost,
	}
}
