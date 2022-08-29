package lesson4

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func Lesson4() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/go-to-the-moon", mainPage).Methods("GET")
	r.HandleFunc("/v1/comment", commentHandler).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		println("err")
	}
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
	txtFileName := "chat_log.txt"
	stringCommentText := time.Now().Format("2006-01-02 15:04:05") + " From: " + comment.Name + " Comment: " + comment.Comment + "\r\n"
	file, err := os.OpenFile(txtFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("not create file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("not close file:", err)
		}
	}(file)
	_, err = file.WriteString(stringCommentText)
	if err != nil {
		fmt.Println("not write in file:", err)
	}
}
