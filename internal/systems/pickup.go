package systems

import (
	"fmt"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	pxginput "github.com/timsims1717/pixel-go-input"
	"github.com/timsims1717/pixel-go-utils/object"
)

func PickUpSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsPlayer) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		_, okC := result.Components[myecs.Character].(*data.Character)
		in, okI := result.Components[myecs.Input].(*pxginput.Input)
		p, okP := result.Components[myecs.Player].(*data.Player)
		if okO && okC && okI && okP {
			if p.Held == nil {
				if in.Get(data.InputAction).Pressed() {
					for _, pickUp := range myecs.Manager.Query(myecs.IsPickUp) {
						objPU, okPO := pickUp.Components[myecs.Object].(*object.Object)
						_, okPC := pickUp.Components[myecs.Character].(*data.Character)
						if okPO && okPC {
							if obj.Rect.Moved(obj.Pos).Contains(objPU.Pos) {
								fmt.Println("pick up")
							}
						}
					}
				}
			}
		}
	}
}
