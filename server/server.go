package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
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
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handlerQuote)
	http.ListenAndServe(":3232", mux)
}

func getQuote() (*Response, error) {
	const apiUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	res, err := http.Get(apiUrl)

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var quote Response
	err = json.Unmarshal(body, &quote)

	if err != nil {
		panic(err)
	}

	return &quote, nil
}

func handlerQuote(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, err := getQuote()

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.USDBRL)
}
