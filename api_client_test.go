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

