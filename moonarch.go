// Package moonarch provides type definitions for moonarch.app's api,
// as well as providing a helper function to fetch API data.
package moonarch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MoonarchResponse represents a moonarch API response.
type MoonarchResponse struct {
	Network string `json:"network"`
	Token   `json:"token"`
	Locks   []Lock `json:"locks"`
}

// APIURL is the base URL that all moonarch API requests
// originate from.
const APIURL = "https://api.moonarch.app/"

// Fetch returns a `MoonarchResponse` for the given address.
func Fetch(address string) (r MoonarchResponse, err error) {
	url := fmt.Sprintf("%v/1.0/tokens/BSC/details/%v", APIURL, address)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	resJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(resJSON, &r)
	return
}
