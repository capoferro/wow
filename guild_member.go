package wow

type GuildMember struct {
	Character *SimpleCharacter
	Rank      int
}

// Cast []*GuildMember to GuildMemberByRank to use stdlib sort. Will sort by Rank.
type ByRank []*GuildMember

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return a[i].Rank < a[j].Rank }
