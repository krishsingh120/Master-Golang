package user

type UserModel struct {
	Email    string
	Username string
	Mobile   int
}

func (u *UserModel) UpdateDetails(email *string, username *string, mobile *int) {
	if email != nil {
		u.Email = *email
	}
	if username != nil {
		u.Username = *username
	}
	if mobile != nil {
		u.Mobile = *mobile
	}
}
