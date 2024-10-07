package data

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

const (
	DangerPool = "dangerpool"
	ToyPool    = "toypool"
	BabyPool   = "babypool"
)

type EntityDefinitions struct {
	StaticEntities       map[string]*StaticEntity              `json:"StaticEntities"`
	DifficultyPool       map[string]map[string]*DifficultyPool `json:"DifficultyPool"`
	DynamicEntities      map[string]*DynamicEntity             `json:"DynamicEntities"`
	BabyPool             []EntityRoll                          `json:"BabyPool"`
	DefaultText          ParentText                            `json:"defaultText"`
	ExpandedEntityPools  map[string][]string
	ExpandedEntityTotals map[string]int
}

type StaticEntity struct {
	Name         string  `json:"key"`
	Sprite       string  `json:"sprite"`
	Damage       int     `json:"damage"`
	Damagetype   string  `json:"damagetype"`
	IsCollidable bool    `json:"IsCollidable"`
	IsPickupable bool    `json:"IsPickupable"`
	IsPushable   bool    `json:"IsPushable"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
}

type DifficultyPool struct {
	Rolls      int          `json:"rolls"`
	EntityPool []EntityRoll `json:"pool"`
}

type EntityRoll struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type DynamicEntity struct {
	Name       string     `json:"key"`
	Parent     string     `json:"parent"`
	Sprite     string     `json:"sprite"`
	HP         int        `json:"hp"`
	Min        int        `json:"min"`
	Max        int        `json:"max"`
	Speed      float64    `json:"speed"`
	TxtBoxY    float64    `json:"txtBoxY"`
	TxtBoxX    float64    `json:"txtBoxX"`
	ParentText ParentText `json:"parentText"`
}

type ParentText struct {
	DropOffText []string `json:"dropOffText"`
	PickUpText  []string `json:"pickUpText"`
	SafeText    []string `json:"safeText"`
	HurtText    []string `json:"hurtText"`
	DeadText    []string `json:"deadText"`
}

func LoadEntityDefinitions(path string) (*EntityDefinitions, error) {
	errMsg := "Load entity definitions failed"
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var entityDefintions EntityDefinitions
	err = decoder.Decode(&entityDefintions)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	return &entityDefintions, nil
}

func PickRandomDynamicEntity(PoolType string) *DynamicEntity {
	roll := LoadedEntities.DynamicEntities[LoadedEntities.ExpandedEntityPools[PoolType][GlobalSeededRandom.Intn(LoadedEntities.ExpandedEntityTotals[PoolType])]]
	return roll
}
