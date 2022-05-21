package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-backend-clean/modules/album"
	"go-backend-clean/modules/user"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

var db *gorm.DB

func main() {
	dsn := "root:roottoor@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleHome)

	userRepository := user.RepositoryMake(db)
	userUsecase := user.UsecaseMake(userRepository)
	userHandler := user.HandlerMake(userUsecase)
	user.RouterMake(userHandler, router)

	albumRepository := album.RepositoryMake(db)
	albumUsecase := album.UsecaseMake(albumRepository)
	albumHandler := album.HandlerMake(albumUsecase)
	album.RouterMake(albumHandler, router)

	http.ListenAndServe(":3000", router)
}
