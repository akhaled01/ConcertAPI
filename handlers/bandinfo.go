package ConcertAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Band struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
}

func BandInfo(w http.ResponseWriter, req *http.Request) {
	var band Band

	request_body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the received body to integer
	band_id, _ := strconv.Atoi(string(request_body))

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(band_id))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	// Ready the body into bytes
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &band)
	if err != nil {
		log.Fatal(err)
	}
	origin := req.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	json_band_data, err := json.Marshal(band)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json_band_data))
}
