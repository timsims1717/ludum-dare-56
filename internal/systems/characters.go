package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func CreateCharacter() {
	obj := object.New().WithID("test")
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	spr := img.NewSprite(data.GhostSpriteKey, data.TestBatchKey).WithOffset(pixel.V(0., 16.))
	character := &data.Character{
		Object: obj,
		Sprite: spr,
	}
	player := &data.Player{}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.Input, data.PlayerInput).
		AddComponent(myecs.Player, player)
	character.Entity = e
}

func CreateNPC() {
	RolledEnttity := PickRandomDynamicEntity()
	obj := object.New().WithID(RolledEnttity.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos.X = data.GetRandomX()
	obj.Pos.Y = data.GetRandomY()
	spr := img.NewSprite(RolledEnttity.Sprite, data.TestBatchKey)
	character := &data.Character{
		Object:   obj,
		Movement: data.Stationary,
		Target:   pixel.ZV,
		Sprite:   spr,
		HP:       RolledEnttity.HP,
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.PickUp, struct{}{})
	character.Entity = e
}

func PickRandomDynamicEntity() data.DynamicEntity {
	roll := data.LoadedEnities.DynamicEntities[data.LoadedEnities.DynamicEnityPoolExpanded[data.GlobalSeededRandom.Intn(data.LoadedEnities.DynamicEntityPoolTotal)]]
	return roll
}
