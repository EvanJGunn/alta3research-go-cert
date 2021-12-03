/*
  Author: Evan Gunn, as part of an Optum Tech University provided class.
*/

package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

// An untyped constant for the API URL, not that it has to be untyped
const SWAPI = "https://swapi.dev/api/"

func main() {
    // A slice of endpoints for the api to query
    var endpoints = []string{"starships/15","planets/55","people/43"}

    // Print all endpoints in formatted string
    for _, endpoint := range endpoints {
        fmt.Printf("Endpoint to be queried: %s%s\n",SWAPI,endpoint)
    }

    // The returned JSON will be in a varied format, make slice of map with len(endpoints)
    responses := make([]map[string](interface{}),len(endpoints))

    // Print out the slice of empty maps
    fmt.Println(responses)

    fmt.Println("The next part takes a second, please be patient!")
    // Make the HTTP requests
    for i, endpoint := range endpoints {
	// Do the GET request
        resp, err := http.Get(fmt.Sprintf("%s%s",SWAPI,endpoint))

	if err != nil {
	    fmt.Println("Error with endpoint:",endpoint,", the error was: ", err)
	} else {
	    // Read in the response body
	    body, respErr := ioutil.ReadAll(resp.Body)

	    if respErr != nil {
	        fmt.Println("Error with a response body")
	        return
	    }

	    // Unmarshalling the body into our map
	    err1 := json.Unmarshal(body, &responses[i])
	    if err1 != nil {
	        fmt.Println("Error unmarshalling data")
	    }
	}
    }

    // Print out all the responses
    for _, response := range responses {
	fmt.Println("\n")
        fmt.Println(response)
    }

    // Print out a thank you to Zach
    fmt.Println("\n ***Thanks Zach!***")
}
