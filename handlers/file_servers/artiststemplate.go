package ConcertAPI

import (
	"net/http"
)

func ArtistTemplate(w http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	http.ServeFile(w, req, "../html/artisttemplate.html")
}
