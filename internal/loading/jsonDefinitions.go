package loading

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
)

func LoadEnities() {
	entities, err := data.LoadEntityDefinitions("assets/EntityDefinitions.json")
	if err != nil {
		panic(err)
	}
	data.Enities = entities
}
