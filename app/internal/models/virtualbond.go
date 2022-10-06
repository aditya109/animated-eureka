package models

// VirtualBond defines the structure of a non-unique timed-virtual bond created between 2 sets of factions available uniquely identified by a UID
type VirtualBond struct {
	VirtualBondId string   `json:"virtual_bond_id"`
	UID           string   `json:"uid"`
	AFactionIds   []string `json:"a_faction_ids"`
	BFactionIds   []string `json:"b_faction_ids"`
	BondEndTime   string   `json:"bond_end_time"`
}
