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

func ListTasks(w http.ResponseWriter, _ *http.Request) {
	var tasks []model.Task
	db.Db.Find(&tasks)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PostFormValue("id")
	var task model.Task
	db.Db.First(&task, id)
	db.Db.Delete(&task)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
