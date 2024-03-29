package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/go-todo/internal/database"
	"github.com/Paienobe/go-todo/utils"
	"github.com/golang-jwt/jwt/v5"
)

type authenticatedUserHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) isAuthorized(handler authenticatedUserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			log.Fatal("JWT_SECRET does not exist in environment")
		}

		tokenString, err := utils.GetBearerToken(r)
		if err != nil {
			log.Fatal(err)
		}

		var mySigningKey = []byte(secretKey)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			utils.ResponsWithError(w, 500, fmt.Sprintf("Failed to parse JWT %v", err))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)

			user, err := apiCfg.DB.GetUserByEmail(r.Context(), email)
			if err != nil {
				utils.ResponsWithError(w, 400, fmt.Sprintf("User not found %v", err))
				return
			}

			handler(w, r, user)
			return
		}

		utils.ResponsWithError(w, 401, "Not Authorised!")
	}
}
