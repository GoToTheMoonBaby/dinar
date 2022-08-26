package lesson4

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Lesson4() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/go-to-the-moon", mainPage).Methods("GET")
	r.HandleFunc("/", handleGet).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		println("err")
	}
	/*http.HandleFunc("/v1/go-to-the-moon", mainPage)
	http.HandleFunc("/v1/comment", commentHandler)
	port := ":9090"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		println("err")
	}*/
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Go"))
}

type Comment struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
	comment := &Comment{}
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(*comment)
}
