package data

import (
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/ludum-dare-56/internal/constants"
	pxginput "github.com/timsims1717/pixel-go-input"
)

var (
	PlayerInput = &pxginput.Input{
		Buttons: map[string]*pxginput.ButtonSet{
			constants.InputLeft:  pxginput.NewJoyless(pixelgl.KeyLeft),
			constants.InputRight: pxginput.NewJoyless(pixelgl.KeyRight),
			constants.InputUp:    pxginput.NewJoyless(pixelgl.KeyUp),
			constants.InputDown:  pxginput.NewJoyless(pixelgl.KeyDown),
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
			"debugMenu":    pxginput.NewJoyless(pixelgl.KeyF7),
			"debugTest":    pxginput.NewJoyless(pixelgl.KeyF8),
			"debugPause":   pxginput.NewJoyless(pixelgl.KeyF9),
			"debugFrame":   pxginput.NewJoyless(pixelgl.KeyF10),
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
