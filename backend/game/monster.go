package game

import (
	"math"

	"github.com/seanhagen/endless_stream/backend/endless"
)

type monster struct {
	creature
	mType    endless.Type
	isFlying bool
	isBoss   bool
	cost     int32
}

type monsterBase struct {
	Name         string
	Description  string
	Type         string
	IsFlying     bool
	IsBoss       bool
	CostMod      int32
	MMod         int32
	GMod         int32
	Strength     int32
	Intelligence int32
	Agility      int32
	Script       string
}

type monsterList map[string]monster

// createMonster reads in the configuration and builds the creature stats
func createMonster(in monsterBase, script string) monster {
	cr := creature{
		Strength:     in.Strength,
		Intelligence: in.Intelligence,
		Agility:      in.Agility,
		Script:       script,
	}
	cr.init()

	var t endless.Type
	if x, ok := endless.Type_value[in.Type]; ok {
		t = endless.Type(x)
	}

	xp, gold, cost := calcMonster(cr, in.CostMod, in.MMod, in.GMod, in.IsBoss)

	cr.XP = xp
	cr.Gold = gold

	return monster{
		creature: cr,
		mType:    t,
		isFlying: in.IsFlying,
		isBoss:   in.IsBoss,
		cost:     cost,
	}
}

func calcMonster(cr creature, costMod, mMod, gMod int32, boss bool) (int32, int32, int32) {
	var mod int32 = mMod + gMod
	var cMod int32 = costMod * 12
	var tmp int32 = cr.MaxVitality +
		mod +
		cr.MaxFocus +
		cr.Willpower +
		cr.Evasion +
		cr.Accuracy +
		cMod

	costTmp := float64(tmp) / 12.0
	cost := int32(math.Ceil(costTmp))

	var goldTmp float64
	if boss {
		goldTmp = float64(cost) * 2.0
	} else {
		goldTmp = float64(cost) / 2.0
	}

	xpTmp := goldTmp / 2.0

	xp := int32(math.Ceil(xpTmp))
	gold := int32(math.Ceil(goldTmp))

	return xp, gold, cost
}

// spawn takes the base monster and creates a copy with it's 'brain' all ready to go
func (m monster) spawn() monster {

	return monster{}
}
