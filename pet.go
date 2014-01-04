package wow

type Pet struct {
	Name                        string
	SpellId                     int
	CreatureId                  int
	Quality                     int
	Icon                        string
	Stats                       *PetStats
	BattlePetGuid               string
	IsFavorite                  bool
	IsFirstAbilitySlotSelected  bool
	IsSecondAbilitySlotSelected bool
	IsThirdAbilitySlotSelected  bool
	CreatureName                string
	CanBattle                   bool
}
