package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var todos []Todo = make([]Todo, 0, 10)

type Todo struct {
	Title string
	Done  bool
}

type InputtingTodo struct {
	Title string
}

type Success struct {
	Success bool
}

func get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	j, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	res.Write(j)
}

func post(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	var body InputtingTodo
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	todo := Todo{Title: body.Title, Done: false}

	todos = append(todos, todo)

	j, err := json.Marshal(Success{Success: true})
	if err != nil {
		panic(err)
	}
	res.Write(j)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", get).Methods(http.MethodGet)
	r.HandleFunc("/todos", post).Methods(http.MethodPost)
	http.ListenAndServe(":3000", r)
}
