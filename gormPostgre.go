package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Veri yapısını tanımlama
type User struct {
	gorm.Model
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

var db *gorm.DB

func main() {
	db = initDB()
	app := fiber.New()

	// Route oluşturuldu.
	app.Get("/UserGet", GetUsers)
	app.Get("/UserNew", NewUser)
	app.Get("/UserDelete", DeleteUser)
	app.Get("/UserUpdate", UpdateUser)

	app.Listen(3000) // 3000. porttan serv edildi
}

func UpdateUser(ctx *fiber.Ctx) {
	db = initDB()
	var user User
	user.Name = "Joseph"
	user.Age = 16

	db.Update(&user)
	ctx.JSON(&user)

	defer db.Close()
}

func DeleteUser(ctx *fiber.Ctx) {
	db = initDB()
	db.Where("ID = ?", 1).Delete(&User{})
	defer db.Close()
}

func initDB() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost user=postgres password=Q1w2E3r4 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		log.Println("57", err)
		panic("Error")
	}
	fmt.Println("Database Connected")
	db.AutoMigrate(&User{})
	return db
}

func GetUsers(ctx *fiber.Ctx) {
	db = initDB()
	var user []User
	db.Find(&user)
	ctx.JSON(&user)

	defer db.Close()
}
func NewUser(ctx *fiber.Ctx) {

	var user User
	user.Name = "James"
	user.Age = 21
	db.Create(&user)
	fmt.Println("Yazıldı")
	ctx.JSON(&user)
	defer db.Close()
}
