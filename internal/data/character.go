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
	Straight
	Random
	Stationary
	Waiting
)

type Character struct {
	Entity                *ecs.Entity
	Object                *object.Object
	Movement              Movement
	Target                pixel.Vec
	TargetDist            float64
	NoStop                bool
	Sprite                *img.Sprite
	Horiz                 Direction
	Vert                  Direction
	InRoom                bool
	Timer                 *timing.Timer
	PickedUp              bool
	HP                    int
	MaxHP                 int
	Speed                 float64
	InvinciblityTimer     *timing.Timer
	IsInvincible          bool
	KidParent             *KidParent
	Kid                   *Kid
	TextBubble            *TextBubble
	TextBoxYOff           float64
	TextBoxXOff           float64
	StaticEnityProperties StaticEntity
}

type Player struct {
	Held *Character
}

type Kid struct {
	DroppedOff bool
	PickedUp   bool
	KidParent  *Character
}

type KidParent struct {
	KidsDropped int
	Kids        []*Character
	ParentState ParentState
	DropOffText []string
	PickUpText  []string
	SafeText    []string
	HurtText    []string
	DeadText    []string
}

type ParentState int

const (
	TimeToDropOff = iota
	DropOffComplete
	TimeToPickUp
	PickingUp
	PickUpComplete
)

var (
	PlayerSpeed = 100.
	NPCSpeed    = 45.
)
