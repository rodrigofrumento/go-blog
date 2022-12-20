package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofrumento/go-blog/internal/conf"
	"github.com/rodrigofrumento/go-blog/internal/store"
	"github.com/rs/zerolog/log"
)

func testSetup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	store.ResetTestDatabase()
	jwtSetup(conf.Config{})
	return setRouter()
}

func userJSON(user store.User) string {
	body, err := json.Marshal(map[string]interface{}{
		"Username": user.Username,
		"Password": user.Password,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Error marshalling JSON body.")
	}
	return string(body)
}

func jsonRes(body *bytes.Buffer) map[string]interface{} {
	jsonValue := &map[string]interface{}{}
	err := json.Unmarshal(body.Bytes(), jsonValue)
	if err != nil {
		log.Panic().Err(err).Msg("Error unmarshalling JSON body.")
	}
	return *jsonValue
}

func performRequest(router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		log.Panic().Err(err).Msg("Error creating new request")
	}
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(rec, req)
	return rec
}
