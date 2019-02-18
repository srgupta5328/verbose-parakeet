package app

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/srgupta5328/verbose-parakeet/helpers"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Coin Market Application in Go")
}

func LatestListingHandler(w http.ResponseWriter, r *http.Request) {
	url := helpers.BASE_URL + "/cryptocurrency/listings/latest"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-CMC_PRO_API_KEY", helpers.API)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error executing the request: %s", err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}
