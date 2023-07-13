package ConcertAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

/* Received JSON will be organized in this struct */
type Recieved_Band struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
}

/* After we fetch the locations, it will be organized in this struct */
type Recieved_Locations struct {
	Locations []string `json:locations`
}

/* Proccessed data will go in this struct, and it's ready to be sent */
type Sent_Band struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Relations    string
}

func BandInfo(w http.ResponseWriter, req *http.Request) {
	var recieved_band Recieved_Band // Will be saving the recived JSON into this struct

	request_body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	/* Convert the received body to integer */
	band_id, _ := strconv.Atoi(string(request_body))

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(band_id))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	/* Read the body into bytes */
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	/* Now put the recived data into received_band structure */
	err = json.Unmarshal(body, &recieved_band)
	if err != nil {
		log.Fatal(err)
	}

	/* set the headers */
	origin := req.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	/* Construct the "struct to be sent" */
	struct_to_be_sent := construct_sent_band(recieved_band)

	/* Convert the struct to JSON */
	json_band_data, err := json.Marshal(struct_to_be_sent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/* send the JSON to the client */
	fmt.Fprint(w, string(json_band_data))
}

/*
	This function will take the recieved data struct

as input and returns the struct of the proccessed data
*/
func construct_sent_band(received Recieved_Band) Sent_Band {
	var sent_band Sent_Band // A struct that holds the data to be sent

	recieved_location := fetch_Locations(received.Locations)

	/* Modifying the struct to be sent */
	sent_band.Id = received.Id
	sent_band.Name = received.Name
	sent_band.Image = received.Image
	sent_band.Members = received.Members
	sent_band.CreationDate = received.CreationDate
	sent_band.FirstAlbum = received.FirstAlbum
	sent_band.Locations = recieved_location.Locations
	sent_band.Relations = received.Relations

	return sent_band
}

/* fetch locations from the given url and returns the array of locations in a struct */
func fetch_Locations(locations_url string) Recieved_Locations {
	var recieved_locations Recieved_Locations // Will be saving the recived JSON into this struct

	/* fetching the url */
	response, err := http.Get(locations_url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	/* Read the body into bytes */
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	/* Modify the recieved_location struct to have the array of locations */
	err = json.Unmarshal(body, &recieved_locations)
	if err != nil {
		log.Fatal(err)
	}

	return recieved_locations
}

/* This function should do the same as the previous funciton */
func fetch_Relations(locations_url string) {

}
