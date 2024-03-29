package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type QuotationResponse struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	mux := http.NewServeMux()
	var err error
	DB, err = sql.Open("sqlite3", "./cotacoes.db")

	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	defer DB.Close()

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS cotacoes(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		timestamp INTEGER
	);`)

	if err != nil {
		log.Fatal("error trying to create table", err)
	}

	mux.HandleFunc("/", handleWelcome)
	mux.HandleFunc("/cotacao", handleQuote)
	http.ListenAndServe(":3232", mux)
}

func handleQuote(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := getExchangeRate(ctx)

	if err != nil {
		log.Fatal(err)
		return
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()
	insertQuote(ctxDB, res)

	response := map[string]string{"valor": res.USDBRL.Bid}
	json.NewEncoder(w).Encode(response)
}

func insertQuote(ctx context.Context, quotation *QuotationResponse) {
	_, err := DB.ExecContext(ctx, `INSERT INTO cotacoes(bid, timestamp) VALUES (?, ?);`, quotation.USDBRL.Bid, time.Now().Unix())
	if err != nil {
		log.Fatal("Error inserting into cotacoes context:", err)
		return
	}
}

func getExchangeRate(ctx context.Context) (*QuotationResponse, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		log.Fatal("error requesting api:", err)
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal("error doing client res", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("error reading body:", err)
	}

	var quotation QuotationResponse

	err = json.Unmarshal(body, &quotation)

	if err != nil {
		return nil, err
	}

	return &quotation, nil
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, para ver a cotação do dólar real acesse '/cotacao'"))
}
