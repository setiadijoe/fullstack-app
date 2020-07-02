package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/setiadijoe/fullstack/app/helpers"
	"github.com/setiadijoe/fullstack/app/internal/models"
	"github.com/setiadijoe/fullstack/app/responses"
)

// Login ...
func (svr *Server) Login(write http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := user.Login(svr.DB, user.Email, user.Password)
	if nil != err {
		errFormatted := helpers.FormatError(err.Error())
		responses.ERROR(write, http.StatusUnprocessableEntity, errFormatted)
		return
	}
	responses.JSON(write, http.StatusOK, token)
}
