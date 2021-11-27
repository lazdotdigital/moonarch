package moonarch

// Check represents an entry from moonarch's rugchecker.
type Check struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Level   int    `json:"level"`
}
