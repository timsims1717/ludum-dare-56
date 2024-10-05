package data

import (
	"github.com/gopxl/pixel/pixelgl"
	pxginput "github.com/timsims1717/pixel-go-input"
)

var (
	PlayerInput = &pxginput.Input{
		Buttons: map[string]*pxginput.ButtonSet{
			InputLeft:   pxginput.NewJoyless(pixelgl.KeyLeft),
			InputRight:  pxginput.NewJoyless(pixelgl.KeyRight),
			InputUp:     pxginput.NewJoyless(pixelgl.KeyUp),
			InputDown:   pxginput.NewJoyless(pixelgl.KeyDown),
			InputAction: pxginput.NewJoyless(pixelgl.KeySpace),
		},
		Mode: pxginput.KeyboardMouse,
	}
	DebugInput = &pxginput.Input{
		Buttons: map[string]*pxginput.ButtonSet{
			"debugConsole": pxginput.NewJoyless(pixelgl.KeyGraveAccent),
			"debug":        pxginput.NewJoyless(pixelgl.KeyF3),
			"debugText":    pxginput.NewJoyless(pixelgl.KeyF4),
			"fullscreen":   pxginput.NewJoyless(pixelgl.KeyF5),
			"fuzzy":        pxginput.NewJoyless(pixelgl.KeyF6),
			"vsync":        pxginput.NewJoyless(pixelgl.KeyF7),
			"layers":       pxginput.NewJoyless(pixelgl.KeyF8),
			//"debugSP":      pxginput.NewJoyless(pixelgl.KeyEqual),
			//"debugSM":      pxginput.NewJoyless(pixelgl.KeyMinus),
			//"camUp":    pxginput.NewJoyless(pixelgl.KeyKP8),
			//"camRight": pxginput.NewJoyless(pixelgl.KeyKP6),
			//"camDown":  pxginput.NewJoyless(pixelgl.KeyKP5),
			//"camLeft":  pxginput.NewJoyless(pixelgl.KeyKP4),
		},
		Mode: pxginput.KeyboardMouse,
	}
)
