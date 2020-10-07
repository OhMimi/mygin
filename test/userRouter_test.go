package test

import (
	"bytes"
	"encoding/json"
	"mygin/router"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGetRouter(t *testing.T) {
	type User struct {
		UserID string `json:"userId"`
	}

	user := User{
		UserID: "123",
	}

	expectedBody, _ := json.Marshal(user)
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+user.UserID, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedBody), w.Body.String())

}

func TestUserLoginRouter(t *testing.T) {
	value := url.Values{}
	value.Add("email", "test@example.com")
	value.Add("password", "qwe123")
	value.Add("password-again", "qwe123")

	router := router.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
