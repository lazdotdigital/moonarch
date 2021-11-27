package moonarch

// Pair represents a token pair.
type Pair struct {
	Type         string  `json:"type"`
	Address      string  `json:"address"`
	WBNBIndex    int     `json:"wbnbIndex"`
	Supply       float64 `json:"supply"`
	ReserveWBNB  float64 `json:"reserveWBNB"`
	ReserveToken float64 `json:"reserveToken"`
	BurntSupply  float64 `json:"burntSupply"`
	LockedSupply float64 `json:"lockedSupply"`
	OwnerSupply  float64 `json:"ownerSupply"`
}
