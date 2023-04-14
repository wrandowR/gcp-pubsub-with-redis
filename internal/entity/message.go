package entity

// Message is the entity that represents a message
type Message struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Message string `json:"message"`
}
