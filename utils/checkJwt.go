package utils

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		secretKey := os.Getenv("JWT_SECRET")
// 		if secretKey == "" {
// 			log.Fatal("JWT_SECRET does not exist in environment")
// 		}

// 		tokenString, err := GetApiKey(r)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var mySigningKey = []byte(secretKey)

// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("There was an error in parsing")
// 			}
// 			return mySigningKey, nil
// 		})

// 		if err != nil {
// 			// var err Error
// 			// err = SetError(err, "Your Token has been expired")
// 			// json.NewEncoder(w).Encode(err)
// 			log.Fatal(err)
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if claims["role"] == "admin" {

// 				r.Header.Set("Role", "admin")
// 				handler.ServeHTTP(w, r)
// 				return

// 			} else if claims["role"] == "user" {

// 				r.Header.Set("Role", "user")
// 				handler.ServeHTTP(w, r)
// 				return
// 			}
// 		}
// 		// var reserr Error
// 		// reserr = SetError(reserr, "Not Authorized")
// 		json.NewEncoder(w).Encode(err)
// 	}
// }
