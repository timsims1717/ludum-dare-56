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

func NonPlayerCharacterSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsNPC) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC && !ch.PickedUp {
			switch ch.Movement {
			case data.Stationary:
				if ch.Timer == nil {
					ch.Timer = timing.New(data.GlobalRand.Float64()*3 + 0.5)
				}
				if ch.Timer.UpdateDone() {
					if data.GlobalRand.Intn(2) == 0 {
						ch.Timer = nil
						newPos := pixel.V(GetRandomX(), GetRandomY())
						count := 0
						for count < 8 {
							if util.Magnitude(obj.Pos.Sub(newPos)) > 20. {
								ch.Target = newPos
								ch.Movement = data.Target
								break
							}
							count++
						}
					} else {
						ch.Timer = timing.New(data.GlobalRand.Float64()*5 + 1.)
						newDir := util.Normalize(pixel.V(GetRandomX(), GetRandomY()))
						ch.Target = newDir
						ch.Movement = data.Random
					}
				}
			case data.Random:
				if ch.Timer.UpdateDone() {
					ch.Movement = data.Stationary
					ch.Target = pixel.ZV
				} else {
					obj.Pos.X += ch.Target.X * data.NPCSpeed * timing.DT
					obj.Pos.Y += ch.Target.Y * data.NPCSpeed * timing.DT
					ch.Target.X += (data.GlobalRand.Float64()*10. - 5.) * timing.DT
					ch.Target.Y += (data.GlobalRand.Float64()*10. - 5.) * timing.DT
					ch.Target = util.Normalize(ch.Target)
				}
			case data.Target:
				if ch.Target.X < obj.Pos.X {
					obj.Pos.X -= data.NPCSpeed * timing.DT
					obj.Flip = true
					if ch.Target.X > obj.Pos.X {
						obj.Pos.X = ch.Target.X
					}
				} else if ch.Target.X > obj.Pos.X {
					obj.Pos.X += data.NPCSpeed * timing.DT
					obj.Flip = false
					if ch.Target.X < obj.Pos.X {
						obj.Pos.X = ch.Target.X
					}
				}
				if ch.Target.Y < obj.Pos.Y {
					obj.Pos.Y -= data.NPCSpeed * timing.DT
					if ch.Target.Y > obj.Pos.Y {
						obj.Pos.Y = ch.Target.Y
					}
				} else if ch.Target.Y > obj.Pos.Y {
					obj.Pos.Y += data.NPCSpeed * timing.DT
					if ch.Target.Y < obj.Pos.Y {
						obj.Pos.Y = ch.Target.Y
					}
				}
				if ch.Target.X == obj.Pos.X && ch.Target.Y == obj.Pos.Y {
					ch.Movement = data.Stationary
					ch.Target = pixel.ZV
				}
			}
		}
	}
}
