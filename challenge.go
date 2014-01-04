package wow

type Challenge struct {
	Realm  *Realm
	Map    *Map
	Groups []*ChallengeGroup
}
