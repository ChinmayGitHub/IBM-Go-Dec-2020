package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/config"
	"github.com/tkmagesh/IBM-Go-Dec-2020/todo-app/models"
)

func GetAllTodos(todo *[]models.Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *models.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *models.Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *models.Todo, id string) (err error) {
	config.DB.save(todo)
	return nil
}

func DeleteATodo(todo *models.Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
