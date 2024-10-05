package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/gopxl/pixel/text"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/states"
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
		Title:  "Typeface Test",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
		VSync:  true,
	}
	options.BilinearFilter = false
	options.VSync = true
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetCursorVisible(false)
	viewport.ILockDefault = true
	viewport.MainCamera = viewport.New(win.Canvas())
	viewport.MainCamera.SetRect(pixel.R(0, 0, winWidth, winHeight))
	viewport.MainCamera.CamPos = pixel.V(winWidth*0.5, winHeight*0.5)
	object.ILock = true

	state.Register(data.GameStateKey, state.New(states.GameState))

	mainFont, err := typeface.LoadTTF("Jive_Talking.ttf", 128.)
	typeface.Atlases["main"] = text.NewAtlas(mainFont, text.ASCII)

	debug.Initialize(&viewport.MainCamera.PostCamPos)
	debug.ShowText = false
	debug.ShowDebug = false

	win.Show()
	timing.Reset()
	for !win.Closed() {
		timing.Update()
		debug.Clear()

		state.Update(win)
		win.Clear(colornames.Black)

		viewport.MainCamera.Update()
		state.Draw(win)

		win.SetSmooth(false)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
