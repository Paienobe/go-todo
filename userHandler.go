package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/types"
	"github.com/Paienobe/go-todo/utils"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {

	params := types.RegistrationParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponsWithError(w, 500, "failed to decode post body")
		return
	}

	hashedPassword, err := utils.GeneratehashPassword(params.Password)
	if err != nil {
		log.Println("Failed to hash password", err)
		return
	}

	err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Email:     params.Email,
		CreatedAt: time.Now().UTC(),
		Password:  hashedPassword,
	})

	if err != nil {
		utils.ResponsWithError(w, 500, "Failed to create user")
	}

	utils.RespondWithJSON(w, 201, types.RegistrationSuccess{
		Success: true, Message: "User created.",
	})
}

func (apiCfg *apiConfig) GetUserTasks(w http.ResponseWriter, r *http.Request, user database.User) {
	tasks, err := apiCfg.DB.GetUserTasks(r.Context(), user.ID)
	if err != nil {
		utils.ResponsWithError(w, 400, "Error fethcing tasks")
		return
	}

	utils.RespondWithJSON(w, 200, dbTasksToTasks(tasks))
}

func (apiCfg *apiConfig) login(w http.ResponseWriter, r *http.Request) {

	params := types.LoginParams{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponsWithError(w, 500, "failed to decode post body")
		return
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)

	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Failed to find user %v", err))
		return
	}

	passwordIsCorrect := utils.CheckPasswordHash(params.Password, user.Password)
	if !passwordIsCorrect {
		utils.ResponsWithError(w, 401, "Password is incorrect")
		return
	}

	tokenString, err := utils.GenerateJWT(user.Email, user.ID)
	if err != nil {
		utils.ResponsWithError(w, 500, fmt.Sprintf("failed to generate JWT %v", err))
		return
	}

	tasks, err := apiCfg.DB.GetUserTasks(r.Context(), user.ID)
	if err != nil {
		utils.ResponsWithError(w, 500, fmt.Sprintf("could'nt find tasks for this user %v", err))
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt_todo",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 10),
	})

	utils.RespondWithJSON(w, 200, types.LoginSuccess[Task]{
		Success: true,
		Tasks:   dbTasksToTasks(tasks),
		Token:   tokenString,
	})

}
