package data

import (
	"github.com/timsims1717/pixel-go-utils/viewport"
	"image/color"
	"math/rand"
	"time"
)

const ( // states
	GameStateKey = "game_state"
)

const ( // sprites and batches
	BatchKeyTest                = "test_batch"
	SpriteKeyGhost              = "ghost_test"
	SpriteKeyMat                = "mat"
	SpriteKeyDoorClosed         = "door_closed"
	SpriteKeyDoorOpen           = "door_open"
	SpriteHeartFull             = "heart_full"
	SpriteHeartEmpty            = "heart_empty"
	SpriteTextLeft              = "text_bubble_left"
	SpriteTextMiddle            = "text_bubble_middle"
	SpriteTextRight             = "text_bubble_right"
	SpriteChalkboardCorner      = "chalkboard_corner"
	SpriteChalkboardSide        = "chalkboard_side"
	SpriteChalkboardWhite       = "chalkboard_white"
	SpriteChalkboardBlackSquare = "chalkboard_black_square"
)

const (
	DialogPause = "pause_menu"
)

var (
	MainCanvas   *viewport.ViewPort
	CanvasWidth  = 740.
	CanvasHeight = 480.
	Ratio        = 0.

	ChalkboardWhite = color.RGBA{
		R: 233,
		G: 233,
		B: 233,
		A: 255,
	}
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
	LoadedStatuses     = new(StatusDefinitions)
	GlobalSeededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	Difficulty         = "medium"
)

func RandomTitle() string {

	return TitleText
}
