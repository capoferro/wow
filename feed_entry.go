package wow

type FeedEntry struct {
	Type string
	Timestamp uint64
	Achievement *Achievement
	FeatOfStrength bool
	Criteria *AchievementCriteria
	Quantity int
	Name string
	ItemId int
}
