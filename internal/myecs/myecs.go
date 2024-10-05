package myecs

import (
	"github.com/bytearena/ecs"
	"github.com/timsims1717/pixel-go-utils/object"
)

var (
	FullCount   = 0
	IDCount     = 0
	LoadedCount = 0
)

type ClearFlag bool

var (
	Manager = ecs.NewManager()

	Object        = Manager.NewComponent()
	Parent        = Manager.NewComponent()
	Temp          = Manager.NewComponent()
	Update        = Manager.NewComponent()
	Interpolation = Manager.NewComponent()

	Drawable = Manager.NewComponent()
	Animated = Manager.NewComponent()

	Tile       = Manager.NewComponent()
	Border     = Manager.NewComponent()
	Block      = Manager.NewComponent()
	Character  = Manager.NewComponent()
	MoveTarget = Manager.NewComponent()
	Input      = Manager.NewComponent()
	Player     = Manager.NewComponent()
	Text       = Manager.NewComponent()
	PickUp     = Manager.NewComponent()

	IsTemp    = ecs.BuildTag(Temp)
	HasUpdate = ecs.BuildTag(Update)

	HasAnimation = ecs.BuildTag(Animated, Object)
	IsDrawable   = ecs.BuildTag(Drawable, Object)

	IsObject         = ecs.BuildTag(Object)
	HasParent        = ecs.BuildTag(Object, Parent)
	HasInterpolation = ecs.BuildTag(Object, Interpolation)

	IsTile      = ecs.BuildTag(Object, Tile)
	HasBorder   = ecs.BuildTag(Object, Border)
	IsBlock     = ecs.BuildTag(Object, Block)
	IsCharacter = ecs.BuildTag(Object, Character)
	IsNPC       = ecs.BuildTag(Object, Character, MoveTarget)
	IsPlayer    = ecs.BuildTag(Object, Character, Player, Input)
	IsPickUp    = ecs.BuildTag(Object, Character, PickUp)
	IsText      = ecs.BuildTag(Object, Drawable, Text)
)

func UpdateManager() {
	LoadedCount = 0
	IDCount = 0
	FullCount = 0
	for _, result := range Manager.Query(IsObject) {
		if t, ok := result.Components[Object].(*object.Object); ok {
			FullCount++
			if t.ID != "" {
				IDCount++
				if !t.Unloaded {
					LoadedCount++
				}
			}
		}
	}
}
