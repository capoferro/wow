package wow

import (
	"testing"
    . "launchpad.net/gocheck"
)

// GoCheck boilerplate
func Test(t *testing.T) { TestingT(t) }

type ApiClientSuite struct{}
var _ = Suite(&ApiClientSuite{})

func (s *ApiClientSuite) Test_NewApiClient_default(c *C) {
	client, _ := NewApiClient("US", "")
	c.Assert(client.Host, Equals, "us.battle.net")
	c.Assert(client.Locale, Equals, "en_US")
}

func (s *ApiClientSuite) Test_NewApiClient_specific(c *C) {
	client, _ := NewApiClient("EU", "fr_FR")
	c.Assert(client.Host, Equals, "eu.battle.net")
	c.Assert(client.Locale, Equals, "fr_FR")
}

func (s *ApiClientSuite) Test_NewApiClient_invalid(c *C) {
	_, err := NewApiClient("China", "it_IT")
	c.Assert(err.Error(), Equals, "Locale 'it_IT' is not valid for region 'China'")
}

func (s *ApiClientSuite) Test_NewApiClient_invalidRegion(c *C) {
	_, err := NewApiClient("Notaregion", "")
	c.Assert(err.Error(), Equals, "Region 'Notaregion' is not valid")
}

func (s *ApiClientSuite) Test_GetAchievement(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetAchievement(2144)

  c.Assert(a.AccountWide, Equals, true)
	c.Assert(a.Description, Equals, "Complete the world events achievements listed below.")
	c.Assert(len(a.Criteria), Equals, 9)
	c.Assert(len(a.RewardItems), Equals, 1)
}

func (s *ApiClientSuite) Test_GetAuctionData(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetAuctionData("Runetotem")
	c.Assert(len(a.Files), Equals, 1)
}

func (s *ApiClientSuite) Test_GetBattlePetAbility(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetBattlePetAbility(640)
	c.Assert(a.Id, Equals, 640)
	c.Assert(a.Cooldown, Equals, 0)
	c.Assert(a.Icon, Equals, "spell_shadow_plaguecloud")
	c.Assert(a.IsPassive, Equals, false)
	c.Assert(a.Name, Equals, "Toxic Smoke")
	c.Assert(a.PetTypeId, Equals, 9)
	c.Assert(a.Rounds, Equals, 1)
	c.Assert(a.ShowHints, Equals, false)
}

func (s *ApiClientSuite) Test_GetBattlePetSpecies(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetBattlePetSpecies(258)

	c.Assert(len(a.Abilities), Equals, 6)
	c.Assert(a.CanBattle, Equals, true)
	c.Assert(a.CreatureId, Equals, 42078)
	c.Assert(a.Description, Equals, "Powerful artillery of the terran army. The Thor is always the first one in and the last one out!")
	c.Assert(a.Icon, Equals, "t_roboticon")
	c.Assert(a.PetTypeId, Equals, 9)
	c.Assert(a.Source, Equals, "Promotion: StarCraft II: Wings of Liberty Collector's Edition")
	c.Assert(a.SpeciesId, Equals, 258)
}

func (s *ApiClientSuite) Test_GetBattlePetStats(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetBattlePet(258, 25, 5, 4)
	
  c.Assert(a.BreedId, Equals, 5)
	c.Assert(a.Health, Equals, 1587)
	c.Assert(a.Level, Equals, 25)
	c.Assert(a.PetQualityId, Equals, 4)
	c.Assert(a.Power, Equals, 315)
	c.Assert(a.SpeciesId, Equals, 258)
	c.Assert(a.Speed, Equals, 297)
}
