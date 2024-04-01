package handler

import (
	"net/http"
	"text/template"

	"htmx.dewey-typical.io/db"
	"htmx.dewey-typical.io/model"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	description := r.PostFormValue("description")

	task := model.Task{Description: description}
	db.Db.Create(&task)

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "task-list-element", model.Task{Description: description})
}
