package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	email  string             `json:"email" bson:"email,omitempty"`
	FullName *FullName         `json:"fullname" bson:fullname,omitempty"`
}

type FullName struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
