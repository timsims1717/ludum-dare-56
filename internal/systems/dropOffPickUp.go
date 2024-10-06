package systems

import (
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/pixel-go-utils/timing"
	"github.com/timsims1717/pixel-go-utils/util"
)

func InitGameplay() {
	CreateDropOffList(len(data.LoadedEntities.DynamicEntityPool))
	data.TheGamePhase = data.ParentDropOff
	data.ParentTimer = nil

}

func CreateDropOffList(count int) {
	var a []int
	for i := 0; i < count; i++ {
		a = append(a, i)
	}
	dropOffInts := util.RandomSample(count, a, data.GlobalSeededRandom)
	data.DropOffList = []string{}
	for _, i := range dropOffInts {
		data.DropOffList = append(data.DropOffList, data.LoadedEntities.DynamicEntityPoolExpanded[i])
	}
	data.ParentIndex = 0
}

func CreatePickUpList() {
	var a []int
	for i := 0; i < len(data.Parents); i++ {
		a = append(a, i)
	}
	data.PickUpList = util.RandomSample(len(data.Parents), a, data.GlobalSeededRandom)
	data.ParentIndex = 0
}

func DropOffSystem() {
	if data.ParentTimer.UpdateDone() &&
		len(data.Parents) < len(data.DropOffList) &&
		data.ParentIndex < len(data.DropOffList) {
		CreateParentAndKids(data.DropOffList[data.ParentIndex])
		data.ParentIndex++
		data.ParentTimer = timing.New(data.GlobalSeededRandom.Float64()*4. + 4.)
	}
	if len(data.Parents) == len(data.DropOffList) {
		doneDroppingOff := true
		for _, parent := range data.Parents {
			if parent.KidParent.ParentState == data.TimeToDropOff {
				doneDroppingOff = false
				break
			}
		}
		if doneDroppingOff {
			data.TheGamePhase = data.Gameplay
			data.ParentTimer = timing.New(data.GameplayTime)
		}
	}
}

func GameplaySystem() {
	if data.ParentTimer.UpdateDone() {
		data.TheGamePhase = data.ParentPickUp
		data.ParentTimer = nil
		CreatePickUpList()
	}
}

func ParentPickUpSystem() {
	pickUpReady := true
	allDone := true
	for _, parent := range data.Parents {
		if parent.KidParent.ParentState != data.PickUpComplete {
			allDone = false
		}
		if parent.KidParent.ParentState == data.TimeToPickUp ||
			parent.KidParent.ParentState == data.PickingUp {
			pickUpReady = false
			break
		}
	}
	if allDone {
		return
	}
	if pickUpReady {
		nextParent := data.Parents[data.PickUpList[data.ParentIndex]]
		nextParent.KidParent.ParentState = data.TimeToPickUp
		nextParent.Target = data.DoorPos
		nextParent.Movement = data.Straight
		nextParent.NoStop = true
		data.ParentIndex++
	}
}
