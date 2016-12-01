package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/webercoder/go-mtg-hipchat-bot/lib"
)

func usage(msg string) {
	if len(msg) == 0 {
		fmt.Print(msg)
	}
	fmt.Printf("Usage: %s port\n", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage("Please provide a port for listening to requests.")
		os.Exit(1)
	}

	port := os.Args[1]
	_, err := strconv.Atoi(port)
	if err != nil {
		usage("Please provide a port for listening to requests.")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle the request
		rc := &lib.HipChatRequestController{}
		hcrsp := rc.HandleRequest(r)

		// Convert the HipChatResponse to JSON and send it.
		b, err := json.Marshal(hcrsp)
		if err != nil {
			fmt.Fprintf(w, "Error creating JSON response for %+v: %+v", hcrsp, err)
			return
		}
		fmt.Fprintf(w, "%s", string(b))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
