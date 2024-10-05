package data

import "github.com/timsims1717/pixel-go-utils/viewport"

const ( // states
	GameStateKey = "game_state"
)

const ( // sprites and batches
	TestBatchKey   = "test_batch"
	GhostSpriteKey = "ghost_test"
	AntSpriteKey   = "ant_test"
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
