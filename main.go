package main

import (
	"html/template"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
}

var (
	db  *gorm.DB
	err error
)

func main() {
	db, err = gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Task{})

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
	http.HandleFunc("/add-task/", createTask)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func listTasks(w http.ResponseWriter, _ *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	description := r.PostFormValue("description")

	task := Task{Description: description}
	db.Create(&task)

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "task-list-element", Task{Description: description})
}
