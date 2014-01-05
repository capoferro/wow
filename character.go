package wow

import (
	"errors"
	"fmt"
)

type Character struct {
	AchievementPoints   int
	Battlegroup         string
	ClassId             int `json:"class"`
	class               string
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
	ApiClient           *ApiClient
}

func NewCharacter(client *ApiClient) *Character {
	return &Character{ApiClient: client}
}

// Character#Class() retrieves the name of the class of the character via the 
func (c *Character) Class() (string, error) {
	if c.ApiClient == nil {
		return "", errors.New("Character instance does not have an ApiClient reference. Please set ApiClient before calling Class(). c.ApiClient = NewApiClient(\"US\", \"\")")
	}
	if c.ClassId == 0 {
		return "", errors.New("Character instance does not have a class id.")
	}

	classes, err := c.ApiClient.GetClasses()
	if err != nil {
		return "", err
	}

	for _, class := range classes {
		if c.ClassId == class.Id {
			c.class = class.Name
		}
	}
	if c.class == "" {
		return "", errors.New(fmt.Sprintf("%d is not a valid class id", c.ClassId))
	}

	return c.class, nil

}
