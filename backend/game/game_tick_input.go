package game

// getCurrentPlayerInput ...
func (g *Game) getCurrentPlayerInput() error {
	// get current actor
	a := g.waveState.current()
	//   get input
	act := a.act()
	if act != nil {
		//   if valid, store and continue
		g.currentAction = act
		g.screenState.Fire(TriggerWaveProcessing)
	}

	return nil
}
