package user

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Password string        `bson:""`
	Note     string        `bson:"note"`
}

func (User) TableName() string {
	return "user"
}
