package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"picture-story/picstore"
	"strings"
)

type Handler struct{}

func (h Handler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {

	path := strings.Split(request.URL.Path, "/")[1:] // First element is ""

	if cap(path) < 2 {
		writeUnvalidAPICallError(rw)
		return
	}

	switch path[1] {
	case "google":
		lookupGoogle(rw, path[2:])
	default:
		writeUnvalidAPICallError(rw)
	}
}

func writeUnvalidAPICallError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, "Not a valid API-Call")
}

func lookupGoogle(rw http.ResponseWriter, args []string) {

	// Google API-Call should look like http:.../api/google/<userid>[/<albumid>[/<picid>]]

	var result interface{}
	var err error

	switch cap(args) {
	case 1:
		result, err = picstore.GAlbums(args[0]) // args[0] should be the userid

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err.Error())
			return
		}

	case 2:
		result = nil
	case 3:
		result = nil
	default:
		writeUnvalidAPICallError(rw)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err.Error())
		return
	}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err.Error())
		return
	}

	fmt.Fprint(rw, string(jsonResult))
}
