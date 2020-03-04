package game

// getCurrentPlayerInput ...
func (g *Game) getCurrentPlayerInput() error {
	g.waveState.tick()

	if g.waveState.proceed() {
		g.screenState.Fire(TriggerWaveProcessing)
	}

	return nil
}
