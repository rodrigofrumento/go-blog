package store

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigofrumento/go-blog/internal/conf"
	"github.com/rodrigofrumento/go-blog/internal/store"
)

func testSetup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	store.ResetTestDatabase()
	cfg := conf.NewConfig("dev")
	jwtSetup(cfg)
	return setRouter(cfg)
}

func addTestUser() (*User, error) {
	user := &User{
		Username: "batman",
		Password: "secret123",
	}
	err := AddUser(user)
	return user, err
}
