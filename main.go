package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task1 := Tasks{
		ID:         "1",
		TaskName:   "New project",
		TaskDetail: "You must lead the project and finish it",
		Date:       "2022-01-22",
	}
	task2 := Tasks{
		ID:         "2",
		TaskName:   "Another project",
		TaskDetail: "You must take the project and destroy it",
		Date:       "2022-02-22",
	}
	task3 := Tasks{
		ID:         "3",
		TaskName:   "Wonderful project",
		TaskDetail: "You must leave this project and ruin it",
		Date:       "2022-03-22",
	}
	tasks = append(tasks, task1, task2, task3)
	fmt.Println("your tasks are", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/")
	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}



func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8082", router))
}
func main() {
	allTasks()
	fmt.Println("Hello there")
	handleRoutes()
}
