package user

type User interface {
	FullName() string
	Username() string
}

type user struct {
	fullName string
	username string
}

func (u user) FullName() string {
	return u.fullName
}

func (u user) Username() string {
	return u.username
}

func New(fullName, username string) User {
	return &user{
		fullName,
		username,
	}
}
