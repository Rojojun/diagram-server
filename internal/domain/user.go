package domain

type User struct {
	ID       string
	Email    string
	Password string
	Name     string
	Role     string
}

func (u *User) Validate() error {
	return nil
}
