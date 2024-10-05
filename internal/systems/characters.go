package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/constants"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func CreateCharacter() {
	obj := object.New().WithID("test")
	obj.Layer = 1
	obj.SetRect(pixel.R(0., 0., 48., 48.))
	spr := img.NewSprite(constants.GhostSpriteKey, constants.TestBatchKey)
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
	obj := object.New().WithID("npc1")
	obj.Layer = 1
	obj.Pos.X = GetRandomX()
	obj.Pos.Y = GetRandomY()
	spr := img.NewSprite(constants.AntSpriteKey, constants.TestBatchKey)
	character := &data.Character{
		Object:   obj,
		Movement: data.Stationary,
		Target:   pixel.ZV,
		Sprite:   spr,
	}
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Character, character).
		AddComponent(myecs.MoveTarget, struct{}{}).
		AddComponent(myecs.PickUp, struct{}{})
	character.Entity = e
}

func GetRandomX() float64 {
	return data.GlobalRand.Float64()*300. - 150.
}

func GetRandomY() float64 {
	return data.GlobalRand.Float64()*200. - 100.
}
