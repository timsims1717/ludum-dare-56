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
		for h1 < data.WoodHeight {
			if data.WoodHeight-h1 < 5 {
				h2 = data.WoodHeight - 1
			} else if data.WoodHeight-h1 < 9 {
				h2 = h1 + data.GlobalSeededRandom.Intn(data.WoodHeight-h1-4)
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
	for x := data.RugLeftIndex; x < data.RugRightIndex; x++ {
		xPos := data.RugWidth * float64(x)
		for y := data.RugBottomIndex; y < data.RugTopIndex; y++ {
			yPos := data.RugWidth * float64(y)
			rugObj := object.New()
			rugObj.Layer = 0
			rugObj.Pos = pixel.V(xPos, yPos)

			sprV := "m"
			sprH := "m"
			if x == data.RugLeftIndex {
				sprH = "l"
			} else if x == data.RugRightIndex-1 {
				sprH = "r"
			}
			if y == data.RugBottomIndex {
				sprV = "b"
			} else if y == data.RugTopIndex-1 {
				sprV = "t"
			}

			myecs.Manager.NewEntity().
				AddComponent(myecs.Object, rugObj).
				AddComponent(myecs.Drawable, img.NewSprite(fmt.Sprintf("rug_%s_%s", sprV, sprH), data.BatchKeyTest))
		}
	}
	// side walls
	for i := 0; i < 32; i++ {
		xLPos := (data.WoodLeftIndex - 1) * data.WoodWidth
		yPos := data.RoomBottom + float64(i)*data.WoodWidth
		wallLObj := object.New()
		wallLObj.Layer = 0
		wallLObj.Flip = true
		wallLObj.Pos = pixel.V(xLPos, yPos)
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, wallLObj).
			AddComponent(myecs.Drawable, img.NewSprite("side_wall", data.BatchKeyTest))
		xRPos := (data.WoodRightIndex) * data.WoodWidth
		wallRObj := object.New()
		wallRObj.Layer = 0
		wallRObj.Pos = pixel.V(xRPos, yPos)
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, wallRObj).
			AddComponent(myecs.Drawable, img.NewSprite("side_wall", data.BatchKeyTest))
	}
	// front walls
	for x := data.WoodLeftIndex; x < data.WoodRightIndex; x++ {
		xPos := data.WoodWidth * float64(x)
		for y := 0; y < 6; y++ {
			yPos := data.RoomBottom + float64(data.WoodHeight+y)*data.WoodWidth
			sprKey := ""
			wallObj := object.New()
			wallObj.Layer = 2
			if x < -2 || x > 2 {
				switch y {
				case 0:
					sprKey = "baseboard"
				case 1, 2, 3, 4, 5:
					sprKey = "back_wall"
				}
			} else if x < -1 || x > 1 {
				switch y {
				case 0:
					sprKey = "baseboard_corner"
				case 1, 2, 3:
					sprKey = "back_wall_door"
				case 4:
					sprKey = "door_upper_corner"
				case 5:
					sprKey = "back_wall"
				}
			} else if y == 4 {
				sprKey = "door_above"
			} else if y == 5 {
				sprKey = "back_wall"
			} else {
				sprKey = "sidewalk"
				wallObj.Layer = 0
			}
			if x < -1 {
				wallObj.Flip = true
			}
			wallObj.Pos = pixel.V(xPos, yPos)
			myecs.Manager.NewEntity().
				AddComponent(myecs.Object, wallObj).
				AddComponent(myecs.Drawable, img.NewSprite(sprKey, data.BatchKeyTest))
		}
	}
	// welcome mat
	matObj := object.New()
	matObj.Layer = 0
	matObj.Pos = data.MatPos
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, matObj).
		AddComponent(myecs.Drawable, img.NewSprite(data.SpriteKeyMat, data.BatchKeyTest))
	// doorway
	doorObj := object.New()
	doorObj.Layer = 1
	doorObj.Pos = data.DoorPos
	data.DoorEntity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, doorObj).
		AddComponent(myecs.Drawable, img.NewSprite(data.SpriteKeyDoorClosed, data.BatchKeyTest))
}

func DoorSystem() {
	doorOpen := false
	for _, result := range myecs.Manager.Query(myecs.HasMoveTarget) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		_, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC {
			if obj.Rect.Moved(obj.Pos).Intersects(data.DoorRect.Moved(data.DoorPos)) {
				doorOpen = true
				break
			}
		}
	}
	if doorOpen {
		data.DoorEntity.AddComponent(myecs.Drawable, img.NewSprite(data.SpriteKeyDoorOpen, data.BatchKeyTest))
	} else {
		data.DoorEntity.AddComponent(myecs.Drawable, img.NewSprite(data.SpriteKeyDoorClosed, data.BatchKeyTest))
	}
}
