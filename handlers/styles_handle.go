package ConcertAPI

import (
	"fmt"
	"net/http"
)

func StylesServ(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/output.css")
}