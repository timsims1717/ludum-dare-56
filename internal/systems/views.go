package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/constants"
	"github.com/timsims1717/pixel-go-utils/viewport"
	"math"
)

func MainViewInit() {
	if constants.MainCanvas == nil {
		constants.MainCanvas = viewport.New(nil)
		constants.MainCanvas.SetRect(pixel.R(0, 0, constants.CanvasWidth, constants.CanvasHeight))
		constants.MainCanvas.CamPos = pixel.ZV
		constants.MainCanvas.PortPos = viewport.MainCamera.CamPos
	}
}

func UpdateViews() {
	constants.MainCanvas.PortPos = viewport.MainCamera.PostCamPos
	ratioY := viewport.MainCamera.Rect.H() / constants.MainCanvas.Rect.H()
	ratioX := viewport.MainCamera.Rect.W() / constants.MainCanvas.Rect.W()
	ratio := math.Min(ratioX, ratioY)
	constants.MainCanvas.PortSize = pixel.V(ratio, ratio)
}
