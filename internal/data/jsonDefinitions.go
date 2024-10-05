package data

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
