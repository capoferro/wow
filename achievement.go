package wow

// Achievements are either individual achievements or achievement
// groups. Groups will contain achievements in Achievements and
// subgroups in Categories.
type Achievement struct {
	AccountWide bool
	Criteria []*AchievementCriteria
	Description string
	Icon string
	Id int
	Points int
	Reward string
	RewardItems []*Item
	Title string
	Achievements []*Achievement
	Categories []*Achievement
	Name string
}

func (a *Achievement) IsGroup() bool {
	return (len(a.Achievements) > 0 || len(a.Categories) > 0)
}
