package data

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/pixel-go-utils/timing"
)

var (
	RoomBorder   = pixel.R(-300., -220., 300., 180.)
	ParentPos    = pixel.V(0., 256.)
	DoorPos      = pixel.V(0., 196.)
	Layers       = true
	DropOffTimer *timing.Timer
)
