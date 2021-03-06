package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type redirect struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirect, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
		}
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

// Handler is an interface to a handler function
type Handler func([]byte, http.Handler) (http.HandlerFunc, error)

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var routes []redirect
	err := yaml.Unmarshal(yml, &routes)
	if err != nil {
		fmt.Printf("cannot unmarshal yaml data: %v", err)
		return nil, err
	}
	pathsToUrls := buildPathMap(routes)
	return MapHandler(pathsToUrls, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc that will attempt to map any paths to
// their corresponding URL.
func JSONHandler(jsonBlob []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var routes []redirect
	err := json.Unmarshal(jsonBlob, &routes)
	if err != nil {
		fmt.Printf("cannot unmarshal json data: %v", err)
	}
	pathsToUrls := buildPathMap(routes)
	return MapHandler(pathsToUrls, fallback), nil
}

func buildPathMap(routes []redirect) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, r := range routes {
		pathsToUrls[r.Path] = r.URL
	}
	return pathsToUrls
}
