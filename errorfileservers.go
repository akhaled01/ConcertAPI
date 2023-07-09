package concertAPI

import "net/http"

func Serve404Html(w http.ResponseWriter, r *http.Request){
	http.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/output.css")
	})
	http.HandleFunc("/sad.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/static/sad.png")
	})
	http.ServeFile(w, r, "html/404.html")
}

func Serve400Html(w http.ResponseWriter, r *http.Request){
	http.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/output.css")
	})
	http.HandleFunc("/sad.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/static/sad.png")
	})
	http.ServeFile(w, r, "html/400.html")
}

func Serve500Html(w http.ResponseWriter, r *http.Request){
	http.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/output.css")
	})
	http.HandleFunc("/sad.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/static/sad.png")
	})
	http.ServeFile(w, r, "html/500.html")
}