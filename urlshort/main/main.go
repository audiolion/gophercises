package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/audiolion/gophercises/urlshort"
	bolt "go.etcd.io/bbolt"
)

func main() {
	var (
		yamlFilename = flag.String("yaml", "urls.yaml", "yaml file of redirects")
		jsonFilename = flag.String("json", "urls.json", "json file of redirects")
		dbFilename   = flag.String("db", "urls.db", "boltdb file of redirects")
	)
	flag.Parse()

	mux := defaultMux()

	db, err := bolt.Open(*dbFilename, 0600, nil)
	if err != nil {
		fmt.Println("panic")
		panic(err)
	}
	defer db.Close()

	pathsToUrls, err := readRedirects(db)

	if err != nil {
		err = createDefaultRedirects(db)
		pathsToUrls, err = readRedirects(db)
	}

	if err != nil {
		fmt.Println(err)
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

func readRedirects(db *bolt.DB) (map[string]string, error) {
	pathsToUrls := make(map[string]string)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("UrlBucket"))

		if b == nil {
			return fmt.Errorf("UrlBucket does not exist")
		}

		b.ForEach(func(k, v []byte) error {
			var key []byte
			var val []byte
			copy(key, k)
			copy(val, v)
			pathsToUrls[string(key)] = string(val)
			return nil
		})

		return nil
	})

	return pathsToUrls, err
}
func createDefaultRedirects(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("UrlBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}
		for k, v := range pathsToUrls {
			err := b.Put([]byte(k), []byte(v))
			if err != nil {
				fmt.Printf("Error: putting %s, %s into bucket, %s\n", k, v, err)
			}
			fmt.Printf("Added key=%s, val=%s\n", k, v)
		}
		return nil
	})
	return err
}
