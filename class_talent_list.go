package wow

type ClassTalentList struct {
	Warrior				*TalentList `json:"1"`
	Paladin				*TalentList `json:"2"`
	Hunter				*TalentList `json:"3"`
	Rogue					*TalentList `json:"4"`
	Priest				*TalentList `json:"5"`
	Deathknight		*TalentList `json:"6"`
	Shaman				*TalentList `json:"7"`
	Mage					*TalentList `json:"8"`
	Warlock				*TalentList `json:"9"`
	Monk					*TalentList `json:"10"`
	Druid					*TalentList `json:"11"`
}
