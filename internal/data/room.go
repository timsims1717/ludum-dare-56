package data

import (
	"github.com/bytearena/ecs"
	"github.com/gopxl/pixel"
)

const (
	WoodWidth      = 16.
	WoodLeftIndex  = -18
	WoodRightIndex = 19
	WoodHeight     = 24
	RugWidth       = 48.
	RugLeftIndex   = -3
	RugRightIndex  = 4
	RugBottomIndex = -3
	RugTopIndex    = 2
)

var (
	RoomBorder    = pixel.R(-300., -224., 300., 160.)
	ParentPos     = pixel.V(0., 256.)
	WaitAtDoorPos = pixel.V(0., 168.)
	InRoomPos     = pixel.V(0., 132.)
	DoorPos       = pixel.V(0., 188.)
	DoorRect      = pixel.R(-8., -8., 8., 8.)
	MatPos        = pixel.V(0., 132.)
	MatRect       = pixel.R(-24., -16., 24., 16.)
	RoomBottom    = -220.

	DoorEntity *ecs.Entity
)
