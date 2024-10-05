package data

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type EntityDefintions struct {
	StaticEntities          map[string]StaticEntity `json:"StaticEntities"`
	StaticEntityPool        []EntityRoll            `json:"StaticEntityPool"`
	StaticEnityPoolExpanded []string
	StaticEntityPoolTotal   int
}

type StaticEntity struct {
	Name   string `json:"key"`
	Sprite string `json:"sprite"`
}

type EntityRoll struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
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
