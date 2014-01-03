package wow

type ItemSet struct {
	Id int
	Items []int
	Name string
	SetBonuses []*SetBonus
}
