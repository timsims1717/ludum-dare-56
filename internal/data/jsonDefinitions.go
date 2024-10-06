package data

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type EntityDefinitions struct {
	StaticEntities            map[string]*StaticEntity `json:"StaticEntities"`
	StaticEntityPool          []EntityRoll             `json:"StaticEntityPool"`
	StaticEntityPoolExpanded  []string
	StaticEntityPoolTotal     int
	DynamicEntities           map[string]*DynamicEntity `json:"DynamicEntities"`
	DynamicEntityPool         []EntityRoll              `json:"DynamicEntityPool"`
	DynamicEntityPoolExpanded []string
	DynamicEntityPoolTotal    int
}

type StaticEntity struct {
	Name         string `json:"key"`
	Sprite       string `json:"sprite"`
	Damage       int    `json:"damage"`
	Damagetype   string `json:"damagetype"`
	IsCollidable bool   `json:"IsCollidable"`
}

type EntityRoll struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type DynamicEntity struct {
	Name   string  `json:"key"`
	Parent string  `json:"parent"`
	Sprite string  `json:"sprite"`
	HP     int     `json:"hp"`
	Min    int     `json:"min"`
	Max    int     `json:"max"`
	Speed  float64 `json:"speed"`
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

func PickRandomDynamicEntity() *DynamicEntity {
	roll := LoadedEntities.DynamicEntities[LoadedEntities.DynamicEntityPoolExpanded[GlobalSeededRandom.Intn(LoadedEntities.DynamicEntityPoolTotal)]]
	return roll
}
