package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/loading"
	"github.com/timsims1717/ludum-dare-56/internal/ui"
)

func LoadDialogs() {
	ui.NewDialog(loading.PauseConstructor)
}
