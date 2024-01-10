package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("http://localhost:3232/cotacao")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
