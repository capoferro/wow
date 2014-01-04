package wow

type CharacterTalentList struct {
	Selected bool
	Talents []*Talent
	Glyphs *GlyphList
	Spec *Spec
	CalcTalent string
	CalcSpec string
	CalcGlyph string
}
