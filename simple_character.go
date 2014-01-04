package wow

// WoW API does not provide the same data for a character when in a
// challenge object
type SimpleCharacter struct {
	Name string
	Realm string
	Battlegroup string
	Class int
	Race int
	Gender int
	Level int
	AchievementPoints int
	Thumbnail string
	Spec *Spec
	Guild string
	GuildRealm string
}
