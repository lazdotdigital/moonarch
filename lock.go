package moonarch

// Lock represents an LP lock.
type Lock struct {
	Type        string  `json:"type"`
	Locker      string  `json:"locker"`
	LockID      int     `json:"lockId"`
	Hash        string  `json:"hash"`
	Start       int     `json:"start"`
	End         int     `json:"end"`
	Amount      float64 `json:"amount"`
	PairAddress string  `json:"pairAddress"`
	Address     string  `json:"address"`
	Network     string  `json:"network"`
}
