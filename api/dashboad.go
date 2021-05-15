package api

import (
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
