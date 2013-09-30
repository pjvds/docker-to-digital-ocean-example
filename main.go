package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var (
	counters = make(map[string]uint64)
)

func main() {
	counterRouter := mux.NewRouter()
	counterRouter.HandleFunc("/counter/{name}", CounterHandler)

	address := "localhost:8080"

	println("Hosting at " + address)
	if err := http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, counterRouter)); err != nil {
		println("Error while serving: " + err.Error())
	}
}

func CounterHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]

	var count uint64
	if current, ok := counters[name]; ok {
		count = current + 1
	} else {
		count = 1
	}

	counters[name] = count
	fmt.Fprintf(response, "%v", count)
}
