package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"time"
)

// Adapter :
type Adapter func(http.Handler) http.Handler

// AllowCors :
func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

// HasLoggedIn : this function checks the validity of the access token
// func HasLoggedIn() Adapter {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// #1 Check Access token,
// 			token := r.Header.Get("Authorization")
// 			platform := r.URL.Query().Get("platform")
// 			if token == "" {
// 				response.With412mV2(w, "Unauthorized", platform)
// 				log.Println("ERROR : middleware - HasLoggedIn - no token")
// 				return
// 			}

// 			auth, err := services.ValidateToken(token)
// 			if err != nil {
// 				log.Println("ERROR : middleware - cannot validate token - " + err.Error())
// 				response.With412mV2(w, "Unauthorized", platform)

// 				return
// 			}
// 			context.Set(r, "loggedIn", true)
// 			context.Set(r, constants.CONTEXTJWTKEY, *auth)
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

//Log : ""
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		platform := r.URL.Query().Get("platform")
		log.Println("platform ==>", platform)
		next.ServeHTTP(w, r)
		duration := time.Since(t)
		log.Println("API ==>", r.RequestURI, " Time taken ===> ", duration.Minutes(), "m")
		fmt.Println()
		fmt.Println()
		fmt.Println()
	})
}
