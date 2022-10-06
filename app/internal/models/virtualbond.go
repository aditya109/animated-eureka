package models

// VirtualBond defines the structure of a non-unique timed-virtual bond created between 2 sets of factions available uniquely identified by a UID
type VirtualBond struct {
	VirtualBondId string
	UID           string
	AFactionIds   []string
	BFactionIds   []string
	BondEndTime   string
}
