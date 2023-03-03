package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Payload struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	To   string `json:"to"`
}

func main() {

	url := "https://api3.greensms.ru/call/send"

	p := Payload{
		User: "dinar",
		Pass: "dinar8888",
		To:   "79178710847",
	}

	structInString, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// POST method with token based auth is recommended. But GET is also supported.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(structInString))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
