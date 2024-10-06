package systems

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func BuildRoom() {
	// wood flooring
	for x := data.WoodLeftIndex; x < data.WoodRightIndex; x++ {
		xPos := data.WoodWidth * float64(x)
		h1 := 0
		h2 := 1
		for h1 < data.FloorHeight {
			if data.FloorHeight-h1 < 5 {
				h2 = data.FloorHeight - 1
			} else if data.FloorHeight-h1 < 9 {
				h2 = h1 + data.GlobalSeededRandom.Intn(data.FloorHeight-h1-4)
			} else {
				h2 = h1 + data.GlobalSeededRandom.Intn(6) + 4
			}
			halfLen := float64(h2-h1) * 0.5
			yPos := data.RoomBottom + float64(h1)*data.WoodWidth + data.WoodWidth*halfLen
			woodObj := object.New()
			woodObj.Layer = 0
			woodObj.Pos = pixel.V(xPos, yPos)

			r := data.GlobalSeededRandom.Intn(8)
			var sprs []*img.Sprite
			for h := h1; h <= h2; h++ {
				spr := img.NewSprite(fmt.Sprintf("wood_floor_%d", r), data.BatchKeyTest).
					WithOffset(pixel.V(0., (float64(h-h1)-halfLen)*data.WoodWidth))
				sprs = append(sprs, spr)
				r++
				r %= 8
			}

			myecs.Manager.NewEntity().
				AddComponent(myecs.Object, woodObj).
				AddComponent(myecs.Drawable, sprs)
			h1 = h2 + 1
		}
	}
	// rug
	//for x := data.RugLeftIndex; x < data.RugRightIndex; x++ {
	//	xPos := data.WoodWidth * float64(x)
	//	h1 := 0
	//	h2 := 1
	//	for h1 < data.FloorHeight {
	//		if data.FloorHeight-h1 < 5 {
	//			h2 = data.FloorHeight - 1
	//		} else if data.FloorHeight-h1 < 9 {
	//			h2 = h1 + data.GlobalSeededRandom.Intn(data.FloorHeight-h1-4)
	//		} else {
	//			h2 = h1 + data.GlobalSeededRandom.Intn(6) + 4
	//		}
	//		halfLen := float64(h2-h1) * 0.5
	//		yPos := data.RoomBottom + float64(h1)*data.WoodWidth + data.WoodWidth*halfLen
	//		woodObj := object.New()
	//		woodObj.Layer = 0
	//		woodObj.Pos = pixel.V(xPos, yPos)
	//
	//		r := data.GlobalSeededRandom.Intn(8)
	//		var sprs []*img.Sprite
	//		for h := h1; h <= h2; h++ {
	//			spr := img.NewSprite(fmt.Sprintf("wood_floor_%d", r), data.BatchKeyTest).
	//				WithOffset(pixel.V(0., (float64(h-h1)-halfLen)*data.WoodWidth))
	//			sprs = append(sprs, spr)
	//			r++
	//			r %= 8
	//		}
	//
	//		myecs.Manager.NewEntity().
	//			AddComponent(myecs.Object, woodObj).
	//			AddComponent(myecs.Drawable, sprs)
	//		h1 = h2 + 1
	//	}
	//}
	// welcome mat
	matObj := object.New()
	matObj.Layer = 0
	matObj.Pos = data.MatPos
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, matObj).
		AddComponent(myecs.Drawable, img.NewSprite(data.SpriteKeyMat, data.BatchKeyTest))
}
