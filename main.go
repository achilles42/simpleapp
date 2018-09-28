package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	FailureModeAllow bool `json:"failure_mode_allow"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatRequest(r))
	response := Response{true}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("check to make sure it works")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func handleRequests() {
	http.HandleFunc("/jobs/metadata", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func formatRequest(r *http.Request) string {
	// Create return string
	fmt.Println("***********")
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v%v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	// Return the request as a string
	return strings.Join(request, "\n")
}

func main() {
	handleRequests()
}
