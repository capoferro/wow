package wow

type Pet struct {
	Name string
	SpellId int
	CreatureId int
	Quality int
	Icon string
	Stats *PetStats
	BattlePetGuid string
	IsFavorite bool
	isFirstAbilitySlotSelected bool
	isSecondAbilitySlotSelected bool
	isThirdAbilitySlotSelected bool
	creatureName string
	canBattle bool
}
