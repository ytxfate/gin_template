package user

func FindUserByUserPwd(username, password string) (*User, error) {
	// TODO: db query
	return &User{
		Username: username,
		Password: password,
		Note:     "xxx",
	}, nil
}
