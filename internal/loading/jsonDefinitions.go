package loading

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
)

func LoadEnities() {
	entities, err := data.LoadEntityDefinitions("assets/EntityDefinitions.json")
	if err != nil {
		panic(err)
	}
	for _, value := range entities.StaticEntityPool {
		for i := 0; i < value.Weight; i++ {
			entities.StaticEnityPoolExpanded = append(entities.StaticEnityPoolExpanded, value.Name)
			entities.StaticEntityPoolTotal++
		}
	}
	data.LoadedEnities = entities
}
