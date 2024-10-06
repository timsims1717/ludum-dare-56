package systems

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/timing"
	"github.com/timsims1717/pixel-go-utils/util"
)

func KidBehaviorSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsKid) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC {
			if ch.Movement == data.Stationary {
				if ch.Timer == nil {
					ch.Timer = timing.New(data.GlobalSeededRandom.Float64()*3 + 0.5)
				}
				if ch.Timer.UpdateDone() {
					if data.GlobalSeededRandom.Intn(2) == 0 {
						ch.Timer = timing.New(data.GlobalSeededRandom.Float64()*7. + 1.)
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
						ch.Timer = timing.New(data.GlobalSeededRandom.Float64()*5. + 1.)
						newDir := util.Normalize(pixel.V(GetRandomX(), GetRandomY()))
						ch.Target = newDir
						ch.Movement = data.Random
					}
				}
			}
		}
	}
}

func KidParentBehaviorSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsKid) {
		_, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		par, okP := result.Components[myecs.KidParent].(*data.KidParent)
		if okO && okC && okP {
			if ch.Movement == data.Stationary {
				if !par.DropOffComplete {
					ch.Target = data.ParentPos
					ch.Movement = data.TargetNoStop
				}
			}
		}
	}
}
