package systems

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/ludum-dare-56/internal/ui"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/util"
	"github.com/timsims1717/pixel-go-utils/viewport"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

func DialogSystem(win *pixelgl.Window) {
	var updated []string
	layer := 100
	for _, dialog := range ui.DialogsOpen {
		dialog.Active = !ui.DialogStackOpen
		updated = append(updated, dialog.Key)
		layer = UpdateDialog(dialog, layer)
	}
	closeKey := ""
	for i, dialog := range ui.DialogStack {
		dialog.Active = i == len(ui.DialogStack)-1
		if dialog.Active {
			closeDlg := false
			// update this dialog box with input
			switch dialog.Key {
			default:
				if data.PlayerInput.Get("escape").JustPressed() {
					data.PlayerInput.Get("escape").Consume()
					closeDlg = true
				}
			}
			if closeDlg {
				closeKey = dialog.Key
			}
		}
		updated = append(updated, dialog.Key)
		layer = UpdateDialog(dialog, layer)
	}
	if closeKey != "" {
		ui.CloseDialog(closeKey)
	}
	layer += 100
	for key, dialog := range ui.Dialogs {
		if !util.ContainsStr(key, updated) {
			//UpdateDialogLayer99(dialog)
			layer = UpdateDialogLayers(dialog, layer)
		}
	}
}

func UpdateDialog(dialog *ui.Dialog, layer int) int {
	dialog.Loaded = true
	dialog.Layer = layer
	if !dialog.NoBorder {
		dialog.BorderVP.Update()
		dialog.BorderObject.Layer = layer
	}
	dialog.ViewPort.Update()
	nextLayer := UpdateSubElements(dialog.Elements, dialog.ViewPort, layer)
	return nextLayer
}

func UpdateSubElements(elements []*ui.Element, vp *viewport.ViewPort, layer int) int {
	nextLayer := layer + 1
	for _, e := range elements {
		e.Object.Unloaded = !vp.RectInside(e.Object.Rect.Moved(e.Object.Pos))
		switch e.ElementType {
		case ui.SpriteElement, ui.ButtonElement, ui.CheckboxElement:
			e.Object.Layer = layer
		case ui.TextElement:
			e.Text.Obj.Layer = layer
		case ui.InputElement:
			e.Layer = nextLayer
			e.BorderObject.Layer = e.Layer
			e.Text.Obj.Layer = e.Layer
			e.CaretObj.Layer = e.Layer
			if !e.Object.Hidden && !e.Object.Unloaded {
				e.BorderVP.Update()
				e.ViewPort.Update()
			}
			nextLayer++
		case ui.ScrollElement, ui.ContainerElement:
			e.BorderObject.Layer = nextLayer
			e.Object.Layer = nextLayer
			e.Layer = nextLayer
			if !e.Object.Hidden && !e.Object.Unloaded {
				e.BorderVP.Update()
				e.ViewPort.Update()
				nextLayer = UpdateSubElements(e.Elements, e.ViewPort, e.Layer)
			} else {
				nextLayer = UpdateSubElementLayers(e.Elements, e.Layer)
			}
		}
	}
	return nextLayer
}

func UpdateDialogLayers(dialog *ui.Dialog, layer int) int {
	dialog.Layer = layer
	if !dialog.NoBorder {
		//dialog.BorderVP.Update()
		dialog.BorderObject.Layer = layer
	}
	//dialog.ViewPort.Update()
	layer = UpdateSubElementLayers(dialog.Elements, layer)
	layer++
	return layer
}

func UpdateSubElementLayers(elements []*ui.Element, layer int) int {
	nextLayer := layer + 1
	for _, e := range elements {
		switch e.ElementType {
		case ui.SpriteElement, ui.ButtonElement, ui.CheckboxElement:
			e.Object.Layer = layer
		case ui.TextElement:
			e.Text.Obj.Layer = layer
		case ui.InputElement:
			e.Layer = nextLayer
			//e.BorderVP.Update()
			e.BorderObject.Layer = e.Layer
			//e.ViewPort.Update()
			e.Text.Obj.Layer = e.Layer
			e.CaretObj.Layer = e.Layer
			nextLayer++
		case ui.ScrollElement, ui.ContainerElement:
			//e.BorderVP.Update()
			e.BorderObject.Layer = nextLayer
			e.Object.Layer = nextLayer
			//e.ViewPort.Update()
			e.Layer = nextLayer
			nextLayer = UpdateSubElementLayers(e.Elements, e.Layer)
		}
	}
	return nextLayer
}

