package tests

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewGinContext() *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header = make(http.Header)
	c.Request = req
	return c
}
