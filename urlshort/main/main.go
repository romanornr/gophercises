package main

import "net/http"
import ".."

func main()  {
	mux := http.DefaultServeMux

	pathsToUrls := map[string]string{
		"bittrex": "https://bittrex.com",
		"binance": "https://binance.com",
	}
	mapHander := urlshort.MapHandler(pathsToUrls, mux)
}