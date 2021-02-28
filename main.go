package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "tanjiro:Kul0nuwun@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Println("Koneksi Sukses")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Ini Service"
	userInput.Email = "service@gmail.com"
	userInput.Occupation = "pujangga"
	userInput.Password = "ini password"

	// setelah itu baru bisa kita pakai service
	userService.RegisterUser(userInput)
	// user := user.User{
	// 	Name: "Hello",
	// }

	// userRepository.Save(user)

	// input dari user
	// handler : mapping input dari user menjadi struct input
	// service : melakukan input dari struct input ke struct user
	// repository : kemudian di passing disimpan ke db
	// db

}
