package states

import (
	"fmt"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/ludum-dare-56/internal/systems"
	"github.com/timsims1717/pixel-go-utils/debug"
	"github.com/timsims1717/pixel-go-utils/img"
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
	systems.CreateCharacter()
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")

	systems.AnimationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *gameState) Draw(win *pixelgl.Window) {
	systems.DrawLayerSystem(win, 1)
	img.Clear()
	systems.TemporarySystem()

	//if options.Updated {
	//	systems.UpdateViews()
	//}
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
