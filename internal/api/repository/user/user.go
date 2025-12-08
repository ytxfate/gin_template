package user

import "gin_template/internal/api/models"

func FindUserByUserPwd(username, password string) (*models.User, error) {
	// TODO: db query
	return &models.User{
		Username: username,
		Password: password,
		Note:     "xxx",
	}, nil
}
