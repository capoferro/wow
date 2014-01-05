package wow

import (
	. "launchpad.net/gocheck"
)

type CharacterSuite struct{}

var _ = Suite(&CharacterSuite{})

func (s *CharacterSuite) Test_CharacterClass_noApiClient(c *C) {
	ch := &Character{}
	_, err := ch.Class()
	c.Assert(err.Error(), Equals, "Character instance does not have an ApiClient reference. Please set ApiClient before calling Class(). c.ApiClient = NewApiClient(\"US\", \"\")")
}

func (s *CharacterSuite) Test_CharacterClass_noClassId(c *C) {
	client, _ := NewApiClient("US", "")
	ch := NewCharacter(client)
	_, err := ch.Class()
 c.Assert(err.Error(), Equals, "Character instance does not have a class id.")
}

func (s *CharacterSuite) Test_CharacterClass(c *C) {
	client, _ := NewApiClient("US", "")
	ch := &Character{ApiClient: client, ClassId: 6}
	class, _ := ch.Class()
	c.Assert(class, Equals, "Death Knight")
}
