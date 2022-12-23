package storage

import (
	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/config"
	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/models"
	"gorm.io/gorm"
)

var Repo *Repository

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *Repository) FindBooks(books *[]models.Book) error {
	return r.db.Find(books).Error
}

func (r *Repository) FindBookById(book *models.Book, id string) error {
	return r.db.First(book, id).Error
}

func (r *Repository) UpdateBook(book *models.Book, id string, fields *models.Book) error {
	return r.db.Model(book).Where("id = ?", id).Updates(fields).Error
}

func (r *Repository) DeleteBook(book *models.Book, id string) error {
	return r.db.Unscoped().Delete(book, id).Error
}

func CreateRepository() *Repository {
	return &Repository{
		db: config.GetDB(),
	}
}

func SetRepository(repo *Repository) {
	Repo = repo
}
