package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	URLShortner "github.com/cameo1221/URL-Shortner/Handler"
)

func main() {
	mux := defaultMux()
	jsonFilePath := flag.String("json", "", "URL-Shortner/jsonfile.json")
	flag.Parse()

	if *jsonFilePath == "" {
		log.Fatal("Please provide a YAML file path using the -yaml flag")
	}
	// Open the file and handle any errors
	file, err := os.Open(*jsonFilePath)
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Now file is an io.Reader, so we can pass it to io.ReadAll
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := URLShortner.MapHandler(pathsToUrls, mux)
	yamlHandler, err := URLShortner.JsonHandler(jsonBytes, mapHandler)
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
