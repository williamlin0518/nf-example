package sbi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	server := &Server{}
	server.GetHello(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello from feature/hello!", w.Body.String())
}

func TestPostHello(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Params = []gin.Param{{Key: "message", Value: "test"}}
	
	server := &Server{}
	server.PostHello(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Received message: test", w.Body.String())
}