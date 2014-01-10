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

func (g *Guild) ItemNews() []*GuildNewsItem{
	itemNews := make([]*GuildNewsItem, 0)
	for _, n := range g.News {
		if n.Type == "itemLoot" {
			itemNews = append(itemNews, n)
		}
	}
	return itemNews
}
