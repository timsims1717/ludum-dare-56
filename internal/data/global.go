package data

import (
	"github.com/timsims1717/pixel-go-utils/viewport"
	"math/rand"
	"time"
)

const ( // states
	GameStateKey = "game_state"
)

const ( // sprites and batches
	BatchKeyTest        = "test_batch"
	SpriteKeyGhost      = "ghost_test"
	SpriteKeyMat        = "mat"
	SpriteKeyDoorClosed = "door_closed"
	SpriteKeyDoorOpen   = "door_open"
	SpriteHeartFull     = "heart_full"
	SpriteHeartEmpty    = "heart_empty"
	SpriteTextLeft      = "text_bubble_left"
	SpriteTextMiddle    = "text_bubble_middle"
	SpriteTextRight     = "text_bubble_right"
)

var (
	MainCanvas   *viewport.ViewPort
	CanvasWidth  = 740.
	CanvasHeight = 480.
)

const ( // player input
	InputLeft   = "left"
	InputRight  = "right"
	InputUp     = "up"
	InputDown   = "down"
	InputAction = "action"
)

const (
	InvincibilityDuration = 2
)

var (
	TitleText          = "LD56"
	LoadedEntities     = new(EntityDefinitions)
	GlobalSeededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	Difficulty         = "medium"
)

func RandomTitle() string {

	return TitleText
}
