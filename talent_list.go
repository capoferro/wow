package wow

type TalentList struct {
	Selected bool
	Talents []*Talent
	Glyphs *GlyphList
	Spec *Spec
	CalcTalent string
	CalcSpec string
	CalcGlyph string
}
