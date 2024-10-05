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
	for _, value := range entities.DynamicEntityPool {
		for i := 0; i < value.Weight; i++ {
			entities.DynamicEnityPoolExpanded = append(entities.DynamicEnityPoolExpanded, value.Name)
			entities.DynamicEntityPoolTotal++
		}
	}
	data.LoadedEnities = entities
}
