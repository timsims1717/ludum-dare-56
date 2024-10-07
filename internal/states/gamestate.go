package states

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/ludum-dare-56/internal/systems"
	"github.com/timsims1717/ludum-dare-56/internal/ui"
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
	systems.LoadDialogs()
	systems.MainViewInit()
	systems.BuildRoom()
	systems.UpdateViews()
	systems.CreateCharacter()
	systems.InitGameplay()
	systems.PopulateLandscape()
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")
	data.PlayerInput.Update(win, viewport.MainCamera.Mat)

	ui.DialogStackOpen = len(ui.DialogStack) > 0
	if !ui.DialogStackOpen && data.PlayerInput.Get("escape").JustPressed() {
		data.PlayerInput.Get("escape").Consume()
		ui.OpenDialogInStack(data.DialogPause)
	}
	systems.DialogSystem(win)

	if !ui.DialogStackOpen {
		// game control systems
		switch data.TheGamePhase {
		case data.ParentDropOff:
			systems.DropOffSystem()
		case data.Gameplay:
			systems.GameplaySystem()
		case data.ParentPickUp:
			systems.ParentPickUpSystem()
		}

		// entity control systems
		systems.KidBehaviorSystem()
		systems.KidParentBehaviorSystem()
		systems.PlayerMoveSystem()
		systems.NonPlayerMoveSystem()
		systems.DoorSystem()
		systems.PickUpSystem()
		systems.NPCCollisions()
		systems.RoomBorderSystem()
		systems.EntityInteractions()
		systems.UpdateInvincibility()
		systems.UpdateStatusEffects()
	}
	systems.HealthBarSystem()

	systems.AnimationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.MainCanvas.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *gameState) Draw(win *pixelgl.Window) {
	data.MainCanvas.Canvas.Clear(pixel.RGBA{})
	systems.DrawLayerSystem(data.MainCanvas.Canvas, -2, false)
	systems.DrawLayerSystem(data.MainCanvas.Canvas, -1, data.Layers)
	systems.DrawLayerSystem(data.MainCanvas.Canvas, 0, false)
	systems.DrawLayerSystem(data.MainCanvas.Canvas, 1, data.Layers)
	systems.DrawLayerSystem(data.MainCanvas.Canvas, 3, data.Layers)
	systems.DrawLayerSystem(data.MainCanvas.Canvas, 4, false)
	systems.DrawLayerSystem(data.MainCanvas.Canvas, 5, false)
	img.Clear()
	data.MainCanvas.Draw(win)

	systems.DialogDrawSystem(win)
	systems.DrawLayerSystem(win, -10, false)
	img.Clear()

	systems.TemporarySystem()

	if options.Updated {
		systems.UpdateViews()
	}
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
