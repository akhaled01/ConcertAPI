package ConcertAPI

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func Artists(w http.ResponseWriter, req *http.Request) {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(body))
}
