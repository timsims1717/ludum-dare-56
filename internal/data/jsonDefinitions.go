package data

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/timsims1717/pixel-go-utils/timing"
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
	Name            string  `json:"key"`
	Sprite          string  `json:"sprite"`
	Damage          int     `json:"damage"`
	Damagetype      string  `json:"damagetype"`
	IsCollidable    bool    `json:"IsCollidable"`
	IsPickupable    bool    `json:"IsPickupable"`
	IsPushable      bool    `json:"IsPushable"`
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	Uses            int     `json:"uses"`
	Zlevel          int     `json:"zlevel"`
	StatusEffect    string  `json:"statuseffect"`
	HasStatusEffect bool    `json:"hasStatusEffect"`
}

func (s StaticEntity) Clone() StaticEntity {
	return StaticEntity{
		Name:            s.Name,
		Sprite:          s.Sprite,
		Damagetype:      s.Damagetype,
		IsCollidable:    s.IsCollidable,
		IsPickupable:    s.IsPickupable,
		IsPushable:      s.IsPushable,
		Width:           s.Width,
		Height:          s.Height,
		Uses:            s.Uses,
		Damage:          s.Damage,
		Zlevel:          s.Zlevel,
		StatusEffect:    s.StatusEffect,
		HasStatusEffect: s.HasStatusEffect,
	}
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

func LoadStatusEffectDefinitions(path string) (*StatusDefinitions, error) {
	errMsg := "Load status effects definitions failed"
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var statusDefinitions StatusDefinitions
	err = decoder.Decode(&statusDefinitions)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	return &statusDefinitions, nil
}

type StatusDefinitions struct {
	StatusEffects map[string]*StatusEffect `json:"StatusDefinitions"`
}

type StatusEffect struct {
	Modifies           string  `json:"Modifies"`
	AugmentationType   string  `json:"AugmentationType"`
	AugmentationFactor float64 `json:"AugmentationFactor"`
	Duration           float64 `json:"Duration"`
	StatusTimer        *timing.Timer
	Name               string `json:"Name"`
}

func (s StatusEffect) Clone() StatusEffect {
	return StatusEffect{
		Modifies:           s.Modifies,
		AugmentationType:   s.AugmentationType,
		AugmentationFactor: s.AugmentationFactor,
		Duration:           s.Duration,
		StatusTimer:        timing.New(s.Duration),
		Name:               s.Name,
	}
}

func SpeedBoostApply(ch *Character, effect string) {
	ch.Speed = LoadedEntities.DynamicEntities[ch.EntityName].Speed * ch.StatusEffects[effect].AugmentationFactor
}

func SpeedBoostClear(ch *Character, effect string) {
	ch.Speed = LoadedEntities.DynamicEntities[ch.EntityName].Speed
}
