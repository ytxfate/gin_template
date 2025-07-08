package userdao

import "go.mongodb.org/mongo-driver/v2/bson"

const UserColl = "user"

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	password string        `bson:"password"`
	Note     string        `bson:"note"`
}

func FindUserByUserPwd(username, password string) (*User, error) {
	// TODO: db query
	return &User{
		Username: username,
		password: password,
		Note:     "xxx",
	}, nil
}
