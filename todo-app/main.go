package main

import (
	"log"

	"github.com/jinzhu/gorm"
	config "github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/config"
	models "github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/models"
	routes "github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildConfig()))

	if err != nil {
		log.Fatalf("Error connecting to database : %s", err.Error())
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})
	r := routes.SetupRouter()

	r.Run()
}
