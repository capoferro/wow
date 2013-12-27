package wow

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
}

