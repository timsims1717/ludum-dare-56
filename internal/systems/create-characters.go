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

func CreateRandomKid() *data.Character {
	entity := data.PickRandomDynamicEntity()
	obj := object.New().WithID(entity.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos.X = GetRandomX()
	obj.Pos.Y = GetRandomY()
	spr := img.NewSprite(entity.Sprite, data.TestBatchKey)
	character := &data.Character{
		Object:   obj,
		Movement: data.Stationary,
		Target:   pixel.ZV,
		Sprite:   spr,
		HP:       entity.HP,
		Speed:    entity.Speed,
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.PickUp, struct{}{}).
		AddComponent(myecs.Kid, struct{}{})
	character.Entity = e
	return character
}

func CreateKid(entity *data.DynamicEntity, pos pixel.Vec) *data.Character {
	obj := object.New().WithID(entity.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos = pos
	spr := img.NewSprite(entity.Sprite, data.TestBatchKey)
	character := &data.Character{
		Object:   obj,
		Movement: data.Stationary,
		Target:   pixel.ZV,
		Sprite:   spr,
		HP:       entity.HP,
		Speed:    entity.Speed,
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.PickUp, struct{}{}).
		AddComponent(myecs.Kid, struct{}{})
	character.Entity = e
	return character
}

func CreateParent(entity string, kidCount int, pos pixel.Vec) *data.Character {
	parent := data.LoadedEnities.DynamicEntities[entity]
	obj := object.New().WithID(parent.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos = pos
	spr := img.NewSprite(parent.Sprite, data.TestBatchKey).WithOffset(pixel.V(0., 16.))
	character := &data.Character{
		Object:   obj,
		Movement: data.Stationary,
		Target:   pixel.ZV,
		Sprite:   spr,
		HP:       parent.HP,
		Speed:    parent.Speed,
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.KidParent, &data.KidParent{
			KidsDropped: kidCount,
		})
	character.Entity = e
	return character
}

func CreateParentAndKids() {
	entity := data.PickRandomDynamicEntity()
	count := data.GlobalSeededRandom.Intn(entity.Max-entity.Min) + entity.Min
	parent := CreateParent(entity.Parent, count, data.ParentPos)
	parent.Target = data.DoorPos
	parent.Movement = data.TargetNoStop
	for i := 0; i < count; i++ {
		pos := data.ParentPos
		switch i % 3 {
		case 0:
			pos.X -= 38.
			pos.Y += float64((i+1)/3) * 38.
		case 1:
			pos.X += 38.
			pos.Y += float64((i+1)/3) * 38.
		case 2:
			pos.Y += 68. + float64((i+1)/3)*38.
		}
		kid := CreateKid(entity, pos)
		kid.Target = data.DoorPos
		kid.Movement = data.TargetNoStop
	}
}

func GetRandomX() float64 {
	return data.GlobalSeededRandom.Float64()*300. - 150.
}

func GetRandomY() float64 {
	return data.GlobalSeededRandom.Float64()*200. - 100.
}
