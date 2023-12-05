package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userTokenId, erro := authentication.GetUserIdFromRequest(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro = json.Unmarshal(requestBody, &post); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	post.CreatorId = userTokenId

	if erro = post.Prepare(); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	repository := repositories.PostRepository(db)
	post.ID, erro = repository.Create(post)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}

func FindPosts(w http.ResponseWriter, r *http.Request) {
	userTokenId, erro := authentication.GetUserIdFromRequest(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.PostRepository(db)

	posts, erro := repository.Find(userTokenId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func FindPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PostRepository(db)

	post, erro := repository.FindById(postId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if (post == models.Post{}) {
		response.Error(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	response.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userTokenId, erro := authentication.GetUserIdFromRequest(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PostRepository(db)
	databasePost, erro := repository.FindById(postId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if (databasePost == models.Post{}) {
		response.Error(w, http.StatusNotFound, errors.New("this post does not exists"))
		return
	}

	if databasePost.CreatorId != userTokenId {
		response.Error(w, http.StatusForbidden, errors.New("you cant update other posts"))
		return
	}

	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro = json.Unmarshal(requestBody, &post); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Prepare(); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.Update(postId, post); erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userTokenId, erro := authentication.GetUserIdFromRequest(r)
	if erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PostRepository(db)
	databasePost, erro := repository.FindById(postId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if (databasePost == models.Post{}) {
		response.Error(w, http.StatusNotFound, errors.New("this post does not exists"))
		return
	}

	if databasePost.CreatorId != userTokenId {
		response.Error(w, http.StatusForbidden, errors.New("you cant remove other posts"))
		return
	}

	if erro = repository.Delete(postId); erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindUserPosts(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.PostRepository(db)
	posts, erro := repository.FindByUser(userId)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PostRepository(db)
	if erro = repository.Like(postId); erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.PostRepository(db)
	if erro = repository.Unlike(postId); erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
