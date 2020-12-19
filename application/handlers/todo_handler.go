package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	model "github.com/S1lvesterTake/simple_todo/application/models"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func CreateTodoHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		description := r.FormValue("description")

		newTodo := &model.TodoItem{Description: description, IsCompleted: false}
		db.Create(&newTodo)
		result := db.Last(&newTodo)

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(result.Value)
	}
	return http.HandlerFunc(fn)
}

func GetListTodoHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		todoList := []model.TodoItem{}

		db.Find(&todoList)

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(todoList)
	}
	return http.HandlerFunc(fn)
}

func GetTodoByIDHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var todo model.TodoItem
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		db.Where("id = ?", id).First(&todo)

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
	return http.HandlerFunc(fn)
}

func UpdateTodoHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var todo model.TodoItem
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		description := r.FormValue("description")
		isCompleted, _ := strconv.ParseBool(r.FormValue("iscompleted"))

		db.Where("id = ?", id).First(&todo)
		todo.Description = description
		todo.IsCompleted = isCompleted

		db.Save(&todo)
		log.WithFields(log.Fields{"ID": id, "Description": description, "IsCompleted": isCompleted}).Info("Success updating todo item")
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
	return http.HandlerFunc(fn)
}

func DeleteTodoHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var todo model.TodoItem
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		db.Where("id = ?", id).First(&todo)

		db.Delete(&todo)
		log.WithFields(log.Fields{"ID": id}).Info("Success delete todo item")
		w.Header().Set("content-type", "application/json")
		io.WriteString(w, `{"Success":true}`)
	}
	return http.HandlerFunc(fn)
}
