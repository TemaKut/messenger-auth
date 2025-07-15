package usermodels

type User struct {
	id       string
	name     string
	lastName string
}

func NewUserFromDb(id string, name string, lastName string) *User {
	return &User{
		id:       id,
		name:     name,
		lastName: lastName,
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) SetId(id string) {
	u.id = id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}
