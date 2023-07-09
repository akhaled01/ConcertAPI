package main

import (
	"net/http"
	"ConcertAPI/handlers"
	"fmt"
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
    	ConcertAPI.Index(w, req)
    case "/output.css": // Serves the css file
    	ConcertAPI.StylesServ(w, req)
    case "/static/vid.mp4":
    	ConcertAPI.VidServe(w, req)
    case "/bands.html":
    	ConcertAPI.Bands(w, req)
    case "/artists":
        ConcertAPI.Artists(w, req)

    default:
        http.Error(w, "404 Not Found", 404)
    }
}
