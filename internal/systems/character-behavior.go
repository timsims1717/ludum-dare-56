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
		kid, okK := result.Components[myecs.Kid].(*data.Kid)
		if okO && okC && okK && !result.Entity.HasComponent(myecs.Parent) {
			if kid.DroppedOff && !kid.PickedUp {
				if ch.Movement == data.Stationary {
					if kid.KidParent.KidParent.ParentState == data.PickingUp {
						if data.MatRect.Moved(data.MatPos).Contains(obj.Pos) {
							kid.PickedUp = true
							ch.Target = data.ParentPos
							ch.Target.X += data.GlobalSeededRandom.Float64()*64. - 32.
							ch.Target.Y += data.GlobalSeededRandom.Float64() * 64.
							ch.Movement = data.Straight
							ch.NoStop = true
							ch.InRoom = false
							obj.Layer = -1
							result.Entity.RemoveComponent(myecs.Collide)
						}
					}
					if !kid.PickedUp {
						if ch.Timer == nil {
							ch.Timer = timing.New(data.GlobalSeededRandom.Float64()*3 + 0.5)
						}
						if ch.Timer.UpdateDone() {
							ChangeKidMovement(ch, obj)
						}
					}
				}
			} else if !kid.DroppedOff {
				if ch.Movement == data.Stationary {
					ch.InRoom = true
					obj.Layer = 1
					ch.NoStop = false
					kid.DroppedOff = true
					ChangeKidMovement(ch, obj)
				}
			}
		}
	}
}

func ChangeKidMovement(ch *data.Character, obj *object.Object) {
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

func KidParentBehaviorSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsKidParent) {
		_, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		par, okP := result.Components[myecs.KidParent].(*data.KidParent)
		if okO && okC && okP {
			if ch.Movement == data.Stationary {
				switch par.ParentState {
				case data.TimeToDropOff:
					dropComplete := true
					for _, kid := range par.Kids {
						if !kid.Kid.DroppedOff {
							dropComplete = false
							break
						}
					}
					if dropComplete {
						ch.Target = data.ParentPos
						ch.Movement = data.Straight
						ch.NoStop = true
						par.ParentState = data.DropOffComplete
					}
				case data.TimeToPickUp:
					par.ParentState = data.PickingUp
				case data.PickingUp:
					pickUpComplete := true
					for _, kid := range par.Kids {
						if !kid.Kid.PickedUp {
							pickUpComplete = false
							break
						}
					}
					if pickUpComplete {
						ch.Target = data.ParentPos
						ch.Movement = data.Straight
						ch.NoStop = true
						par.ParentState = data.PickUpComplete
						for _, kid := range par.Kids {
							kid.Target = data.ParentPos
							kid.Target.X += data.GlobalSeededRandom.Float64()*64. - 32.
							kid.Target.Y += data.GlobalSeededRandom.Float64()*64. - 32.
							kid.Movement = data.Straight
							kid.NoStop = true
						}
					}
				}
			}
		}
	}
}
