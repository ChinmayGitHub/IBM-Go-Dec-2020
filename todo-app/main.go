package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	config "github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/config"
	"github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/models"
	"github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/routes"
)

var err error

func main(){
	config.DB, err := gorm.Open("mysql", config.DbURL(config.BuildConfig()))

	if err != nil {
		log.Fatalf("Error connecting to database : %s", err.Error())
	}
	defer config.DB.Close()
	r := routes.SetupRouter()

	r.run()
}