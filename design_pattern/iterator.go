package design_pattern

type Iterator interface {
	HasNext() bool
	GetNext() *User
}

type Collection interface {
	CreateIterator() Iterator
}

type User struct {
	Name string
	Age  int
}

type UserCollection struct {
	Users []*User
}

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	return len(u.users)-u.index > 0
}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index += 1
		return user
	}
	return nil
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.Users,
	}
}
