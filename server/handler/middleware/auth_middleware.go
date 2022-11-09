package middleware

import (
	"net/http"
	"tgnotify/server/getter"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		m := getter.MustGetUserList(ctx)
		user := ctx.GetHeader("user")
		code := ctx.GetHeader("code")
		realcode, ok := m[user]
		if !ok {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		if realcode != code {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
