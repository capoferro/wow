package wow

type Guild struct {
	Name              string
	Realm             string
	Battlegroup       string
	LastModified      uint64
	Level             int
	Members           []*GuildMember
	AchievementPoints int
	Emblem            *GuildEmblem
	Side              int
	Achievements      *AchievementList
	News              []*GuildNewsItem
	Challenge         []*Challenge
}
