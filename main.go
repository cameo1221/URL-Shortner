package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	URLShortner "github.com/cameo1221/URL-Shortner/Handler"
)

func main() {
	mux := defaultMux()
	db, err := sql.Open("postgres", "user=postgres password=url1234 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()
	// jsonFilePath := flag.String("json", "", "URL-Shortner/jsonfile.json")
	// flag.Parse()

	// if *jsonFilePath == "" {
	// 	log.Fatal("Please provide a YAML file path using the -yaml flag")
	// }
	// // Open the file and handle any errors
	// file, err := os.Open(*jsonFilePath)
	// if err != nil {
	// 	log.Fatalf("Error opening JSON file: %v", err)
	// }
	// defer file.Close()

	// // Now file is an io.Reader, so we can pass it to io.ReadAll
	// jsonBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Fatalf("Error reading YAML file: %v", err)
	// }

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := URLShortner.MapHandler(pathsToUrls, mux)
	// yamlHandler, err := URLShortner.JsonHandler(jsonBytes, mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
	dbHandler, err := URLShortner.DBHandler(db, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
