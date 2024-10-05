package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	pxginput "github.com/timsims1717/pixel-go-input"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/timing"
)

func PlayerCharacterSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsPlayer) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		in, okI := result.Components[myecs.Input].(*pxginput.Input)
		if okO && okC && okI {
			if in.Get(data.InputLeft).JustPressed() {
				ch.Horiz = data.Left
			} else if in.Get(data.InputRight).JustPressed() {
				ch.Horiz = data.Right
			} else if !in.Get(data.InputLeft).Pressed() && !in.Get(data.InputRight).Pressed() {
				ch.Horiz = data.NoDirection
			}
			if in.Get(data.InputUp).JustPressed() {
				ch.Vert = data.Up
			} else if in.Get(data.InputDown).JustPressed() {
				ch.Vert = data.Down
			} else if !in.Get(data.InputUp).Pressed() && !in.Get(data.InputDown).Pressed() {
				ch.Vert = data.NoDirection
			}
			if ch.Horiz == data.Left && in.Get(data.InputLeft).Pressed() {
				obj.Pos.X -= data.PlayerSpeed * timing.DT
			} else if ch.Horiz == data.Right && in.Get(data.InputRight).Pressed() {
				obj.Pos.X += data.PlayerSpeed * timing.DT
			} else if in.Get(data.InputLeft).Pressed() {
				obj.Pos.X -= data.PlayerSpeed * timing.DT
			} else if in.Get(data.InputRight).Pressed() {
				obj.Pos.X += data.PlayerSpeed * timing.DT
			}
			if ch.Horiz == data.Down && in.Get(data.InputDown).Pressed() {
				obj.Pos.Y -= data.PlayerSpeed * timing.DT
			} else if ch.Horiz == data.Up && in.Get(data.InputUp).Pressed() {
				obj.Pos.Y += data.PlayerSpeed * timing.DT
			} else if in.Get(data.InputDown).Pressed() {
				obj.Pos.Y -= data.PlayerSpeed * timing.DT
			} else if in.Get(data.InputUp).Pressed() {
				obj.Pos.Y += data.PlayerSpeed * timing.DT
			}
			if ch.Horiz == data.Left {
				obj.Flip = true
			} else if ch.Horiz == data.Right {
				obj.Flip = false
			}
		}
	}
}

//func NonPlayerCharacterSystem() {
//	for _, result := range myecs.Manager.Query(myecs.IsNPC) {
//		obj, okO := result.Components[myecs.Object].(*object.Object)
//		ch, okC := result.Components[myecs.Character].(*data.Character)
//		in, okI := result.Components[myecs.Input].(*pxginput.Input)
//		if okO && okC && okI {
//			if in.Get(data.InputLeft).JustPressed() {
//				ch.Horiz = data.Left
//			} else if in.Get(data.InputRight).JustPressed() {
//				ch.Horiz = data.Right
//			} else if !in.Get(data.InputLeft).Pressed() && !in.Get(data.InputRight).Pressed() {
//				ch.Horiz = data.NoDirection
//			}
//			if in.Get(data.InputUp).JustPressed() {
//				ch.Vert = data.Up
//			} else if in.Get(data.InputDown).JustPressed() {
//				ch.Vert = data.Down
//			} else if !in.Get(data.InputUp).Pressed() && !in.Get(data.InputDown).Pressed() {
//				ch.Vert = data.NoDirection
//			}
//			if ch.Horiz == data.Left && in.Get(data.InputLeft).Pressed() {
//				obj.Pos.X -= data.PlayerSpeed * timing.DT
//			} else if ch.Horiz == data.Right && in.Get(data.InputRight).Pressed() {
//				obj.Pos.X += data.PlayerSpeed * timing.DT
//			} else if in.Get(data.InputLeft).Pressed() {
//				obj.Pos.X -= data.PlayerSpeed * timing.DT
//			} else if in.Get(data.InputRight).Pressed() {
//				obj.Pos.X += data.PlayerSpeed * timing.DT
//			}
//			if ch.Horiz == data.Down && in.Get(data.InputDown).Pressed() {
//				obj.Pos.Y -= data.PlayerSpeed * timing.DT
//			} else if ch.Horiz == data.Up && in.Get(data.InputUp).Pressed() {
//				obj.Pos.Y += data.PlayerSpeed * timing.DT
//			} else if in.Get(data.InputDown).Pressed() {
//				obj.Pos.Y -= data.PlayerSpeed * timing.DT
//			} else if in.Get(data.InputUp).Pressed() {
//				obj.Pos.Y += data.PlayerSpeed * timing.DT
//			}
//		}
//	}
//}
