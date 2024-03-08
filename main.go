package main

import (
	"fmt"
	"net/http"
	"github.com/cameo1221/URL-Shortener/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// yaml := `
	// -path: /urlshort
	// url: http://www.google.co.uk/
	// -path: /urlshort-final
	// url: https`
	// yamlHandler. err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil{
	// 	panic(err)
	// }
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)

}
