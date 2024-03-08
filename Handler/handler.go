package URLShortner

import (
	"database/sql"
	"net/http"
)

// type pathUrl struct {
// 	Path string `json:"path"`
// 	URL  string `json:"url"`
// }

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)

	}
}

//	func JsonHandler(JsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
//		pathUrls, err := Parsejson(JsonBytes)
//		if err != nil {
//			log.Printf("Error unmarshalling JSON: %v", err)
//			return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
//		}
//		pathsToUrls := buildMap(pathUrls)
//		return MapHandler(pathsToUrls, fallback), nil
//	}
func DBHandler(db *sql.DB, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls := make(map[string]string)

	rows, err := db.Query("SELECT path, url FROM path_urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var path, url string
		if err := rows.Scan(&path, &url); err != nil {
			return nil, err
		}
		pathsToUrls[path] = url
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return MapHandler(pathsToUrls, fallback), nil
}

// func Parsejson(data []byte) ([]pathUrl, error) {
// 	var pathUrls []pathUrl
// 	err := json.Unmarshal(data, &pathUrls)
// 	return pathUrls, err
// }
// func buildMap(pathUrls []pathUrl) map[string]string {
// 	pathsToUrls := map[string]string{}
// 	for _, purl := range pathUrls {
// 		pathsToUrls[purl.Path] = purl.URL
// 	}
// 	return pathsToUrls
// }
