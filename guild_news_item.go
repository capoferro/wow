package wow

type GuildNewsItem struct {
	Type        string
	Character   string
	Timestamp   uint64
	ItemId      int
	Achievement *Achievement
}
