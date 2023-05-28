// Router for the app
package routes

import (
	"caching-with-redis/services"

	"github.com/gorilla/mux"
)

// func CreateRouter() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
// 	return router
// }

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts/{id}", services.GetPost).Methods("GET")

	return router
}
