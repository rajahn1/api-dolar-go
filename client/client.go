package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3232/cotacao", nil)

	if err != nil {
		log.Fatal("Error doing request with context", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal("Error doing request", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Error reading body:", err)
	}

	file, err := os.Create("cotacao.txt")

	if err != nil {
		log.Fatal(err, "Error creating file")
	}

	_, err = file.WriteString(fmt.Sprintf("Dólar: %v", string(body)))

	if err != nil {
		log.Fatal(err, "Error writing to file")
	}

	fmt.Println(string(body))
}
