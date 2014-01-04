package wow

type ItemClass struct {
	Class      int
	Name       string
	Subclasses []*ItemSubclass
}
