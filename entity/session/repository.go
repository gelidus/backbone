package session

// Repository is the base interface for the session
type Repository interface {
	Create(session *Model) (error)
	ReadByLink(id string) (*Model, error)
}
