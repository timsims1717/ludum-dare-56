package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func CreateEntity() {
	entityRoll := PickRandomStaticEntity()
	obj := object.New().WithID(entityRoll.Name)
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Layer = 1
	spr := img.NewSprite(entityRoll.Sprite, data.TestBatchKey)
	character := &data.Character{
		Object: obj,
		Sprite: spr,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.StaticEnity, character)
}

func PickRandomStaticEntity() data.StaticEntity {
	roll := data.LoadedEnities.StaticEntities[data.LoadedEnities.StaticEnityPoolExpanded[data.GlobalSeededRandom.Intn(data.LoadedEnities.StaticEntityPoolTotal)]]
	return roll
}
