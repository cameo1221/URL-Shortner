package main

import (
	"fmt"
	"net/http"

	URLShortner "github.com/cameo1221/URL-Shortner/Handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := URLShortner.MapHandler(pathsToUrls, mux)
	yaml := `
- path: /urlshort
  url: https://google.com
- path: /urlshort-final
  url: https://godoc.org/gopkg.in/yaml.v2`
	yamlHandler, err := URLShortner.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
