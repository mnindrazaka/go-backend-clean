package album

import "github.com/gorilla/mux"

func RouterMake(handler Handler, router *mux.Router) {
	router.HandleFunc("/albums", handler.HandlerCreate).Methods("POST")
	router.HandleFunc("/albums", handler.HandlerGet)
	router.HandleFunc("/albums/{id}", handler.HandlerUpdate).Methods("PUT")
	router.HandleFunc("/albums/{id}", handler.HandlerDelete).Methods("DELETE")
	router.HandleFunc("/albums/{id}", handler.HandlerShow)
}
