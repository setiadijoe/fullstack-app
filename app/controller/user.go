package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/setiadijoe/fullstack/app/auth"
	"github.com/setiadijoe/fullstack/app/helpers"
	"github.com/setiadijoe/fullstack/app/internal/models"
	"github.com/setiadijoe/fullstack/app/responses"

	"github.com/gorilla/mux"
)

// CreateUser ...
func (svr *Server) CreateUser(write http.ResponseWriter, req *http.Request) {
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
	err = user.Validate("")
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.Save(svr.DB)
	if nil != err {
		errFormated := helpers.FormatError(err.Error())
		responses.ERROR(write, http.StatusInternalServerError, errFormated)
		return
	}

	write.Header().Set("Location", fmt.Sprintf("%s%s/%d", req.Host, req.RequestURI, userCreated.ID))
	responses.JSON(write, http.StatusCreated, userCreated)
}

// GetUsers ...
func (svr *Server) GetUsers(write http.ResponseWriter, req *http.Request) {
	user := models.User{}

	users, err := user.FindAll(svr.DB)
	if nil != err {
		responses.ERROR(write, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(write, http.StatusOK, users)
}

// GetUser ...
func (svr *Server) GetUser(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindByID(svr.DB, uint32(id))
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}
	responses.JSON(write, http.StatusOK, userGotten)
}

// UpdateUser ...
func (svr *Server) UpdateUser(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}
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
	tokenID, err := auth.ExtractTokenID(req)
	if nil != err {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(id) {
		responses.ERROR(write, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateByID(svr.DB, uint32(id), user.Email, user.Nickname, user.Password)
	if nil != err {
		errFormated := helpers.FormatError(err.Error())
		responses.ERROR(write, http.StatusInternalServerError, errFormated)
		return
	}
	responses.JSON(write, http.StatusOK, updatedUser)
}

// DeleteUser ...
func (svr *Server) DeleteUser(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user := models.User{}

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(req)
	if nil != err {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(id) {
		responses.ERROR(write, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteByID(svr.DB, uint32(id))
	if nil != err {
		responses.ERROR(write, http.StatusInternalServerError, err)
		return
	}
	write.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(write, http.StatusNoContent, nil)
}
