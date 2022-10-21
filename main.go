package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized app",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		var message string
		if name == "" {
			message = "Hello World"
		} else {
			message = "Hello " + name
		}
		response := map[string]string{
			"message": message,
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/getJson", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		user := &Person{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("got user:", user)
		w.WriteHeader(http.StatusCreated)
	})

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}

	log.Println("Server is running!")
	fmt.Println(http.ListenAndServe(":8081", router))
}
