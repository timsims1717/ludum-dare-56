package data

import (
	"github.com/bytearena/ecs"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/typeface"
)

type TextBubble struct {
	Text      *typeface.Text
	LeftObj   *object.Object
	Left      *ecs.Entity
	MiddleObj *object.Object
	Middle    *ecs.Entity
	RightObj  *object.Object
	Right     *ecs.Entity
}
