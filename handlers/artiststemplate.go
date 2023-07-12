package ConcertAPI

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	id           int      `json:"id"`
	image        string   `json:"image"`
	members      []string `json:"members"`
	creationData int      `json:"creationData"`
	firstAlbum   string   `json:"firstAlbum"`
	locations    string   `json:"locations"`
	relations    string   `json:"relations"`
}

func ArtistTemplate(w http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	http.ServeFile(w, req, "../html/artisttemplate.html")
}
