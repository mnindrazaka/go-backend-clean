package album

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Usecase struct {
	repo Repository
}

func UsecaseMake(repo Repository) Usecase {
	return Usecase{repo}
}

func (usecase Usecase) UsecaseGet() ([]T, error) {
	return usecase.repo.RepositoryGet()
}

func (usecase Usecase) UsecaseShow(r *http.Request) (T, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	var formattedId, err = strconv.Atoi(id)

	if err != nil {
		return T{}, err
	}

	return usecase.repo.RepositoryShow(formattedId)
}

func (usecase Usecase) UsecaseCreate(r *http.Request) error {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var album T
	json.Unmarshal(reqBody, &album)

	return usecase.repo.RepositoryCreate(album)
}

func (usecase Usecase) UsecaseUpdate(r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	var album T
	json.Unmarshal(reqBody, &album)

	var formattedId, err = strconv.Atoi(id)

	if err != nil {
		return err
	}

	return usecase.repo.RepositoryUpdate(album, formattedId)
}

func (usecase Usecase) UsecaseDelete(r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	var formattedId, err = strconv.Atoi(id)

	if err != nil {
		return err
	}

	return usecase.repo.RepositoryDelete(formattedId)
}
