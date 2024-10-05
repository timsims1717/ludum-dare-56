package states

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/ludum-dare-56/internal/systems"
	"github.com/timsims1717/pixel-go-utils/debug"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/options"
	"github.com/timsims1717/pixel-go-utils/state"
	"github.com/timsims1717/pixel-go-utils/viewport"
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
	systems.MainViewInit()
	systems.UpdateViews()
	systems.CreateCharacter()
	systems.CreateEntity()
	systems.CreateNPC()
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")
	data.PlayerInput.Update(win, viewport.MainCamera.Mat)

	systems.PlayerCharacterSystem()
	systems.NonPlayerCharacterSystem()
	systems.PickUpSystem()
	systems.RoomBorderSystem()

	systems.AnimationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.MainCanvas.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *gameState) Draw(win *pixelgl.Window) {
	data.MainCanvas.Canvas.Clear(pixel.RGBA{})
	systems.DrawLayerSystem(data.MainCanvas.Canvas, 1)
	img.Clear()
	data.MainCanvas.Draw(win)

	systems.TemporarySystem()

	if options.Updated {
		systems.UpdateViews()
	}
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
