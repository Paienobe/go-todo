package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/types"
	"github.com/Paienobe/go-todo/utils"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createTask(w http.ResponseWriter, r *http.Request, user database.User) {

	params := types.CreateTaskParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Failed to parse body %v", err))
		return
	}

	task, err := apiCfg.DB.CreateTask(r.Context(), database.CreateTaskParams{
		ID:          uuid.New(),
		Name:        params.Name,
		Iscompleted: false,
		UserID:      user.ID,
	})

	if err != nil {
		fmt.Println("err from create Task", err)
		utils.ResponsWithError(w, 400, fmt.Sprintf("Failed to create task %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, dbTaskToTask(task))

}

func (apiCfg *apiConfig) updateTask(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		IsCompleted bool      `json:"isCompleted"`
		ID          uuid.UUID `json:"id"`
	}

	params := parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Error parsing body %v", err))
		return
	}

	task, err := apiCfg.DB.ToggleTaskStatus(r.Context(), database.ToggleTaskStatusParams{
		Iscompleted: params.IsCompleted, ID: params.ID,
	})
	fmt.Println()

	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Error updating task: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, dbTaskToTask(task))
}

func (apiCfg *apiConfig) deleteTask(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Error parsing body %v", err))
		return
	}

	err = apiCfg.DB.DeleteTask(r.Context(), params.ID)
	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Error deleting task %v", err))
		return
	}

	type success struct {
		Msg string `json:"msg"`
	}
	utils.RespondWithJSON(w, 200, success{Msg: "task deleted"})
}
