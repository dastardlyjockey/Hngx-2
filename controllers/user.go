package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dastardlyjockey/hngx-2/internal/database"
	"github.com/dastardlyjockey/hngx-2/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type parameters struct {
	Name string `json:"name"`
}

func (c *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to decode the error: %v", err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user, err := c.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed creating the user, %v", err.Error()))
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}

func (c *ApiConfig) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid userID, %v", err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user, err := c.DB.GetUserById(ctx, userID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("User not found in the database, %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

func (c *ApiConfig) UpdateUserNameById(w http.ResponseWriter, r *http.Request) {
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error decoding the JSON, %v", err))
		return
	}

	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error retrieving the User ID from the URL, %v", err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	updatedUser, err := c.DB.UpdateUserNameById(ctx, database.UpdateUserNameByIdParams{
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		ID:        userID,
	})
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update the user name, %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(updatedUser))
}

func (c *ApiConfig) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string `json:"message"`
	}

	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error retrieving the User ID from the URL, %v", err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exist, err := c.DB.UserExistsById(ctx, userId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("User not found in the database, %v", err))
		return
	}

	if !exist {
		RespondWithJSON(w, http.StatusNotFound, response{Message: "User not found"})
		return
	}

	err = c.DB.DeleteUserById(ctx, userId)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to delete the user, %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, response{Message: "User deleted successfully"})
}
