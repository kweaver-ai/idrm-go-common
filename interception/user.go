package interception

import (
	"context"
	"github.com/gin-gonic/gin"
)

const contextKeyUser = "GoCommon/interception.User"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewContextWithUser 生成一个包含 User 的 context.Context
func NewContextWithUser(parent context.Context, u *User) context.Context {
	return context.WithValue(parent, contextKeyUser, u)
}

func SetGinContextWithUser(c *gin.Context, u *User) {
	c.Set(contextKeyUser, u)
}

// UserFromGinContext 从 Context 获取 User。User 可能是应用或用户。
func UserFromGinContext(c *gin.Context) (*User, error) {
	v, exists := c.Get(contextKeyUser)
	if !exists {
		return nil, ErrNotExist
	}

	u, ok := v.(*User)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return u, nil
}

// UserIDFromGinContext 从 Context 获取 UserID。UserID 可能是应用 ID 或用户 ID。
func UserIDFromGinContext(c *gin.Context) (string, error) {
	v, exists := c.Get(contextKeyUser)
	if !exists {
		return "", ErrNotExist
	}

	u, ok := v.(*User)
	if !ok {
		return "", ErrUnexpectedType
	}

	return u.ID, nil
}
