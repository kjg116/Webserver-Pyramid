package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/kjg116/webserver-pyramid/internal/models"
)

// AreTheyPyramids ..
func AreTheyPyramids(w http.ResponseWriter, r *http.Request) {
	var uReq models.Request
	showDetails := getBoolQueryParam(r.URL.Query().Get("show_details"))

	// Serialize the Request into a Request Struct
	err := json.NewDecoder(r.Body).Decode(&uReq)
	if err != nil {
		// Respond with Error is the Request Cannot be Serialized
		WriteErrorResponse(w, http.StatusBadRequest, "invalid_request", err.Error())
		return
	}

	words := uReq.Words
	var wg sync.WaitGroup
	var response models.Response

	// Interate through the given words
	for _, word := range words {
		wg.Add(1)
		go func(w string, showDetails bool, wg *sync.WaitGroup) {
			defer wg.Done()

			var result models.Result
			var charSlice models.CharacterMapSlice
			result.Word = w
			result.IsPyramid, charSlice = isThisAPyramid(w)
			if showDetails {
				result.CharacterMapSlice = charSlice
			}
			response.Results = append(response.Results, result)
		}(word, showDetails, &wg)
	}
	wg.Wait()

	WriteSuccessResponse(w, response)
	return
}

func isThisAPyramid(word string) (bool, models.CharacterMapSlice) {
	// Create a Slice of Key/Value Pairs
	// Keys are the individual characters
	// Values are the number of times the character is seen
	charaterMapSlice := models.NewCharacterMapSlice(word)

	// Assumption: A pyramid cannot have less than three characters
	if len(word) < 3 {
		return false, charaterMapSlice
	}

	// Sorts the Slice by the Value
	charaterMapSlice.SortByValue()

	// Iterate through the slice of Key/Value Pairs
	// If the difference between the indexed character and the character after is greater or less than one then it is not a pyramid
	for i := range charaterMapSlice {
		if i == (len(charaterMapSlice) - 1) {
			continue
		}
		if !isThisPlusOne(charaterMapSlice[i+1].Value, charaterMapSlice[i].Value) {
			return false, charaterMapSlice
		}
	}

	return true, charaterMapSlice
}

func isThisPlusOne(x, y int) bool {
	if x-y == 1 {
		return true
	}
	return false
}

func getBoolQueryParam(param string) bool {
	if param == "" {
		return false
	}
	b, _ := strconv.ParseBool(param)
	return b
}
