package user

import "github.com/gorilla/mux"

func RouterMake(handler Handler, router *mux.Router) {
	router.HandleFunc("/login", handler.HandlerLogin).Methods("POST")
	router.HandleFunc("/register", handler.HandlerRegister).Methods("POST")
	router.HandleFunc("/users", handler.HandlerGet)
}
