package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Currency struct {
	valorConvertido int    `json:"valorConvertido"`
	simboloMoeda    string `json:"simboloMoeda"`
}

func main() {
	uri := mux.NewRouter()

	uri.HandleFunc("exchange/{amount}/{from}/{to}/{rate}", conversion).Methods("GET")
}

func conversion(writer http.ResponseWriter, request *http.Request) {
	symbols := map[string]sting{
		"USD": "U$$",
		"EUR": "€",
		"BTC": "₿",
	}

	db, err := sql.Open("mysql", "test:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	fromCurency := params["from"]
	toCurency := params["to"]
	amount := params["amount"]
	rate := params["rate"]
	if fromCurency == "BRL" {
		if toCurency == "USD" || toCurency == "EUR" || toCurency == "BTC" {
			symbol := symbols[params["to"]]
			insert, err := db.Query("INSERT INTO CURRENCY VALUES ( '", fromCurency, "', '", toCurency, "',", amount, ",", rate, " )")
			if err != nil {
				panic(err.Error())
			}
			apiResponse(symbol, amount, params["rate"])
		}
	} else if fromCurency == "USD" || fromCurency == "EUR" || fromCurency == "BTC" {
		if toCurency == "BRL" {
			symbol := "R$"
			insert, err := db.Query("INSERT INTO CURRENCY VALUES ( '", fromCurency, "', '", toCurency, "',", amount, ",", rate, " )")
			if err != nil {
				panic(err.Error())
			}
			apiResponse(symbol, amount, rate)
		}
	}

}

func apiResponse(writer http.ResponseWriter, symbol string, amount int, rate int) {
	calc := amount * rate
	var exchange Currency
	exchange.valorConvertido = calc
	exchange.simboloMoeda = symbol
	json.NewEncoder(writer).Encode(exchange)
}
