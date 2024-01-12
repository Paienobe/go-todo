package main

import (
	"net/http"

	"github.com/Paienobe/go-todo/utils"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	type home struct {
		Msg string `json:"msg"`
	}
	utils.RespondWithJSON(w, 200, home{Msg: "A golang todo app"})
}
