package game

import (
	"fmt"
	"runtime"
)

// Lock ...
func (g *Game) Lock() {
	fr := getFrame(1)
	fmt.Printf("Game lock called by %v\n", fr.Function)
	g.lock.Lock()
	fmt.Printf("Game lock granted to %v\n", fr.Function)
}

// Unlock ...
func (g *Game) Unlock() {
	fr := getFrame(1)
	fmt.Printf("Game unlock called by %v\n", fr.Function)
	g.lock.Unlock()
	fmt.Printf("Game unlocked by %v\n", fr.Function)
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
