package ConcertAPI

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	http.ServeFile(w, req, "../html/index.html")
	w.Header().Set("Content-Type", "text/html")
}
