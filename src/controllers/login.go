package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)

	databaseUser, erro := repository.FindByEmail(user.Email)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	erro = security.ValidatePassword(user.Password, databaseUser.Password)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CreateToken(databaseUser.ID)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, token)
}
