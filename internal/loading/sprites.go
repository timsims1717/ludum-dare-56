package loading

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/pixel-go-utils/img"
)

func LoadSprites() {
	testSheet, err := img.LoadSpriteSheet("assets/TestMap.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(data.BatchKeyTest, testSheet, true, true)
}
