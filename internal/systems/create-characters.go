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
	spr := img.NewSprite(data.SpriteKeyGhost, data.BatchKeyTest).WithOffset(pixel.V(0., 16.))
	character := &data.Character{
		Object: obj,
		Sprite: spr,
		InRoom: true,
	}
	player := &data.Player{}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.Input, data.PlayerInput).
		AddComponent(myecs.Player, player)
	character.Entity = e
	data.ThePlayer = character
}

func CreateRandomKid() *data.Character {
	entity := data.PickRandomDynamicEntity()
	obj := object.New().WithID(entity.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos.X = GetRandomX()
	obj.Pos.Y = GetRandomY()
	spr := img.NewSprite(entity.Sprite, data.BatchKeyTest)
	character := &data.Character{
		Object:     obj,
		Movement:   data.Stationary,
		Target:     pixel.ZV,
		TargetDist: 16.,
		Sprite:     spr,
		HP:         entity.HP,
		Speed:      entity.Speed,
		Kid:        &data.Kid{},
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.PickUp, struct{}{}).
		AddComponent(myecs.Kid, character.Kid).
		AddComponent(myecs.Collide, struct{}{})
	character.Entity = e
	data.Kids = append(data.Kids, character)
	return character
}

func CreateKid(entity *data.DynamicEntity, pos pixel.Vec) *data.Character {
	obj := object.New().WithID(entity.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos = pos
	spr := img.NewSprite(entity.Sprite, data.BatchKeyTest)
	character := &data.Character{
		Object:     obj,
		Movement:   data.Stationary,
		Target:     pixel.ZV,
		TargetDist: 16.,
		Sprite:     spr,
		HP:         entity.HP,
		Speed:      entity.Speed,
		Kid:        &data.Kid{},
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.PickUp, struct{}{}).
		AddComponent(myecs.Kid, character.Kid).
		AddComponent(myecs.Collide, struct{}{})
	character.Entity = e
	data.Kids = append(data.Kids, character)
	return character
}

func CreateParent(entity string, kidCount int, pos pixel.Vec) *data.Character {
	parent := data.LoadedEntities.DynamicEntities[entity]
	obj := object.New().WithID(parent.Name)
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 32., 32.))
	obj.Pos = pos
	spr := img.NewSprite(parent.Sprite, data.BatchKeyTest).WithOffset(pixel.V(0., 16.))
	character := &data.Character{
		Object:   obj,
		Movement: data.Stationary,
		Target:   pixel.ZV,
		Sprite:   spr,
		HP:       parent.HP,
		Speed:    parent.Speed,
		KidParent: &data.KidParent{
			KidsDropped: kidCount,
		},
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.KidParent, character.KidParent)
	character.Entity = e
	data.Parents = append(data.Parents, character)
	return character
}

func CreateParentAndKids(entityName string) {
	entity := data.LoadedEntities.DynamicEntities[entityName]
	r := entity.Max - entity.Min
	count := 1
	if r > 0 {
		count = data.GlobalSeededRandom.Intn(r) + entity.Min
	}
	parent := CreateParent(entity.Parent, count, data.ParentPos)
	parent.Target = data.DoorPos
	parent.Movement = data.Straight
	parent.NoStop = true
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
		pos.X += data.GlobalSeededRandom.Float64()*12. - 6.
		pos.Y += data.GlobalSeededRandom.Float64()*12. - 6.
		kid := CreateKid(entity, pos)
		kid.Target = data.InRoomPos
		kid.Movement = data.Straight
		kid.NoStop = true
		kid.Kid.KidParent = parent
		parent.KidParent.Kids = append(parent.KidParent.Kids, kid)
	}
}

func GetRandomX() float64 {
	return data.GlobalSeededRandom.Float64()*300. - 150.
}

func GetRandomY() float64 {
	return data.GlobalSeededRandom.Float64()*200. - 100.
}
