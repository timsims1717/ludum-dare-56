package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/gopxl/pixel/text"
	"github.com/timsims1717/ludum-dare-56/embed"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/loading"
	"github.com/timsims1717/ludum-dare-56/internal/states"
	"github.com/timsims1717/ludum-dare-56/internal/systems"
	"github.com/timsims1717/pixel-go-utils/debug"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/options"
	"github.com/timsims1717/pixel-go-utils/state"
	"github.com/timsims1717/pixel-go-utils/timing"
	"github.com/timsims1717/pixel-go-utils/typeface"
	"github.com/timsims1717/pixel-go-utils/viewport"
	"golang.org/x/image/colornames"
)

func run() {
	winWidth := 1600.
	winHeight := 900.
	options.RegisterResolution(pixel.V(winWidth, winHeight))
	cfg := pixelgl.WindowConfig{
		Title:  "Ludum Dare 56",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
		VSync:  true,
	}
	options.BilinearFilter = false
	options.VSync = true
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//win.SetCursorVisible(false)
	viewport.ILockDefault = true
	viewport.MainCamera = viewport.New(win.Canvas())
	viewport.MainCamera.SetRect(pixel.R(0, 0, winWidth, winHeight))
	viewport.MainCamera.CamPos = pixel.V(winWidth*0.5, winHeight*0.5)
	object.ILock = false

	state.Register(data.GameStateKey, state.New(states.GameState))

	loading.LoadSprites()
	loading.LoadEnities()
	loading.LoadStatuses()

	mainFont, err := typeface.LoadBytes(embed.JiveTalking, 128.)
	typeface.Atlases["main"] = text.NewAtlas(mainFont, text.ASCII)

	debug.Initialize(&viewport.MainCamera.PostCamPos)
	debug.ShowText = false
	debug.ShowDebug = false

	systems.MainViewInit()

	win.Show()
	timing.Reset()
	for !win.Closed() {
		timing.Update()
		debug.Clear()
		options.WindowUpdate(win)
		if options.Updated {
			viewport.MainCamera.CamPos = pixel.V(viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.Rect.H()*0.5)
		}

		data.DebugInput.Update(win, viewport.MainCamera.Mat)
		if data.DebugInput.Get("debugPause").JustPressed() {
			state.ToggleDebugPause()
		}
		if data.DebugInput.Get("debugFrame").JustPressed() || data.DebugInput.Get("debugFrame").Repeated() {
			state.DebugFrameAdvance()
		}
		if data.DebugInput.Get("fullscreen").JustPressed() {
			options.FullScreen = !options.FullScreen
		}
		if data.DebugInput.Get("fuzzy").JustPressed() {
			options.BilinearFilter = !options.BilinearFilter
		}
		if data.DebugInput.Get("vsync").JustPressed() {
			options.VSync = !options.VSync
		}
		if data.DebugInput.Get("debugText").JustPressed() {
			debug.ShowText = !debug.ShowText
		}
		if data.DebugInput.Get("debug").JustPressed() {
			debug.ShowDebug = !debug.ShowDebug
		}
		if data.DebugInput.Get("layers").JustPressed() {
			data.Layers = !data.Layers
		}

		state.Update(win)
		win.Clear(colornames.Aliceblue)

		viewport.MainCamera.Update()
		state.Draw(win)

		win.SetSmooth(false)
		debug.Draw(win)
		win.SetSmooth(options.BilinearFilter)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
