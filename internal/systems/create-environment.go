package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func PopulateLandscape() {
	for i := 0; i < data.LoadedEntities.DifficultyPool[data.Difficulty][data.DangerPool].Rolls; i++ {
		CreateEntity(data.DangerPool)
	}
	for i := 0; i < data.LoadedEntities.DifficultyPool[data.Difficulty][data.ToyPool].Rolls; i++ {
		CreateEntity(data.ToyPool)
	}
}

func CreateEntity(PoolType string) {
	entityRoll := PickRandomStaticEntity(PoolType)
	obj := object.New().WithID(entityRoll.Name)
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Layer = 1
	obj.Pos.X = data.GetRandomX()
	obj.Pos.Y = data.GetRandomY()
	spr := img.NewSprite(entityRoll.Sprite, data.BatchKeyTest)
	character := &data.Character{
		Object: obj,
		Sprite: spr,
		Damage: entityRoll.Damage,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.StaticEnity, character)
}

func PickRandomStaticEntity(PoolType string) *data.StaticEntity {
	roll := data.LoadedEntities.StaticEntities[data.LoadedEntities.ExpandedEntityPools[PoolType][data.GlobalSeededRandom.Intn(data.LoadedEntities.ExpandedEntityTotals[PoolType])]]
	return roll
}
