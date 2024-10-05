package loading

import (
	"github.com/timsims1717/ludum-dare-56/internal/constants"
	"github.com/timsims1717/pixel-go-utils/img"
)

func LoadSprites() {
	testSheet, err := img.LoadSpriteSheet("assets/TestMap.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(constants.TestBatchKey, testSheet, true, true)
}
