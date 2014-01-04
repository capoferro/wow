package wow

type BattlePetSpecies struct {
	Abilities   []*BattlePetAbility
	CanBattle   bool
	CreatureId  int
	Description string
	Icon        string
	PetTypeId   int
	Source      string
	SpeciesId   int
}
