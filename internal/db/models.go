package db // Same package name as the one used in db.go

// Card represents a Kanban card in the database.
type Card struct {
	ID   int      `json:"id"`
	Text string   `json:"text"`
	Tags []string `json:"tags,omitempty"`
}
