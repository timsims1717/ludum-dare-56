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
	TestBatchKey            = "test_batch"
	GhostSpriteKey          = "ghost_test"
	AntSpriteKey            = "ant_test"
	CactusSpriteKey         = "cactus_sprite"
	BearTrapSpriteKey       = "bear_trap_sprite"
	AggressiveVineSpriteKey = "aggressive_vine_sprite"
)

var (
	MainCanvas   *viewport.ViewPort
	CanvasWidth  = 640.
	CanvasHeight = 480.
)

const ( // player input
	InputLeft   = "left"
	InputRight  = "right"
	InputUp     = "up"
	InputDown   = "down"
	InputAction = "action"
)

var (
	TitleText          = "LD56"
	LoadedEnities      = new(EntityDefintions)
	GlobalSeededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandomTitle() string {

	return TitleText
}
