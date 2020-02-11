package game

var statToMod = map[int]int{
	0:  -3,
	1:  -2,
	2:  -1,
	3:  0,
	4:  1,
	5:  2,
	6:  3,
	7:  4,
	8:  5,
	9:  6,
	10: 7,
}

// vitalityForStat also produces the focus for the stat
func vitalityForStat(stat int) int {
	return 20 + (stat * 5)
}

// evasionForStat also produces the willpower for the stat
func evasionForStat(stat int) int {
	return (stat * 2) + 5
}

// xpForLevel returns how much XP is required to get to the given level
func xpForLevel(level int) int {
	if level == 1 {
		return 0
	}

	if level == 2 {
		return 6
	}

	return level + xpForNextLevel(level-1)
}

func budgetForWave(num int) int {
	if num == 1 {
		return 2
	}

	return num + budgetForWave(num-1)
}
