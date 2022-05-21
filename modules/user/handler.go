package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	usecase Usecase
}

func HandlerMake(usecase Usecase) Handler {
	return Handler{usecase}
}

func (handler Handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	albums, err := handler.usecase.UsecaseGet(r)

	if err != nil {
		w.Write([]byte("Failed to get data"))
	}

	json.NewEncoder(w).Encode(albums)
}

func (handler Handler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	tokenString, err := handler.usecase.UsecaseRegister(r)

	if err != nil {
		w.Write([]byte("Failed to register"))
	}

	w.Write([]byte(tokenString))
}

func (handler Handler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	tokenString, err := handler.usecase.UsecaseLogin(r)

	if err != nil {
		w.Write([]byte("Failed to login"))
	}

	w.Write([]byte(tokenString))
}
