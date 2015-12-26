package session

// Repository is the base interface for the session
type Repository interface {
	Create(session *Session) (error)
	ReadByLink(id string) (*Session, error)
}
