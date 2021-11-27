package moonarch

// Token represents a BSC token.
type Token struct {
	Network      string  `json:"network"`
	Address      string  `json:"address"`
	AddressCheck string  `json:"addressCheck"`
	Checks       []Check `json:"checks"`
	Pairs        []Pair  `json:"pairs"`
	Verified     bool    `json:"verified"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Decimals     int     `json:"decimals"`
	Creator      string  `json:"creator"`
	CreationDate int     `json:"creationDate"`
	Supply       int     `json:"supply"`
	Owner        string  `json:"owner"`
	BurntSupply  float64 `json:"burntSupply"`
	LockedSupply float64 `json:"lockedSupply"`
	OwnerSupply  float64 `json:"ownerSupply"`
	PoolsSupply  float64 `json:"poolsSupply"`
	Updated      int     `json:"updated"`
}
