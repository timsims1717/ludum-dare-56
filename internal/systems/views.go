package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/ui"
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
	data.Ratio = math.Min(ratioX, ratioY)
	data.MainCanvas.PortSize = pixel.V(data.Ratio, data.Ratio)
	for _, dialog := range ui.Dialogs {
		UpdateDialogViews(dialog)
	}
}

func UpdateDialogViews(dialog *ui.Dialog) {
	posRatX := viewport.MainCamera.Rect.W() / 1600.
	posRatY := viewport.MainCamera.Rect.H() / 900.
	nPos := pixel.V(dialog.Pos.X*posRatX, dialog.Pos.Y*posRatY)
	if !dialog.NoBorder {
		dialog.BorderVP.PortPos = viewport.MainCamera.PostCamPos.Add(nPos)
		dialog.BorderVP.PortSize = pixel.V(data.Ratio, data.Ratio)
	}
	dialog.ViewPort.PortPos = viewport.MainCamera.PostCamPos.Add(nPos)
	dialog.ViewPort.PortSize = pixel.V(data.Ratio, data.Ratio)
}
