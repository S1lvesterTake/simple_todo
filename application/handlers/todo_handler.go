package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/S1lvesterTake/simple_todo/application/models"
	"github.com/jinzhu/gorm"
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
