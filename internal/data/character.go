package data

import (
	"github.com/bytearena/ecs"
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
	TargetNoStop
	Random
	Stationary
	Waiting
)

type Character struct {
	Entity            *ecs.Entity
	Object            *object.Object
	Movement          Movement
	Target            pixel.Vec
	Sprite            *img.Sprite
	Horiz             Direction
	Vert              Direction
	Timer             *timing.Timer
	PickedUp          bool
	HP                int
	Speed             float64
	InvinciblityTimer *timing.Timer
	IsInvincible      bool
	Damage            int
}

type Player struct {
	Held *Character
}

type KidParent struct {
	KidsDropped     int
	DropOffComplete bool
	PickUpComplete  bool
}

var (
	PlayerSpeed = 100.
	NPCSpeed    = 45.
)
