package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
)

func CreateCharacter() {
	obj := object.New().WithID("test")
	obj.Layer = 1
	spr := img.NewSprite(data.GhostSpriteKey, data.TestBatchKey)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr)
}
