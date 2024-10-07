package data

import (
	"github.com/timsims1717/pixel-go-utils/timing"
)

var (
	Layers       = true
	TheGamePhase = ParentDropOff
	ParentTimer  *timing.Timer
	PickUpTime   bool
	ThePlayer    *Character
	Parents      []*Character
	Kids         []*Character
	DropOffList  []string
	PickUpList   []int
	ParentIndex  int
	GameplayTime = 15.
)

type GamePhase int

const (
	ParentDropOff = iota
	Gameplay
	ParentPickUp
)

func GetRandomX() float64 {
	return RoomBorder.Center().X + GlobalSeededRandom.Float64()*(RoomBorder.W()-32) - (RoomBorder.W()-32)*0.5
}

func GetRandomY() float64 {
	return RoomBorder.Center().Y + GlobalSeededRandom.Float64()*(RoomBorder.H()-32) - (RoomBorder.H()-32)*0.5
}
