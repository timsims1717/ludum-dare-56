package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/pixel-go-utils/timing"
	"github.com/timsims1717/pixel-go-utils/util"
)

func DropOffSystem() {
	if data.DropOffTimer == nil || len(data.Parents) == 0 { // the level is starting
		CreateDropOffList()
		data.DropOffTimer = timing.New(0.)
	}
	if data.DropOffTimer.UpdateDone() &&
		len(data.Parents) < len(data.DropOffList) &&
		data.DropOffIndex < len(data.DropOffList) {
		CreateParentAndKids(data.DropOffList[data.DropOffIndex])
		data.DropOffIndex++
		data.DropOffTimer = timing.New(data.GlobalSeededRandom.Float64()*6. + 6.)
	}
}

func CreateDropOffList() {
	var a []int
	for i := 0; i < len(data.LoadedEntities.DynamicEntityPool); i++ {
		a = append(a, i)
	}
	dropOffInts := util.RandomSample(len(data.LoadedEntities.DynamicEntityPool), a, data.GlobalSeededRandom)
	data.DropOffList = []string{}
	for _, i := range dropOffInts {
		data.DropOffList = append(data.DropOffList, data.LoadedEntities.DynamicEntityPoolExpanded[i])
	}
	data.DropOffIndex = 0
}
