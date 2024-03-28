package blog

import (
	"context"

	"github.com/ohayouarmaan/learning_go/internal/config"
)

func GetAllBlogs() map[string]string {
	db := config.CreateDBInstance("monogodb://localhost:27017/blocks")
	cur := (db.C("blogs")).Find(context.TODO())
	if cur != nil {
		panic("Cursor is nil")
	}
	return make(map[string]string)
}
