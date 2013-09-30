// APIHandler
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"picture-story/google"
)

type APIHandler struct {
}

func (h APIHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {

	parameter := request.URL.Query()

	var result interface{}
	var err error

	switch parameter["func"][0] {
	case "albums":
		result, err = fetchAlbums(parameter["userid"][0])
	default:
		err = errors.New("no function selected")
	}

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, err.Error())
		return
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, err.Error())
		return
	}

	fmt.Fprint(rw, jsonResult)
}

func fetchAlbums(userid string) (interface{}, error) {

	albums, err := google.Albums(userid)

	return albums, err
}
