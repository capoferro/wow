package wow

import (
	"time"
)

type ChallengeGroup struct {
	Date *time.Time
	Faction string
	IsRecurring bool
	Medal string
	Members []*ChallengeMember
	Ranking int
	Time *ChallengeTime
}

