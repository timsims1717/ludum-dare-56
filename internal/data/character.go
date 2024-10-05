package data

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/timing"
)

type Direction int

const (
	Left = iota
	Right
	Up
	Down
	NoDirection
)

type Movement int

const (
	Target = iota
	Random
	Stationary
)

type Character struct {
	Object   *object.Object
	Movement Movement
	Target   pixel.Vec
	Sprite   *img.Sprite
	Horiz    Direction
	Vert     Direction
	Timer    *timing.Timer
	PickedUp bool
}

type Player struct {
	Held *Character
}

var (
	PlayerSpeed = 100.
	NPCSpeed    = 45.
)
