package user

// Repository is any User repository that matches this
// interface pattern
type Repository interface {
	Create(user *User) (error)
	ReadByID(id string) (*User, error)
	ReadByEmail(email string) (*User, error)
}
