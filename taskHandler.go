package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/utils"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createTask(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
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
		utils.ResponsWithError(w, 400, fmt.Sprintf("Failed to create task %v", err))
		return
	}

	utils.RespondWithJSON(w, 201, dbTaskToTask(task))

}
