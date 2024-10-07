package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/object"
)

func UpdateStatusEffects() {
	for _, result := range myecs.Manager.Query(myecs.HasStatusEffect) {
		_, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC {
			for _, effect := range ch.StatusEffects {
				if effect.StatusTimer.UpdateDone() {
					data.SpeedBoostClear(ch, effect.Name)
					delete(ch.StatusEffects, effect.Name)
					if len(ch.StatusEffects) == 0 {
						ch.Entity.RemoveComponent(myecs.StatusEffect)
					}
				}
			}
		}
	}
}
