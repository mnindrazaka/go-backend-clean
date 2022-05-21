package album

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
	var albums []T
	result := repo.con.Table("album").Find(&albums)
	return albums, result.Error
}

func (repo Repository) RepositoryShow(id int) (T, error) {
	var album = T{ID: id}
	result := repo.con.Table("album").First(&album)
	return album, result.Error
}

func (repo Repository) RepositoryCreate(album T) error {
	result := repo.con.Table("album").Create(&album)
	return result.Error
}

func (repo Repository) RepositoryUpdate(album T, id int) error {
	result := repo.con.Table("album").Save(&T{ID: id, Title: album.Title, Artist: album.Artist, Price: album.Price})
	return result.Error
}

func (repo Repository) RepositoryDelete(id int) error {
	result := repo.con.Table("album").Delete(&T{ID: id})
	return result.Error
}
