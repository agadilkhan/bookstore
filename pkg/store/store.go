package store

import (
	"github.com/agadilkhan/simple-rest-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func (s *Store) Connection() {
	dsn := "host=db port=5432 user=postgres password=Alfarabi2004 dbname=go_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Book{})
	s.DB = db
}
func (s *Store) GetBooks() (*[]models.Book, error) {
	var books []models.Book
	_, err := s.DB.Find(&books).Rows()
	if err != nil {
		return nil, err
	}
	return &books, nil
}
func (s *Store) CreateBook(book *models.Book) error {
	result := s.DB.Create(book)
	return result.Error
}
func (s *Store) GetBookByID(id int) (*models.Book, error) {
	var book models.Book
	err := s.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func (s *Store) DeleteBook(id int) error {
	var book models.Book
	result := s.DB.Delete(&book, id)
	return result.Error
}
func (s *Store) UpdateBook(id int, updatedBook *models.Book) error {
	var book models.Book
	s.DB.First(&book, id)
	result := s.DB.Model(&book).Updates(updatedBook)
	result.Save(&book)
	return result.Error
}
func (s *Store) SearchBook(title string) (*models.Book, error) {
	var book models.Book
	result := s.DB.Where("title LIKE ?", title + "%").First(&book)
	return &book, result.Error
}
func (s *Store) OrderBy(order string) (*[]models.Book, error) {
	var books []models.Book
	result := s.DB.Order("cost " + order).Find(&books)
	return &books, result.Error
}