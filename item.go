package wow

type Item struct {
	Icon string
	Id int
	Name string
	Quality int
	ItemLevel int
	TooltipParams *TooltipParams
	Stats []*Stat
	Armor int
	WeaponInfo *WeaponInfo
}
