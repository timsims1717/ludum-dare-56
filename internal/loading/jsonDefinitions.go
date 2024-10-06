package loading

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
)

func LoadEnities() {
	entities, err := data.LoadEntityDefinitions("assets/EntityDefinitions.json")
	if err != nil {
		panic(err)
	}
	for _, value := range entities.DifficultyPool[data.Difficulty].StaticEntityPool {
		for i := 0; i < value.Weight; i++ {
			entities.StaticEntityPoolExpanded = append(entities.StaticEntityPoolExpanded, value.Name)
			entities.StaticEntityPoolTotal++
		}
	}
	for _, value := range entities.DynamicEntityPool {
		for i := 0; i < value.Weight; i++ {
			entities.DynamicEntityPoolExpanded = append(entities.DynamicEntityPoolExpanded, value.Name)
			entities.DynamicEntityPoolTotal++
		}
	}
	data.LoadedEntities = entities
}
