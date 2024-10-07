package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/img"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/typeface"
	"golang.org/x/image/colornames"
	"math"
)

func SetTextBubble(ch *data.Character, raw string, xWidth, yOff float64) {
	if ch.TextBubble == nil {
		CreateTextBubble(ch, raw, xWidth, yOff)
	} else {
		ch.TextBubble.Text.SetText(raw)
		ch.TextBubble.MiddleObj.Offset = pixel.V(xWidth+ch.TextBubble.Text.Width*0.5, yOff)
		ch.TextBubble.MiddleObj.Sca.X = math.Max(ch.TextBubble.Text.Width*0.5/8., 1.)
		ch.TextBubble.LeftObj.Offset = pixel.V(ch.TextBubble.Text.Width*-0.5, 0.)
		ch.TextBubble.RightObj.Offset = pixel.V(ch.TextBubble.Text.Width*0.5, 0.)
		ch.TextBubble.Text.Hidden = false
		ch.TextBubble.LeftObj.Hidden = false
		ch.TextBubble.RightObj.Hidden = false
		ch.TextBubble.MiddleObj.Hidden = false
	}
}

func CreateTextBubble(ch *data.Character, raw string, xWidth, yOff float64) {
	txt := typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.0625, 0, 0)
	//txt.SetPos(pixel.V(6., 10.))
	txt.Obj.Layer = 5
	txt.Obj.Offset = pixel.V(xWidth, yOff+7.)
	txt.SetColor(pixel.ToRGBA(colornames.Black))
	txt.SetText(raw)
	te := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, txt.Obj).
		AddComponent(myecs.Drawable, txt).
		AddComponent(myecs.Parent, ch.Object)

	bmObj := object.New()
	bmObj.Layer = 4
	bmObj.Offset = pixel.V(xWidth+txt.Width*0.5, yOff)
	bmObj.Sca = pixel.V(math.Max(txt.Width*0.5/8., 1.), 1.)
	me := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, bmObj).
		AddComponent(myecs.Drawable, img.NewSprite(data.SpriteTextMiddle, data.BatchKeyTest)).
		AddComponent(myecs.Parent, ch.Object)

	blObj := object.New()
	blObj.Layer = 4
	blObj.Offset = pixel.V(txt.Width*-0.5, 0.)
	le := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, blObj).
		AddComponent(myecs.Drawable, img.NewSprite(data.SpriteTextLeft, data.BatchKeyTest)).
		AddComponent(myecs.Parent, bmObj)

	brObj := object.New()
	brObj.Layer = 4
	brObj.Offset = pixel.V(txt.Width*0.5, 0.)
	re := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, brObj).
		AddComponent(myecs.Drawable, img.NewSprite(data.SpriteTextRight, data.BatchKeyTest)).
		AddComponent(myecs.Parent, bmObj)

	ch.TextBubble = &data.TextBubble{
		Text:       txt,
		TextEntity: te,
		LeftObj:    blObj,
		Left:       le,
		MiddleObj:  bmObj,
		Middle:     me,
		RightObj:   brObj,
		Right:      re,
	}
}

func HideTextBubble(ch *data.Character) {
	if ch.TextBubble != nil {
		ch.TextBubble.Text.Hidden = true
		ch.TextBubble.LeftObj.Hidden = true
		ch.TextBubble.RightObj.Hidden = true
		ch.TextBubble.MiddleObj.Hidden = true
	}
}

func DisposeTextBubble(ch *data.Character) {
	if ch.TextBubble != nil {
		myecs.Manager.DisposeEntity(ch.TextBubble.TextEntity)
		myecs.Manager.DisposeEntity(ch.TextBubble.Left)
		myecs.Manager.DisposeEntity(ch.TextBubble.Middle)
		myecs.Manager.DisposeEntity(ch.TextBubble.Right)
		ch.TextBubble = nil
	}
}

func TextBubbleSystem() {

}
