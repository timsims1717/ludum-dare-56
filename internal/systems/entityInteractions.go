package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/object"
	"github.com/timsims1717/pixel-go-utils/timing"
)

func EntityInteractions() {
	for _, result := range myecs.Manager.Query(myecs.HasMoveTarget) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC {
			for _, result2 := range myecs.Manager.Query(myecs.IsStaticEntity) {
				obj2, okO2 := result2.Components[myecs.Object].(*object.Object)
				hitEntity, okC2 := result2.Components[myecs.Character].(*data.Character)
				if okO2 && okC2 {
					if obj.Rect.Moved(obj.Pos).Intersects(obj2.Rect.Moved(obj2.Pos)) && !ch.IsInvincible &&
						ch.HP >= 0 && !result.Entity.HasComponent(myecs.Parent) && !result2.Entity.HasComponent(myecs.Parent) {
						if hitEntity.StaticEnityProperties.Damage > 0 {
							ch.HP = ch.HP - hitEntity.StaticEnityProperties.Damage
							ch.Entity.AddComponent(myecs.Invincible, struct{}{})
							ch.InvinciblityTimer = timing.New(data.InvincibilityDuration)
							ch.IsInvincible = true
							if ch.HP <= 0 {
								myecs.Manager.DisposeEntity(result.Entity)
							}
						}
						if hitEntity.StaticEnityProperties.HasStatusEffect {
							if ch.StatusEffects == nil {
								ch.StatusEffects = make(map[string]data.StatusEffect)
							}
							ch.StatusEffects[hitEntity.StaticEnityProperties.StatusEffect] =
								data.LoadedStatuses.StatusEffects[hitEntity.StaticEnityProperties.StatusEffect].Clone()
							ch.Entity.AddComponent(myecs.StatusEffect, struct{}{})
							data.SpeedBoostApply(ch, hitEntity.StaticEnityProperties.StatusEffect)
						}

						if hitEntity.StaticEnityProperties.Uses > 0 {
							hitEntity.StaticEnityProperties.Uses--
							if hitEntity.StaticEnityProperties.Uses == 0 {
								myecs.Manager.DisposeEntity(result2.Entity)
							}
						}
					}
				}
			}
		}
	}
}
