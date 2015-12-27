package user

// Repository is any User repository that matches this
// interface pattern
type Repository interface {
	Create(user *Model) (error)
	ReadByID(id string) (*Model, error)
	ReadByEmail(email string) (*Model, error)
}
