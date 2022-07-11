// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/gopl", handler) // each request calls handler
	http.HandleFunc("/gopl/count", counter)
	http.HandleFunc("/gopl/image", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			panic(err)
		}

		if cycleStr := request.Form["cycle"]; cycleStr != nil {
			cycle, err := strconv.Atoi(cycleStr[0])
			if err != nil {
				panic(err)
			}
			lissajous(writer, float64(cycle))
		}
		lissajous(writer, 0)

	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(writer http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
