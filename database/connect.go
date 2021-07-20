package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func ConnectDb() error {
	var err error
	//dsn := "host=localhost user=postgres password=docker dbname=postgres port=5432 sslmode=disable"
	//var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=docker")

	if err != nil {
		fmt.Println("veritabanına bağlanılamadı.")
		panic("failed : veritabanına bağlanılamadı.")
	}

	/*user := new(models.User)
	user.Name= "yadi"
	user.Email = "aydngltp"
	user.Password="1234567"
	user.Surname="gultepe"

	DB.Create(&user)
	var user2 models.User
	DB.Find(&user2, 1)
	fmt.Println()
	fmt.Println(user2)
	fmt.Println("Veritabanına başarıyla bağlanıldı.")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed migrate")
	}
	fmt.Println("Veritabanı migrate edildi.")

	var users []models.User
	DB.Find(&users)
	fmt.Println(users)

	for i,v :=  range users {
		fmt.Println(i,v)
	}
	return*/
	return nil
}
