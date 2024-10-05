package data

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type EntityDefintions struct {
	StaticEntities   []StaticEntity `json:"StaticEntities"`
	StaticEntityPool []EntityRolls  `json:"StaticEntityPool"`
}

type StaticEntity struct {
	Name   string `json:"name"`
	Sprite string `json:"sprite"`
}

type EntityRolls struct {
	Name   string `json:"name"`
	Weight string `json:"weight"`
}

func LoadEntityDefinitions(path string) (*EntityDefintions, error) {
	errMsg := "Load entity definitions failed"
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var entityDefintions EntityDefintions
	err = decoder.Decode(&entityDefintions)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	return &entityDefintions, nil
}
