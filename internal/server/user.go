package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofrumento/go-blog/internal/store"
)

func signUp(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := store.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Sign Up successfully",
		"jwt": "1234565789",
	})
}

func signIn(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	user, err := store.Authenticate(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign In failed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Sign In successfully",
		"jwt": "1234565789",
	})
}
