package domain

type User struct {
	id       string
	email    string
	password string
	name     string
}

func NewUser(email, password, name string) User {
	if len(password) < 8 {
	}

	return User{email: email, password: password, name: name}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetID(id string) {
	u.id = id
}

func (u *User) SetEmail(encrypt string) {
	u.password = encrypt
}
