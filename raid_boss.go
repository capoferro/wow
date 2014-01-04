package wow

type RaidBoss struct {
	Id              int
	Name            string
	NormalKills     int
	NormalTimestamp uint64
	HeroicKills     int
	HeroicTimestamp uint64
	LFRKills        int
	LFRTimestamp    uint64
	FlexKills       int
	FlexTimestamp   uint64
}
