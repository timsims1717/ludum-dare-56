package data

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/pixel-go-utils/timing"
)

var (
	RoomBorder   = pixel.R(-300., -220., 300., 176.)
	ParentPos    = pixel.V(0., 256.)
	DoorPos      = pixel.V(0., 184.)
	InRoomPos    = pixel.V(0., 148.)
	Layers       = true
	DropOffTimer *timing.Timer
	ThePlayer    *Character
	Parents      []*Character
	Kids         []*Character
	DropOffList  []string
	DropOffIndex int
)

func GetRandomX() float64 {
	return GlobalSeededRandom.Float64()*300. - 150.
}

func GetRandomY() float64 {
	return GlobalSeededRandom.Float64()*200. - 100.
}
