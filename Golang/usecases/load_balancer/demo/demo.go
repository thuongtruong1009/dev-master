package main

import (
	"fmt"
	"net/http"
	"sync"
	"zoomer/pkg/load_balancer"
)

func serveBackend(name string, port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Backend server name: %v\n", name)
		fmt.Fprintf(w, "Response headers: %v\n", r.Header)
	}))
	http.ListenAndServe(port, mux)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)

	go func() {
		load_balancer.LoadBalancer()
		wg.Done()
	}()



	go func() {
		serveBackend("backend3", ":8082")
		wg.Done()
	}()

	go func() {
		serveBackend("backend4", ":8083")
		wg.Done()
	}()

	go func() {
		serveBackend("backend5", ":8084")
		wg.Done()
	}()

	go func() {
		serveBackend("backend2", ":8085")
		wg.Done()
	}()

	wg.Wait()
}
