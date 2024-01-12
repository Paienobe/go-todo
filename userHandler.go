package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/utils"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponsWithError(w, 500, "failed to decode post body")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(), Name: params.Name, Email: params.Email, CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		utils.ResponsWithError(w, 500, "Failed to create user")
	}

	utils.RespondWithJSON(w, 201, dbUserToUser(user))
}

func (apiCfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, 200, dbUserToUser(user))
}
