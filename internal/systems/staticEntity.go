package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func CreateEntity() {
	obj := object.New().WithID("cactus")
	obj.Layer = 1
	entityRoll := PickRandomEntity()
	spr := img.NewSprite(data.LoadedEnities.StaticEntities[entityRoll.Name].Sprite, data.TestBatchKey)
	character := &data.Character{
		Object: obj,
		Sprite: spr,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character)
}

func PickRandomEntity() data.EntityRoll {
	roll := data.LoadedEnities.StaticEntityPool[data.GlobalSeededRandom.Intn(len(data.LoadedEnities.StaticEntityPool))]
	return roll
}
