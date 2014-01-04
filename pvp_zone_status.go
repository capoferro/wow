package wow

type PvPZoneStatus struct {
	Area               int
	ControllingFaction int `json:"controlling-faction"`
	Status             int
	Next               uint64
}
