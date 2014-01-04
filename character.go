package wow

type Character struct {
	AchievementPoints   int
	Battlegroup         string
	Class               int
	CalcClass           string
	Gender              int
	Level               int
	Name                string
	Race                int
	Realm               string
	Thumbnail           string
	LastModified        uint
	Guild               *SimpleGuild
	Feed                []*FeedEntry
	Items               *ItemList
	Stats               *CharacterStats
	Professions         *ProfessionList
	Reputation          []*Reputation
	Titles              []*Title
	Achievements        *AchievementList
	Talents             []*CharacterTalentList
	Appearance          *CharacterAppearance
	Mounts              *MountList
	Pets                *PetList
	PetSlots            []*PetSlot
	Progression         *ProgressionList
	PvP                 *PvPList
	Quests              []int
	TotalHonorableKills int
}
