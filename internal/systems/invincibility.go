package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/myecs"
	"github.com/timsims1717/pixel-go-utils/object"
)

func UpdateInvincibility() {
	for _, result := range myecs.Manager.Query(myecs.IsInvincible) {
		_, okO := result.Components[myecs.Object].(*object.Object)
		ch, okC := result.Components[myecs.Character].(*data.Character)
		if okO && okC {
			ch.IsInvincible = !ch.InvinciblityTimer.UpdateDone()
			if !ch.IsInvincible {
				ch.Entity.RemoveComponent(myecs.Invincible)
			}
		}
	}
}
