package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const ok_hello = `{"message":"` + OK_HELLO + `"}`

func TestHello(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller := &HelloController{}

	if assert.NoError(t, controller.Hello(c)) {
		newl := len(rec.Body.String()) - 1

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ok_hello, rec.Body.String()[:newl])
	}

}
