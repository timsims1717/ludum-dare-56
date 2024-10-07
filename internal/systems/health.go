package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func HealthBarSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsCharacter) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC && ch.HP > -1 && ch.MaxHP > 0 {
			for i := 1; i <= ch.MaxHP; i++ {
				xOff := 0.
				if ch.MaxHP%2 == 0 {
					xOff = float64(i*10) - float64(ch.MaxHP-1)*10/2
				} else {
					xOff = float64(i*10) - float64(ch.MaxHP)*10/2
				}
				spr := data.SpriteHeartFull
				if i > ch.HP {
					spr = data.SpriteHeartEmpty
				}
				heartObj := object.New()
				heartObj.Layer = 3
				heartObj.Offset = pixel.V(xOff, obj.Rect.H()*0.5+2)
				myecs.Manager.NewEntity().
					AddComponent(myecs.Object, heartObj).
					AddComponent(myecs.Drawable, img.NewSprite(spr, data.BatchKeyTest)).
					AddComponent(myecs.Parent, obj).
					AddComponent(myecs.Temp, myecs.ClearFlag(true))
			}
		}
	}
}
