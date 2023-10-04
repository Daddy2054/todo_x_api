package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
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
	// fmt.Println("your tasks are", tasks)
	fmt.Println("Got all tasks")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
//	fmt.Println(r.Body)

	 w.Header().Set("Content-Type", "application/json")

	// _ = json.NewDecoder(r.Body).Decode(&task)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var task Tasks
	//var book models.Book
	json.Unmarshal(body, &task)
	//fmt.Println(body)

	fmt.Println(task)
	fmt.Println(tasks)

	maxNum := big.NewInt(10000000000000000)
	randNum, _ := rand.Int(rand.Reader, maxNum.Add(maxNum, big.NewInt(1)))
	task.ID = randNum.String()
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode("Created")
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
	//	fmt.Println("Hello there")
	handleRoutes()
}
