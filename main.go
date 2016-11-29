package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rc := &RequestController{}
		decoder := json.NewDecoder(r.Body)

		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Fprintf(w, "Could not find the requested card.")
			return
		}

		rsp := rc.HandleRequest(&req)
		if rsp != nil {
			fmt.Fprintf(w, "Error loading the requested card.")
		}

		fmt.Fprintf(w, "Found your card: \n%+v", rsp)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
