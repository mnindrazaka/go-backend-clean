package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	con *gorm.DB
}

func RepositoryMake(con *gorm.DB) Repository {
	return Repository{con}
}

func (repo Repository) RepositoryGet() ([]T, error) {
	var users []T
	result := repo.con.Table("user").Find(&users)
	return users, result.Error
}

func (repo Repository) RepositoryGetBy(user T) ([]T, error) {
	var users []T
	result := repo.con.Table("user").Where(&user).Find(&users)
	return users, result.Error
}

func (repo Repository) RepositoryCreate(user T) error {
	result := repo.con.Table("user").Create(&user)
	return result.Error
}
