package wow

type RealmStatus struct {
	Type        string
	Population  string
	Queue       bool
	Wintergrasp *PvPZoneStatus
	TolBarad    *PvPZoneStatus `json:"tol-barad"`
	Status      bool
	Name        string
	Slug        string
	Battlegroup string
	Locale      string
	Timezone    string
}
