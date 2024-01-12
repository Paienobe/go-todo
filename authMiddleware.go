package main

import (
	"fmt"
	"net/http"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/utils"
)

type authenticatedUserHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) authMiddleware(handler authenticatedUserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := utils.GetApiKey(r)
		if err != nil {
			utils.ResponsWithError(w, 400, fmt.Sprintf("Auth err: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.ResponsWithError(w, 500, fmt.Sprintf("Error fetching user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
