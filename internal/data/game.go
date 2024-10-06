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
	return GlobalSeededRandom.Float64()*300. - 150.
}

func GetRandomY() float64 {
	return GlobalSeededRandom.Float64()*200. - 100.
}
