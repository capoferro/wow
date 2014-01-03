package wow

type PetSlot struct {
	Slot int
	BattlePetGuid string
	IsEmpty bool
	IsLocked bool
	Abilities []int
}
