package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id      string `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Todos []Todo

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("[Home Page]")

}

func getTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("[Get Todos]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("[Get Todo]")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["id"]
	for _, todo := range Todos {
		if todo.Id == key {
			json.NewEncoder(w).Encode(todo)
			break
		}
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("[Create New Todo]")
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo Todo
	json.Unmarshal(reqBody, &todo)
	Todos = append(Todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("[Delete Todo]")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["id"]

	for index, todo := range Todos {
		if todo.Id == key {
			Todos = append(Todos[:index], Todos[index+1:]...)
			break
		}
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("[Update Todo]")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newTodo Todo
	json.Unmarshal(reqBody, &newTodo)

	for index, todo := range Todos {
		if todo.Id == key {
			Todos[index] = newTodo
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/", homePage)
	router.HandleFunc("/api/todos", getTodos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/api/todo/", createTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")
	router.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("Go Rest API - Mux Routers")

	Todos = []Todo{
		{Id: "1", Author: "Anna", Title: "Travel", Content: "Travel to France"},
		{Id: "2", Author: "John", Title: "Sports", Content: "Swimming"},
	}

	handleRequests()
}
