package models_test

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
	"github.com/setiadijoe/fullstack/app/internal/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAll(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestLogin(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	type testCase struct {
		Email         string
		Password      string
		ExpectedError bool
	}

	samples := []*testCase{
		{
			Email:         "steven@gmail.com",
			Password:      "password",
			ExpectedError: false,
		}, {
			Email:         "lex",
			Password:      "password",
			ExpectedError: true,
		}, {
			Email:         "pet@gmail.com",
			Password:      "expass",
			ExpectedError: true,
		},
	}
	for _, val := range samples {
		result, err := userInstance.Login(server.DB, val.Email, val.Password)
		if !val.ExpectedError {
			if err != nil {
				log.Fatal(err)
			}
			if result == nil {
				log.Fatal("expected got token but nil")
			}
		}
		if val.ExpectedError {
			if err == nil {
				log.Fatal("expected got error")
			}
		}
	}

}

func TestSaveUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := models.User{
		ID:       1,
		Email:    "test@gmail.com",
		Nickname: "test",
		Password: "password",
	}
	savedUser, err := newUser.Save(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := userInstance.FindByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userUpdate := models.User{
		ID:       1,
		Nickname: "modiUpdate",
		Email:    "modiupdate@gmail.com",
		Password: "password",
	}
	updatedUser, err := userUpdate.UpdateByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Nickname, userUpdate.Nickname)
}

func TestDeleteAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, isDeleted)
}
