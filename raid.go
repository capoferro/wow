package wow

type Raid struct {
	Name   string
	Normal int
	Heroic int
	Id     int
	Bosses []*RaidBoss
}
