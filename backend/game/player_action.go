package game

import (
	"log"

	"github.com/seanhagen/endless_stream/backend/endless"
)

// setActionSkill ...
func (p *player) setActionSkill(sk *endless.Input_Skill) actionMessage {
	log.Printf("player wants to use a skill: %#v", sk)
	return skipMsg{}
}

// setActionItem ...
func (p *player) setActionItem(it *endless.Input_Item) actionMessage {
	log.Printf("player wants to use an item: %#v", it)
	return skipMsg{}
}

// setActionMove ...
func (p *player) setActionMove(m *endless.Input_Move) actionMessage {
	log.Printf("player wants to move: %#v", m)
	return skipMsg{}
}
