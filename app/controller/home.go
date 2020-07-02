package controller

import (
	"net/http"

	"github.com/setiadijoe/fullstack/app/responses"
)

// HomePage ...
func (svr *Server) HomePage(write http.ResponseWriter, req http.Request) {
	responses.JSON(write, http.StatusOK, "welcome_budum")
}
