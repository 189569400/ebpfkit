package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthCheckHandler)
	r.HandleFunc("/api/products", ProductsHandler).Methods("POST")
	r.HandleFunc("/api/articles", ArticlesHandler)
	r.NotFoundHandler = NotFoundHandler()
	r.MethodNotAllowedHandler = MethodNotAllowedHandler()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func logPrefix(r *http.Request) string {
	var url string
	if r.URL != nil {
		url = r.URL.String()
	} else {
		url = r.RequestURI
	}

	return fmt.Sprintf("%s - %s %s", r.RemoteAddr, r.Method, url)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(NewMessage("OK", http.StatusOK))
	w.Write(data)
	log.Printf("%s - %d\n", logPrefix(r), http.StatusOK)
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%s - couldn't read request body !\n", logPrefix(r))
	}
	w.Write([]byte("Product ok !\n"))
	log.Printf("%s - (%d) %s - %d\n", logPrefix(r), len(data), string(data), http.StatusOK)
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%s - couldn't read request body !\n", logPrefix(r))
	}
	w.Write([]byte("Articles ok !\n"))
	log.Printf("%s - (%d) %s - %d\n", logPrefix(r), len(data), string(data), http.StatusOK)
}

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %d\n", logPrefix(r), http.StatusNotFound)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
}

func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %d\n", logPrefix(r), http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})
}