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
		rc := &lib.RequestController{}
		decoder := json.NewDecoder(r.Body)

		var req lib.Request
		err = decoder.Decode(&req)
		if err != nil {
			fmt.Fprintf(w, "Difficulty parsing HipChat request: %+v", err)
			return
		}

		rsp, err := rc.HandleRequest(&req)
		if err != nil {
			fmt.Fprintf(w, "Error finding your card: %+v", err)
			return
		}

		b, err := json.Marshal(rsp)
		if err != nil {
			fmt.Fprintf(w, "Error creating JSON response for %+v: %+v", rsp, err)
			return
		}

		fmt.Fprintf(w, "%s", string(b))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
