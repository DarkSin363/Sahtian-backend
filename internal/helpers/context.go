package helpers

import (
	"context"
	"github.com/BigDwarf/sahtian/internal/model"
)

const userKey = "userKey"

func WithUser(c context.Context, user *model.User) context.Context {
	return context.WithValue(c, userKey, user)
}

func GetUser(c context.Context) *model.User {
	return c.Value(userKey).(*model.User)
}
