package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestbody, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestbody, &user); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("signup"); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)

	user.ID, erro = repository.Create(user)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)

	user, erro := repository.Find(userId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if (user == models.User{}) {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)

	users, erro := repository.FindAll(nameOrNick)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if len(users) <= 0 {
		response.JSON(w, http.StatusOK, struct{}{})
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)

	user, erro := repository.Find(userId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if (user == models.User{}) {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	if erro = repository.Delete(userId); erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, erro := strconv.ParseUint(parameters["id"], 10, 64)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	tokenUserId, erro := authentication.GetUserIdFromRequest(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if tokenUserId != userId {
		fmt.Printf("%d %d", tokenUserId, userId)
		response.Error(w, http.StatusForbidden, errors.New("you cant update this user"))
		return
	}

	if erro = user.Prepare("update"); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)

	if erro = repository.Update(userId, user); erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
