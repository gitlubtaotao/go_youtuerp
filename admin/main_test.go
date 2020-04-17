package main

import (
	"encoding/json"
	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initTestApp(t *testing.T) *httpexpect.Expect {
	e := httptest.New(t, NewAppInfo())
	return e
}
func TestLoginNotData(t *testing.T) {
	e := initTestApp(t)
	content := e.POST("/user/login").Expect().Status(httptest.StatusBadRequest).Body()
	result := stringToMap(content.Raw())
	assert.Equal(t, float64(400), result["code"])
}
func TestPasswordIsNull(t *testing.T) {
	e := initTestApp(t)
	e.POST("/user/login").WithJSON(map[string]interface{}{
		"username": "admin@example.com",
	}).Expect().Status(httptest.StatusUnauthorized)
}
func TestLogin(t *testing.T) {
	e := initTestApp(t)
	e.POST("/user/login").WithJSON(map[string]interface{}{
		"username": "admin@example.com",
		"password": "password",
	}).Expect().Status(httptest.StatusOK).Body().Raw()
}



func stringToMap(content string) map[string]interface{} {
	b := []byte(content)
	m := make(map[string]interface{})
	_ = json.Unmarshal(b, &m)
	return m
}
