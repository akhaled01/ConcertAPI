package main

import (
	ConcertAPI_handlers "ConcertAPI/handlers"
	ConcertAPI_fileServers "ConcertAPI/handlers/file_servers"
	"fmt"
	"net/http"
)

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	switch req.URL.Path {
	case "/":
		ConcertAPI_fileServers.Index(w, req)
	case "/output.css": // Serves the css file
		ConcertAPI_fileServers.StylesServ(w, req)
	case "/static/vid.mp4":
		ConcertAPI_fileServers.VidServe(w, req)
	case "/bands.html":
		ConcertAPI_fileServers.Bands(w, req)
	case "/artists":
		ConcertAPI_handlers.Artists(w, req)
	case "/scripts.js":
		ConcertAPI_fileServers.ScriptsServ(w, req)
	case "/bandinfo":
		ConcertAPI_handlers.BandInfo(w, req)
	case "/artisttemplate.html":
		ConcertAPI_fileServers.ArtistTemplate(w, req)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}
