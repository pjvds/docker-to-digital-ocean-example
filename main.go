package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var (
	counters = make(map[string]uint64)
	host     = flag.String("host", "localhost", "the host to bind on")
	port     = flag.Int("port", 8080, "the port to bind on")
)

func main() {
	flag.Parse()

	counterRouter := mux.NewRouter()
	counterRouter.HandleFunc("/counter/{name}", CounterHandler)

	address := fmt.Sprintf("%v:%v", *host, *port)

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
