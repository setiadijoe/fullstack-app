package controller

import (
	"net/http"

	"github.com/setiadijoe/fullstack/app/responses"
)

// Home ...
func (svr *Server) Home(write http.ResponseWriter, req *http.Request) {
	host := req.URL.Host
	responses.JSON(write, http.StatusOK, host)
}
