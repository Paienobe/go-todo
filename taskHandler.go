package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/types"
	"github.com/Paienobe/go-todo/utils"
	"github.com/go-chi/chi"
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

	taskId, err := uuid.Parse(chi.URLParam(r, "task_id"))
	if err != nil {
		log.Println("Failed to parse id", err)
	}

	err = apiCfg.DB.DeleteTask(r.Context(), taskId)
	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Error deleting task %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, types.Success{Msg: "task deleted"})
}

func (apiCfg *apiConfig) deleteAllUserTasks(w http.ResponseWriter, r *http.Request, user database.User) {
	err := apiCfg.DB.DeleteAllUserTasks(r.Context(), user.ID)
	if err != nil {
		utils.ResponsWithError(w, 400, fmt.Sprintf("Failed to delete tasks %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, types.Success{Msg: "All tasks deleted"})
}
