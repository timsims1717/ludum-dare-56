package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	pxginput "github.com/timsims1717/pixel-go-input"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/timing"
	"github.com/timsims1717/pixel-go-utils/util"
)

func PlayerMoveSystem() {
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
			mov := pixel.ZV
			if ch.Horiz == data.Left && in.Get(data.InputLeft).Pressed() {
				mov.X = -1
			} else if ch.Horiz == data.Right && in.Get(data.InputRight).Pressed() {
				mov.X = 1
			} else if in.Get(data.InputLeft).Pressed() {
				mov.X = -1
			} else if in.Get(data.InputRight).Pressed() {
				mov.X = 1
			}
			if ch.Vert == data.Down && in.Get(data.InputDown).Pressed() {
				mov.Y = -1
			} else if ch.Vert == data.Up && in.Get(data.InputUp).Pressed() {
				mov.Y = 1
			} else if in.Get(data.InputDown).Pressed() {
				mov.Y = -1
			} else if in.Get(data.InputUp).Pressed() {
				mov.Y = 1
			}
			if ch.Horiz != data.NoDirection || ch.Vert != data.NoDirection {
				mov = util.Normalize(mov)
				obj.Pos.X += mov.X * data.PlayerSpeed * timing.DT
				obj.Pos.Y += mov.Y * data.PlayerSpeed * timing.DT
			}
			if ch.Horiz == data.Left {
				obj.Flip = true
			} else if ch.Horiz == data.Right {
				obj.Flip = false
			}
		}
	}
}

func NonPlayerMoveSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasMoveTarget) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC {
			if ch.PickedUp {
				ch.Movement = data.Stationary
				ch.Target = pixel.ZV
			} else {
				switch ch.Movement {
				case data.Random:
					if !ch.NoStop && ch.Timer.UpdateDone() {
						ch.Movement = data.Stationary
						ch.Target = pixel.ZV
						ch.Timer = nil
					} else {
						if ch.Target.X < 0 {
							obj.Flip = true
						} else if ch.Target.X > 0 {
							obj.Flip = false
						}
						obj.Pos.X += ch.Target.X * ch.Speed * timing.DT
						obj.Pos.Y += ch.Target.Y * ch.Speed * timing.DT
						ch.Target.X += (data.GlobalSeededRandom.Float64()*10. - 5.) * timing.DT
						ch.Target.Y += (data.GlobalSeededRandom.Float64()*10. - 5.) * timing.DT
						ch.Target = util.Normalize(ch.Target)
					}
				case data.Target:
					if !ch.NoStop && ch.Timer.UpdateDone() {
						ch.Movement = data.Stationary
						ch.Target = pixel.ZV
						ch.Timer = nil
					} else {
						mov := pixel.ZV
						horiz := data.NoDirection
						vert := data.NoDirection
						if ch.Target.X < obj.Pos.X {
							mov.X = -1
							obj.Flip = true
							horiz = data.Left
						} else if ch.Target.X > obj.Pos.X {
							mov.X = 1
							obj.Flip = false
							horiz = data.Right
						}
						if ch.Target.Y < obj.Pos.Y {
							mov.Y = -1
							vert = data.Down
						} else if ch.Target.Y > obj.Pos.Y {
							mov.Y = 1
							vert = data.Up
						}
						MoveInDir(ch, obj, mov, data.Direction(horiz), data.Direction(vert))
					}
				case data.Straight:
					if !ch.NoStop && ch.Timer.UpdateDone() {
						ch.Movement = data.Stationary
						ch.Target = pixel.ZV
						ch.Timer = nil
					} else {
						horiz := data.NoDirection
						vert := data.NoDirection
						if ch.Target.X < obj.Pos.X {
							obj.Flip = true
							horiz = data.Left
						} else if ch.Target.X > obj.Pos.X {
							obj.Flip = false
							horiz = data.Right
						}
						if ch.Target.Y < obj.Pos.Y {
							vert = data.Down
						} else if ch.Target.Y > obj.Pos.Y {
							vert = data.Up
						}
						mov := ch.Target.Sub(obj.Pos)
						MoveInDir(ch, obj, mov, data.Direction(horiz), data.Direction(vert))
					}
				}
			}
		}
	}
}

func MoveInDir(ch *data.Character, obj *object.Object, mov pixel.Vec, horiz, vert data.Direction) {
	if horiz != data.NoDirection || vert != data.NoDirection {
		mov = util.Normalize(mov)
		obj.Pos.X += mov.X * ch.Speed * timing.DT
		obj.Pos.Y += mov.Y * ch.Speed * timing.DT

		if horiz == data.Left && ch.Target.X > obj.Pos.X {
			obj.Pos.X = ch.Target.X
		} else if horiz == data.Right && ch.Target.X < obj.Pos.X {
			obj.Pos.X = ch.Target.X
		}
		if vert == data.Down && ch.Target.Y > obj.Pos.Y {
			obj.Pos.Y = ch.Target.Y
		} else if vert == data.Up && ch.Target.Y < obj.Pos.Y {
			obj.Pos.Y = ch.Target.Y
		}
	}
	if (ch.Target.X == obj.Pos.X && ch.Target.Y == obj.Pos.Y) ||
		(ch.TargetDist > util.Magnitude(ch.Target.Sub(obj.Pos))) {
		ch.Movement = data.Stationary
		ch.Target = pixel.ZV
		ch.Timer = nil
	}
}

func RoomBorderSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsCharacter) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC && !result.Entity.HasComponent(myecs.Parent) && ch.InRoom {
			if obj.Pos.X+obj.HalfWidth > data.RoomBorder.Max.X {
				obj.Pos.X = data.RoomBorder.Max.X - obj.HalfWidth
			} else if obj.Pos.X-obj.HalfWidth < data.RoomBorder.Min.X {
				obj.Pos.X = data.RoomBorder.Min.X + obj.HalfWidth
			}
			if obj.Pos.Y+obj.HalfHeight > data.RoomBorder.Max.Y {
				obj.Pos.Y = data.RoomBorder.Max.Y - obj.HalfHeight
			} else if obj.Pos.Y-obj.HalfHeight < data.RoomBorder.Min.Y {
				obj.Pos.Y = data.RoomBorder.Min.Y + obj.HalfHeight
			}
		}
	}
}

func NPCCollisions() {
	for i, result := range myecs.Manager.Query(myecs.IsCollide) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO && !result.Entity.HasComponent(myecs.Parent) {
			for j, result2 := range myecs.Manager.Query(myecs.IsCollide) {
				if j > i {
					obj2, okO2 := result2.Components[myecs.Object].(*object.Object)
					if okO2 && !result2.Entity.HasComponent(myecs.Parent) {
						d := obj.HalfWidth + obj2.HalfWidth
						v := obj.Pos.Sub(obj2.Pos)
						m := util.Magnitude(v)
						if m < d {
							p := (d - m) * 0.5
							n := util.Normalize(v).Scaled(p)
							obj.Pos = obj.Pos.Add(n)
							obj2.Pos = obj2.Pos.Add(pixel.V(-n.X, -n.Y))
						}
					}
				}
			}
		}
	}
}
