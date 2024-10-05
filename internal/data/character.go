package data

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

type Character struct {
	Object *object.Object
	Target *pixel.Vec
	Sprite *img.Sprite
}
