package wow

type MountList struct {
	NumCollected    int
	NumNotCollected int
	Collected       []*Mount
}
