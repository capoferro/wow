package wow

import (
	"encoding/json"
)

type Item struct {
	Icon                   string
	Id                     int
	Name                   string
	Quality                int
	TooltipParams          *TooltipParams
	BonusStats             []*Stat
	Stats                  []*Stat
	Armor                  int
	BaseArmor              int
	WeaponInfo             *WeaponInfo
	BuyPrice               int
	ContainerSlots         int
	Description            string
	DisenchantingSkillRank int
	DisplayInfoId          int
	Equipable              bool
	HasSockets             bool
	HeroicTooltip          bool
	InventoryType          int
	IsAuctionable          bool
	ItemBind               int
	ItemClass              int
	ItemLevel              int
	ItemSource             *ItemSource
	ItemSpells             []*int
	ItemSubclass           int
	MaxCount               int
	MaxDurability          int
	MinFactionId           int
	MinReputation          int
	NameDescription        string
	NameDescriptionColor   string
	RequiredLevel          int
	RequiredSkill          int
	RequiredSkillRank      int
	SellPrice              int
	Stackable              int
	Upgradable             bool
}

func NewItemFromJson(jsonBlob []byte) (*Item, error) {
	item := &Item{}
	err := json.Unmarshal(jsonBlob, item)
	if err != nil {
		return nil, err
	}

	if len(item.Stats) == 0 {
		item.Stats = item.BonusStats
	}
	if len(item.BonusStats) == 0 {
		item.BonusStats = item.Stats
	}

	return item, nil
}
