package ConcertAPI

import (
	"net/http"
)

func VidServe(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../html/static/vid.mp4")
}