func UpdateDialogLayer99(dialog *ui.Dialog) {
	dialog.Layer = 99
	if !dialog.NoBorder {
		//dialog.BorderVP.Update()
		dialog.BorderObject.Layer = 99
	}
	//dialog.ViewPort.Update()
	UpdateSubElementLayer99(dialog.Elements)
}

func UpdateSubElementLayer99(elements []*ui.Element) {
	for _, e := range elements {
		switch e.ElementType {
		case ui.SpriteElement, ui.ButtonElement, ui.CheckboxElement:
			e.Object.Layer = 99
		case ui.TextElement:
			e.Text.Obj.Layer = 99
		case ui.InputElement:
			e.Layer = 99
			//e.BorderVP.Update()
			e.BorderObject.Layer = e.Layer
			//e.ViewPort.Update()
			e.Text.Obj.Layer = e.Layer
			e.CaretObj.Layer = e.Layer
		case ui.ScrollElement, ui.ContainerElement:
			//e.BorderVP.Update()
			e.BorderObject.Layer = 99
			e.Object.Layer = 99
			//e.ViewPort.Update()
			e.Layer = 99
			UpdateSubElementLayer99(e.Elements)
		}
	}
}

func DialogDrawSystem(win *pixelgl.Window) {
	for _, dialog := range ui.DialogsOpen {
		DrawDialog(dialog, win)
	}
	for _, dialog := range ui.DialogStack {
		DrawDialog(dialog, win)
	}
}

func DrawDialog(dialog *ui.Dialog, win *pixelgl.Window) {
	if !dialog.NoBorder {
		dialog.BorderVP.Canvas.Clear(color.RGBA{})
		BorderSystem(dialog.Layer)
		img.Batchers[data.BatchKeyTest].Draw(dialog.BorderVP.Canvas)
		dialog.BorderVP.Draw(win)
		img.Clear()
	}
	// draw elements w/no sub elements
	dialog.ViewPort.Canvas.Clear(colornames.Darkgray)
	DrawLayerSystem(dialog.ViewPort.Canvas, dialog.Layer, false)
	img.Clear()
	// draw elements w/sub elements
	for _, e := range dialog.Elements {
		DrawSubElements(e, dialog.ViewPort)
	}
	dialog.ViewPort.Draw(win)
	img.Clear()
}

// DrawSubElements draws the border and sub elements of ui elements
// with sub elements.
func DrawSubElements(element *ui.Element, vp *viewport.ViewPort) {
	if element == nil {
		return
	}
	if element.Object.Hidden {
		return
	}
	if element.Object.Unloaded {
		return
	}
	switch element.ElementType {
	case ui.InputElement:
		// draw border
		element.BorderVP.Canvas.Clear(color.RGBA{})
		BorderSystem(element.Layer)
		img.Batchers[data.BatchKeyTest].Draw(element.BorderVP.Canvas)
		element.BorderVP.Draw(vp.Canvas)
		img.Clear()
		// draw input
		element.ViewPort.Canvas.Clear(pixel.RGBA{})
		DrawLayerSystem(element.ViewPort.Canvas, element.Layer, false)
		element.ViewPort.Draw(vp.Canvas)
		img.Clear()
	case ui.ScrollElement, ui.ContainerElement:
		// draw border
		element.BorderVP.Canvas.Clear(color.RGBA{})
		BorderSystem(element.Layer)
		img.Batchers[data.BatchKeyTest].Draw(element.BorderVP.Canvas)
		element.BorderVP.Draw(vp.Canvas)
		img.Clear()
		// draw container elements
		element.ViewPort.Canvas.Clear(pixel.RGBA{})
		DrawLayerSystem(element.ViewPort.Canvas, element.Layer, false)
		img.Clear()
		for _, e := range element.Elements {
			DrawSubElements(e, element.ViewPort)
		}
		element.ViewPort.Draw(vp.Canvas)
		img.Clear()
	}
}

