package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"practicing-go/models"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=sucuk password=sucuk123 dbname=practising-go port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("veritabanına bağlanılamadı.")
		panic("failed : veritabanına bağlanılamadı.")
	}

	fmt.Println("Veritabanına başarıyla bağlanıldı.")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed migrate")
	}
	fmt.Println("Veritabanı migrate edildi.")
}
