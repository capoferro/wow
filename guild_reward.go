package wow

type GuildReward struct {
	MinGuildLevel int
	MinGuildRepLevel int
	Races []int
	Achievement *Achievement
	Item *Item
}
