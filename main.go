package main

import (
	"html/template"
	"log"
	"net/http"

	"htmx.dewey-typical.io/db"
	"htmx.dewey-typical.io/handler"
	"htmx.dewey-typical.io/model"
)

func main() {
	db := db.CreateDb()
	db.AutoMigrate(&model.Task{})
	//h1 := func(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("index.html"))
	//tmpl.Execute(w, )
	//}

	/* 	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		description := r.PostFormValue("description")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "task-list-element", Task{Description: description})
	} */
	http.HandleFunc("/", listTasks)
	http.HandleFunc("/add-task/", handler.CreateTask)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func listTasks(w http.ResponseWriter, _ *http.Request) {
	var tasks []model.Task
	db.Db.Find(&tasks)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, tasks)
}
