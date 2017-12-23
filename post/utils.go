package post

import (
	"fmt"
	"net/http"
)

// IndexHandler returns data to the default port
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("incorrect path: %s", r.URL.Path)))
		return
	}

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%v", `¯\_(ツ)_/¯`)))
		return
	case "POST":
		postHandler(w, r)
	default:
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}
