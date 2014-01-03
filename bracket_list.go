package wow

type BracketList struct {
	ArenaBracket2v2 *ArenaBracket `json:"ARENA_BRACKET_2v2"`
	ArenaBracket3v3 *ArenaBracket `json:"ARENA_BRACKET_3v3"`
	ArenaBracket5v5 *ArenaBracket `json:"ARENA_BRACKET_5v5"`
	ArenaBracketRBG *ArenaBracket `json:"ARENA_BRACKET_RBG"`
}
