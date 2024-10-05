package states

import (
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/pixel-go-utils/debug"
	"github.com/timsims1717/pixel-go-utils/state"
)

var (
	GameState = &gameState{}
)

type gameState struct {
	*state.AbstractState
}

func (s *gameState) Unload(win *pixelgl.Window) {

}

func (s *gameState) Load(win *pixelgl.Window) {

}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")

}

func (s *gameState) Draw(win *pixelgl.Window) {

}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
