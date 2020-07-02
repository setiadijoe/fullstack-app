package controller

import (
	"net/http"

	"github.com/setiadijoe/fullstack/app/responses"
)

// Home ...
func (svr *Server) Home(write http.ResponseWriter, req *http.Request) {
	responses.JSON(write, http.StatusOK, "welcome_budum")
}
