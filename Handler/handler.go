package URLShortner

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

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

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := ParseYaml(yamlBytes)
	if err != nil {
		log.Printf("Error unmarshalling YAML: %v", err)
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func ParseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	return pathUrls, err
}
func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := map[string]string{}
	for _, purl := range pathUrls {
		pathsToUrls[purl.Path] = purl.URL
	}
	return pathsToUrls
}
