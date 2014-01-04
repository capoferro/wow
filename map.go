package wow

type Map struct {
	Id               int
	Name             string
	Slug             string
	HasChallengeMode bool
	Bronze           *ChallengeTime `json:"bronzeCriteria"`
	Silver           *ChallengeTime `json:"silverCriteria"`
	Gold             *ChallengeTime `json:"goldCriteria"`
}
