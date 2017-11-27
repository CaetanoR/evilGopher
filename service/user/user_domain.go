package user

type User struct {

	Name string
	Following []*User
	Followers []*User

}

func (u *User) Follow(user *User) {
	u.Following = append(u.Following, user)
	user.Followers = append(user.Followers, u)
}

func NewUser(name string) *User {
	return &User{name, nil, nil}
}
