package data

import "github.com/gopxl/pixel"

const (
	WoodWidth      = 16.
	WoodLeftIndex  = -16
	WoodRightIndex = 8
	FloorHeight    = 25
	RugWidth       = 48.
	RugLeftIndex   = -2
	RugRightIndex  = 3
	RugBottomIndex = -2
	RugTopIndex    = 1
)

var (
	RoomBorder = pixel.R(-300., -220., 300., 176.)
	ParentPos  = pixel.V(0., 256.)
	DoorPos    = pixel.V(0., 184.)
	InRoomPos  = pixel.V(0., 148.)
	MatPos     = pixel.V(0., 148.)
	MatRect    = pixel.R(-24., -16., 24., 16.)
	RoomBottom = -220.
)
