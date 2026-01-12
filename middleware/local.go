package middleware

import (
	"context"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kweaver-ai/idrm-go-common/interception"
)

func LocalToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenID := c.GetHeader("Authorization")
		userInfo := &User{
			ID:   c.GetHeader("user_id"),
			Name: c.GetHeader("user_name"),
		}
		c.Set(interception.InfoName, userInfo)
		c.Set(interception.Token, tokenID)
		interception.SetGinContextWithAuth(c, tokenID)
		c.Set(interception.TokenType, interception.TokenTypeUser)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.InfoName, userInfo))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.Token, tokenID))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.TokenType, interception.TokenTypeUser))
		c.Request = c.Request.WithContext(interception.NewContextWithAuth(c.Request.Context(), tokenID))
		c.Next()
	}
}

func LocalAppToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenID := c.GetHeader("Authorization")
		token := strings.TrimPrefix(tokenID, "Bearer ")
		tokenType, _ := strconv.Atoi(c.GetHeader("token_type"))
		userInfo := &User{
			ID:       c.GetHeader("id"),
			Name:     c.GetHeader("name"),
			UserType: tokenType,
		}

		c.Set(interception.InfoName, userInfo)
		c.Set(interception.Token, token)
		c.Set(interception.TokenType, tokenType)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.InfoName, userInfo))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.Token, token))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.TokenType, tokenType))

		c.Next()
	}
}
