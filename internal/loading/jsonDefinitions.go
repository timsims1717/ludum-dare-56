package loading

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
)

func LoadEntities() {
	entities, err := data.LoadEntityDefinitions("assets/EntityDefinitions.json")
	if err != nil {
		panic(err)
	}
	entities.ExpandedEntityPools = make(map[string][]string)
	entities.ExpandedEntityTotals = make(map[string]int)
	entities.ExpandedEntityPools[data.DangerPool] = []string{}
	entities.ExpandedEntityPools[data.BabyPool] = []string{}
	entities.ExpandedEntityPools[data.ToyPool] = []string{}
	entities.ExpandedEntityTotals[data.DangerPool] = 0
	entities.ExpandedEntityTotals[data.BabyPool] = 0
	entities.ExpandedEntityTotals[data.ToyPool] = 0
	for _, value := range entities.DifficultyPool[data.Difficulty][data.DangerPool].EntityPool {
		for i := 0; i < value.Weight; i++ {
			entities.ExpandedEntityPools[data.DangerPool] = append(entities.ExpandedEntityPools[data.DangerPool], value.Name)
			entities.ExpandedEntityTotals[data.DangerPool]++
		}
	}
	for _, value := range entities.BabyPool {
		for i := 0; i < value.Weight; i++ {
			entities.ExpandedEntityPools[data.BabyPool] = append(entities.ExpandedEntityPools[data.BabyPool], value.Name)
			entities.ExpandedEntityTotals[data.BabyPool]++
		}
	}
	for _, value := range entities.DifficultyPool[data.Difficulty][data.ToyPool].EntityPool {
		for i := 0; i < value.Weight; i++ {
			entities.ExpandedEntityPools[data.ToyPool] = append(entities.ExpandedEntityPools[data.ToyPool], value.Name)
			entities.ExpandedEntityTotals[data.ToyPool]++
		}
	}
	for _, e := range entities.DynamicEntities {
		if e.Parent == "" {
			e.ParentText.DropOffText = append(e.ParentText.DropOffText, entities.DefaultText.DropOffText...)
			e.ParentText.PickUpText = append(e.ParentText.PickUpText, entities.DefaultText.PickUpText...)
			e.ParentText.SafeText = append(e.ParentText.SafeText, entities.DefaultText.SafeText...)
			e.ParentText.HurtText = append(e.ParentText.HurtText, entities.DefaultText.HurtText...)
			e.ParentText.DeadText = append(e.ParentText.DeadText, entities.DefaultText.DeadText...)
		}
	}
	data.LoadedEntities = entities
}
