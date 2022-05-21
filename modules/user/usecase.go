package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type Usecase struct {
	repo Repository
}

func UsecaseMake(repo Repository) Usecase {
	return Usecase{repo}
}

func (usecase Usecase) UsecaseGet(r *http.Request) ([]T, error) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("my secret"), nil
	})

	if err != nil {
		return []T{}, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		users, err := usecase.repo.RepositoryGet()

		if err != nil {
			return []T{}, err
		}

		return users, nil
	} else {
		return []T{}, err
	}
}

func (usecase Usecase) UsecaseRegister(r *http.Request) (string, error) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var user T
	json.Unmarshal(reqBody, &user)

	err := usecase.repo.RepositoryCreate(user)

	if err != nil {
		return "", err
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	}).SignedString([]byte("my secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (usecase Usecase) UsecaseLogin(r *http.Request) (string, error) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var user T
	json.Unmarshal(reqBody, &user)

	users, err := usecase.repo.RepositoryGetBy(user)

	if err != nil {
		return "", err
	}

	if len(users) > 0 {
		tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
		}).SignedString([]byte("my secret"))

		if err != nil {
			return "", err
		}

		return tokenString, nil
	} else {
		return "", err
	}
}
