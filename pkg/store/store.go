package main

import (
	"github.com/agadilkhan/rest-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func (s *Store) Init() {
	dsn := "user=postgres password=Alfarabi2004 dbname=go_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Book{})
	s.DB = db
}
