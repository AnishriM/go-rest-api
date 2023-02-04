package http

// Handler - stores pointer to our comment service
type Handler struct{}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}
