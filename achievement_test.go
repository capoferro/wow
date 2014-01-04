package wow

import (
    . "launchpad.net/gocheck"
)

type AchievementSuite struct{}
var _ = Suite(&AchievementSuite{})

func (s *AchievementSuite) Test_IsGroup(c *C) {
	a := &Achievement{Achievements: []*Achievement{&Achievement{}}}
	c.Assert(a.IsGroup(), Equals, true)
}

func (s *AchievementSuite) Test_IsGroup_2(c *C) {
	a := &Achievement{Categories: []*Achievement{&Achievement{}}}
	c.Assert(a.IsGroup(), Equals, true)
}

func (s *AchievementSuite) Test_IsGroup_3(c *C) {
	a := &Achievement{Achievements: []*Achievement{}}
	c.Assert(a.IsGroup(), Equals, false)
}
