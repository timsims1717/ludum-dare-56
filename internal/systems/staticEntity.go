package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/constants"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func CreateEntity() {
	obj := object.New().WithID("cactus")
	obj.Layer = 1
	spr := img.NewSprite(constants.AggressiveVineSpriteKey, constants.TestBatchKey)
	character := &data.Character{
		Object: obj,
		Sprite: spr,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character)
}

func PickRandomEntity() {

}
