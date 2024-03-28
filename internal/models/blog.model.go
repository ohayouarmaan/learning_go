package models

import "gopkg.in/mgo.v2/bson"

type Blog struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Content     string        `json:"content" bson:"content"`
}
