package wow

type SimpleGuild struct {
	Name string
	Realm string
	Battlegroup string
	Level int
	Members int
	AchievementPoints int
	Emblem *GuildEmblem
}