func BorderSystem(layer int) {
	for _, result := range myecs.Manager.Query(myecs.HasBorder) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		bord, okB := result.Components[myecs.Border].(*ui.Border)
		if okO && okB && obj.Layer == layer {
			switch bord.Style {
			case ui.FancyBorder:
				DrawFancyBorder(bord, obj)
			case ui.ThinBorder:
				DrawThinBorder(bord, obj)
			}
		}
	}
}

func DrawFancyBorder(bord *ui.Border, obj *object.Object) {
	for y := 0; y < bord.Height+1; y++ {
		if y == 0 || y == bord.Height {
			for x := 0; x < bord.Width+1; x++ {
				DrawFancyBorderSection(x, y, bord, obj)
				if !bord.Empty && y == bord.Height && x != 0 {
					DrawBlackSquare(x, y, bord, obj)
				}
			}
		} else {
			for x := 0; x < bord.Width+1; x++ {
				if x == 0 || x == bord.Width {
					DrawFancyBorderSection(x, y, bord, obj)
					if !bord.Empty && x == bord.Width {
						DrawBlackSquare(x, y, bord, obj)
					}
				} else if !bord.Empty {
					DrawBlackSquare(x, y, bord, obj)
				}
			}
		}
	}
}

func DrawFancyBorderSection(x, y int, bord *ui.Border, obj *object.Object) {
	mat := pixel.IM
	offset := pixel.V(16.*(float64(x)-float64(bord.Width)*0.5), 16.*(float64(y)-float64(bord.Height)*0.5))
	sKey := data.SpriteChalkboardSide
	if (x == 0 || x == bord.Width) && (y == 0 || y == bord.Height) {
		sKey = data.SpriteChalkboardCorner
	}
	if y == 0 {
		if x > 0 && x < bord.Width {
			mat = mat.Rotated(pixel.ZV, 0.5*math.Pi)
		} else if x == bord.Width {
			mat = mat.ScaledXY(pixel.ZV, pixel.V(-1., 1.))
		}
	} else if y == bord.Height {
		if x > 0 && x < bord.Width {
			mat = mat.Rotated(pixel.ZV, -0.5*math.Pi)
		} else if x == 0 {
			mat = mat.ScaledXY(pixel.ZV, pixel.V(1., -1.))
		} else if x == bord.Width {
			mat = mat.ScaledXY(pixel.ZV, pixel.V(-1., -1.))
		}
	} else if x == bord.Width {
		mat = mat.ScaledXY(pixel.ZV, pixel.V(-1., 1.))
	}
	img.Batchers[data.BatchKeyTest].DrawSpriteColor(sKey, mat.Moved(obj.PostPos).Moved(offset), colornames.White)
}

func DrawThinBorder(bord *ui.Border, obj *object.Object) {
	matTB := pixel.IM.ScaledXY(pixel.ZV, pixel.V(bord.Rect.W()+2, 1.))
	matLR := pixel.IM.ScaledXY(pixel.ZV, pixel.V(1., bord.Rect.H()+2))
	// top
	img.Batchers[data.BatchKeyTest].DrawSprite(data.SpriteChalkboardWhite, matTB.Moved(obj.PostPos).Moved(pixel.V(0, bord.Rect.H()*0.5+0.5)))
	// right
	img.Batchers[data.BatchKeyTest].DrawSprite(data.SpriteChalkboardWhite, matLR.Moved(obj.PostPos).Moved(pixel.V(bord.Rect.W()*0.5+0.5, 0)))
	// bottom
	img.Batchers[data.BatchKeyTest].DrawSprite(data.SpriteChalkboardWhite, matTB.Moved(obj.PostPos).Moved(pixel.V(0, bord.Rect.H()*-0.5-0.5)))
	// left
	img.Batchers[data.BatchKeyTest].DrawSprite(data.SpriteChalkboardWhite, matLR.Moved(obj.PostPos).Moved(pixel.V(bord.Rect.W()*-0.5-0.5, 0)))
}

func DrawBlackSquare(x, y int, bord *ui.Border, obj *object.Object) {
	mat := pixel.IM
	offset := pixel.V(16.*(float64(x)-float64(bord.Width+1)*0.5), 16.*(float64(y)-float64(bord.Height+1)*0.5))
	sKey := data.SpriteChalkboardBlackSquare
	img.Batchers[data.BatchKeyTest].DrawSpriteColor(sKey, mat.Moved(obj.PostPos).Moved(offset), colornames.White)
}
