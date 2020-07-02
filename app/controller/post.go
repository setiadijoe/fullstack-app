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

// CreatePost ...
func (svr *Server) CreatePost(write http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}
	userID, err := auth.ExtractTokenID(req)
	if nil != err {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if userID != post.AuthorID {
		responses.ERROR(write, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	postCreated, err := post.Save(svr.DB)
	if nil != err {
		errFormated := helpers.FormatError(err.Error())
		responses.ERROR(write, http.StatusInternalServerError, errFormated)
		return
	}
	write.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", req.Host, req.URL.Path, postCreated.ID))
	responses.JSON(write, http.StatusCreated, postCreated)
}

// GetPosts ...
func (svr *Server) GetPosts(write http.ResponseWriter, req *http.Request) {
	post := models.Post{}
	posts, err := post.FindAll(svr.DB)
	if nil != err {
		responses.ERROR(write, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(write, http.StatusOK, posts)
}

// GetPost ...
func (svr *Server) GetPost(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	postID, err := strconv.ParseUint(vars["id"], 10, 64)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}
	post := models.Post{}

	postReceived, err := post.FindByID(svr.DB, postID)
	if nil != err {
		responses.ERROR(write, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(write, http.StatusOK, postReceived)
}

// UpdatePost ...
func (svr *Server) UpdatePost(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// Check if the post id is valid
	postID, err := strconv.ParseUint(vars["id"], 10, 64)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}

	//Check if the auth token is valid and  get the user id from it
	userID, err := auth.ExtractTokenID(req)
	if nil != err {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	post := models.Post{}
	postReceived, err := post.FindByID(svr.DB, postID)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	if nil == postReceived {
		responses.ERROR(write, http.StatusNotFound, errors.New("post_not_found"))
		return
	}

	// If a user attempt to update a post not belonging to him
	if userID != post.AuthorID {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data posted
	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	postUpdate := models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if userID != postUpdate.AuthorID {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	postUpdate.ID = post.ID //this is important to tell the model the post id to update, the other update field are set above

	postUpdated, err := postUpdate.UpdateByID(svr.DB, postID, postUpdate.Title, postUpdate.Content)

	if nil != err {
		errFormatted := helpers.FormatError(err.Error())
		responses.ERROR(write, http.StatusInternalServerError, errFormatted)
		return
	}
	responses.JSON(write, http.StatusOK, postUpdated)
}

// DeletePost ...
func (svr *Server) DeletePost(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	// Is a valid post id given to us?
	postID, err := strconv.ParseUint(vars["id"], 10, 64)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	userID, err := auth.ExtractTokenID(req)
	if nil != err {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	post := models.Post{}
	postReceived, err := post.FindByID(svr.DB, postID)
	if nil != err {
		responses.ERROR(write, http.StatusUnprocessableEntity, err)
		return
	}

	if nil == postReceived {
		responses.ERROR(write, http.StatusNotFound, errors.New("post_not_found"))
		return
	}

	// Is the authenticated user, the owner of this post?
	if userID != post.AuthorID {
		responses.ERROR(write, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = post.DeleteByID(svr.DB, postID, userID)
	if nil != err {
		responses.ERROR(write, http.StatusBadRequest, err)
		return
	}
	write.Header().Set("Entity", fmt.Sprintf("%d", postID))
	responses.JSON(write, http.StatusNoContent, nil)
}
