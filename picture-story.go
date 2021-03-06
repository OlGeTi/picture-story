// picture-story project main.go
package main

import (
	"github.com/stvp/go-toml-config"
	"log"
	"net/http"
	"picture-story/api"
	"picture-story/story"
)

var (
	dings = config.Bool("dings", false)
)

func main() {

	log.Println("Startup...")

	http.Handle("/api/", api.Handler{})
	http.Handle("/story", story.Handler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
