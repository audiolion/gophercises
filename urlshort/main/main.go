package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/audiolion/gophercises/urlshort"
)

func main() {
	var (
		yamlFilename = flag.String("yaml", "urls.yaml", "yaml file of redirects")
		jsonFilename = flag.String("json", "urls.json", "json file of redirects")
	)
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler := makeHandler(yamlFilename, urlshort.YAMLHandler, mapHandler)

	jsonHandler := makeHandler(jsonFilename, urlshort.JSONHandler, yamlHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func makeHandler(filename *string, createHandler urlshort.Handler, fallback http.HandlerFunc) http.HandlerFunc {
	data := readFile(*filename)
	handler, err := createHandler(data, fallback)
	if err != nil {
		panic(err)
	}
	return handler
}
