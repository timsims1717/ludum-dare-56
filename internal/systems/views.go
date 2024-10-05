package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/pixel-go-utils/viewport"
	"math"
)

func MainViewInit() {
	if data.MainCanvas == nil {
		data.MainCanvas = viewport.New(nil)
		data.MainCanvas.SetRect(pixel.R(0, 0, data.CanvasWidth, data.CanvasHeight))
		data.MainCanvas.CamPos = pixel.ZV
		data.MainCanvas.PortPos = viewport.MainCamera.CamPos
	}
}

func UpdateViews() {
	data.MainCanvas.PortPos = viewport.MainCamera.PostCamPos
	ratioY := viewport.MainCamera.Rect.H() / data.MainCanvas.Rect.H()
	ratioX := viewport.MainCamera.Rect.W() / data.MainCanvas.Rect.W()
	ratio := math.Min(ratioX, ratioY)
	data.MainCanvas.PortSize = pixel.V(ratio, ratio)
}
