package main

import (
	"encoding/json"
	"fmt"
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

func (apiCfg *apiConfig) login(w http.ResponseWriter, r *http.Request) {
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

	user, err := apiCfg.DB.GetUserByNameAndEmail(r.Context(), database.GetUserByNameAndEmailParams{
		Name:  params.Name,
		Email: params.Email,
	})

	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Failed to find user %v", err))
		return
	}

	tokenString, err := utils.GenerateJWT(user.Name, user.Email)
	if err != nil {
		utils.ResponsWithError(w, 500, fmt.Sprintf("failed to generate JWT %v", err))
	}

	tasks, err := apiCfg.DB.GetUserTasks(r.Context(), user.ID)
	if err != nil {
		utils.ResponsWithError(w, 500, fmt.Sprintf("could'nt find tasks for this user %v", err))
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "go_todo_jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 10),
	})

	type Success struct {
		Tasks []Task `json:"tasks"`
	}

	utils.RespondWithJSON(w, 200, Success{Tasks: dbTasksToTasks(tasks)})

}
