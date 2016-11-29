package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/webercoder/go-mtg-service/lib"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rc := &lib.RequestController{}
		decoder := json.NewDecoder(r.Body)

		var req lib.Request
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Fprintf(w, "Could not find the requested card.")
			return
		}

		rsp := rc.HandleRequest(&req)

		fmt.Fprintf(w, "Found your card: \n%+v", rsp)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
