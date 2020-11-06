package main

import (
	"net/http"

	"github.com/joshcarp/electionbot"
)

func main() {
	http.HandleFunc("/", electionbot.RespondHandler)
	http.ListenAndServe(":8081", nil)
}
