package wow

type PetList struct {
	NumCollected    int
	NumNotCollected int
	Collected       []*Pet
}
