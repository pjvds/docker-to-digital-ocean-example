package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	counters = make(map[string]uint64)
	host     = flag.String("host", "localhost", "the host to bind on")
	port     = flag.Int("port", 8080, "the port to bind on")
	mutex    sync.Mutex
)

func main() {
	flag.Parse()

	counterRouter := mux.NewRouter()
	counterRouter.HandleFunc("/{name}", CounterHandler)

	address := fmt.Sprintf("%v:%v", *host, *port)

	println("Hosting at " + address)
	if err := http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, counterRouter)); err != nil {
		println("Error while serving: " + err.Error())
	}
}

func CounterHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]

	fmt.Fprintf(response, "%v", count(name))
}

func count(name string) {
	lower := strings.ToLower(name)
	mutex.Lock()
	defer mutex.Unlock()

	var count uint64
	if current, ok := counters[lower]; ok {
		count = current + 1
	} else {
		count = 1
	}
	counters[lower] = count
}
