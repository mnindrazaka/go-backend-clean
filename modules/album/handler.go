package album

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
	albums, err := handler.usecase.UsecaseGet()

	if err != nil {
		w.Write([]byte("Failed to get data"))
	}

	json.NewEncoder(w).Encode(albums)
}

func (handler Handler) HandlerShow(w http.ResponseWriter, r *http.Request) {
	album, err := handler.usecase.UsecaseShow(r)

	if err != nil {
		w.Write([]byte("Failed to get data"))
	}

	json.NewEncoder(w).Encode(album)
}

func (handler Handler) HandlerCreate(w http.ResponseWriter, r *http.Request) {
	err := handler.usecase.UsecaseCreate(r)

	if err != nil {
		w.Write([]byte("Failed to create data"))
	}

	w.Write([]byte("Success to create data"))
}

func (handler Handler) HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	err := handler.usecase.UsecaseUpdate(r)

	if err != nil {
		w.Write([]byte("Failed to update data"))
	}

	w.Write([]byte("Success to update data"))
}

func (handler Handler) HandlerDelete(w http.ResponseWriter, r *http.Request) {
	err := handler.usecase.UsecaseDelete(r)

	if err != nil {
		w.Write([]byte("Failed to delete data"))
	}

	w.Write([]byte("Success to delete data"))
}
