package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TestType struct {
	Isvalid bool
	Value   string
	Prop    string
}

func test(res http.ResponseWriter, req *http.Request) {
	data := &TestType{
		Isvalid: true,
		Value:   "test",
		Prop:    "test prop",
	}

	log.Output(2, data.Prop)
	log.Output(2, "ok ok ")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.Write(j)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", test).Methods(http.MethodGet)
	http.ListenAndServe(":3000", r)
}
