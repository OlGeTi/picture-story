package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"picture-story/picstore"
)

type Handler struct {
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {

	parameter := request.URL.Query()

	var albums *[]picstore.Album
	var err error

	switch parameter["func"][0] {
	case "albums":
		albums, err = fetchAlbums(parameter["userid"][0])
	default:
		err = errors.New("no function selected")
	}

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, err.Error())
		return
	}

	jsonResult, err := json.Marshal(*albums)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, err.Error())
		return
	}

	fmt.Fprint(rw, jsonResult)
}

func fetchAlbums(userid string) (*[]picstore.Album, error) {
	return picstore.GAlbums(userid)
}
