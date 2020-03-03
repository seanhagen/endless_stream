package game

import (
	"math"

	"github.com/seanhagen/endless_stream/backend/endless"
)

var _ actor = &monster{}

type monster struct {
	*creature
	isBoss bool
	cost   int32
}

type monsterBase struct {
	Name         string
	Description  string
	Type         string
	IsFlying     bool
	IsBoss       bool
	Mod          int32
	CostMod      int32
	GoldMod      float64
	XPMod        float64
	Strength     int32
	Intelligence int32
	Agility      int32
	Script       string
}

type monsterMap map[string]monster

const eqThres = 1e-9

func equalEnough(a, b float64) bool {
	return math.Abs(a-b) <= eqThres
}

// createMonster reads in the configuration and builds the creature stats
func createMonster(id string, in monsterBase, script string) monster {
	if equalEnough(in.XPMod, 0.0) {
		in.XPMod = 1
	}

	if equalEnough(in.GoldMod, 0.0) {
		in.GoldMod = 2.0
	}

	var t endless.Type
	if x, ok := endless.Type_value[in.Type]; ok {
		t = endless.Type(x)
	}
	p := endless.Position_Right
	cr := &creature{
		Id:           id,
		Name:         in.Name,
		Description:  in.Description,
		Strength:     in.Strength,
		Intelligence: in.Intelligence,
		Agility:      in.Agility,
		Script:       script,
		Position:     &p,
		mType:        t,
		isFlying:     in.IsFlying,
	}
	cr.init()

	mod := in.Mod + (cr.Strength - 3)
	// fmt.Printf("mod = m_mod ( str - 3 ) = %v + ( %v - 3 ) = %v\n", in.Mod, cr.Strength, mod)
	cost := calcMonster(cr, in.CostMod, mod, in.GoldMod, in.XPMod, in.IsBoss)
	return monster{
		creature: cr,
		isBoss:   in.IsBoss,
		cost:     cost,
	}
}

func calcMonster(cr *creature, costMod, mod int32, goldMod, xpMod float64, isBoss bool) int32 {
	var tmp int32 = cr.MaxVitality +
		mod +
		cr.MaxFocus +
		cr.Willpower +
		cr.Evasion +
		cr.Accuracy +
		(costMod * 12)

	// fmt.Printf("roundup((%v + %v + %v + %v + %v + %v + (%v * 12))/12) = %v\n",
	// 	cr.MaxVitality,
	// 	mod,
	// 	cr.MaxFocus,
	// 	cr.Willpower,
	// 	cr.Evasion,
	// 	cr.Accuracy,
	// 	costMod,
	// 	tmp)

	// =ROUNDUP(
	//    (
	//      Vitality +
	//      MOD +
	//      Focus +
	//      Willpower +
	//      Evasion +
	//      Accuracy +
	//      (Special * 12)
	//    ) / 12)

	costTmp := float64(tmp) / 12.0
	// fmt.Printf("cost tmp: %v\n", costTmp)
	cost := int32(math.Ceil(costTmp))

	// fmt.Printf("cost: %v\n", cost)

	var goldTmp float64
	if isBoss {
		goldTmp = float64(cost) * goldMod
		// fmt.Printf("gold: %v * %v\n", cost, goldMod)
	} else {
		goldTmp = float64(cost) / goldMod
		// fmt.Printf("gold: %v / %v\n", cost, goldMod)
	}

	// fmt.Printf("gold tmp: %v\n", goldTmp)

	xpTmp := (goldTmp / goldMod) * xpMod

	xp := int32(math.Ceil(xpTmp))
	gold := int32(math.Ceil(goldTmp))

	cr.XP = xp
	cr.Gold = gold

	return cost
}

// spawn takes the base monster and creates a copy with it's 'brain' all ready to go
func (m monster) spawn() (*monster, error) {
	cc, err := m.creature.spawn()
	if err != nil {
		return nil, err
	}
	return &monster{creature: cc, isBoss: m.isBoss, cost: m.cost}, nil
}
