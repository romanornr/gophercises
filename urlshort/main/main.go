package main

import "net/http"
import (
	"github.com/romanornr/gophercises/urlshort"
	"fmt"
)

func main()  {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"bittrex": "https://bittrex.com",
		"binance": "https://binance.com",
	}
	mapHander := urlshort.MapHandler(pathsToUrls, mux)
	http.ListenAndServe(":8080", mapHander)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}