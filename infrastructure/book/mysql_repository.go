package book

import (
	"book-ddd/domain/book"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) book.Repository {
	return &mysqlRepo{db}
}

func (r *mysqlRepo) FindAll() ([]book.Book, error) {
	var books []book.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *mysqlRepo) FindByID(id uint) (book.Book, error) {
	var b book.Book
	err := r.db.First(&b, id).Error
	return b, err
}

func (r *mysqlRepo) Create(b book.Book) (book.Book, error) {
	err := r.db.Create(&b).Error
	return b, err
}

func (r *mysqlRepo) Update(b book.Book) (book.Book, error) {
	err := r.db.Save(&b).Error
	return b, err
}

func (r *mysqlRepo) Delete(id uint) error {
	return r.db.Delete(&book.Book{}, id).Error
}
