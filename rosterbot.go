package rosterbot

import (
	"context"
	"github.com/joshcarp/rosterbot/database"
	"log"
	"net/http"
	"os"

	"github.com/joshcarp/rosterbot/roster"
)

func RespondHandler(w http.ResponseWriter, r *http.Request) {
	ser, err := server()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := ser.Respond(context.Background(), ser.Database.Client); err != nil {
		log.Println(err)
	}
}

func Enroll(w http.ResponseWriter, r *http.Request) {
	ser, err := server()
	message, err := ser.Enroll(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		log.Println(err)
		return
	}
	w.Write([]byte(message))
}

func server() (roster.Server, error) {
	fire, err := database.NewFirestore(os.Getenv("PROJECT_ID"))
	if err != nil {
		return roster.Server{}, err
	}
	return roster.NewServer(
		os.Getenv("SLACK_CLIENT_ID"),
		os.Getenv("SLACK_CLIENT_SECRET"),
		fire,
		http.DefaultClient,
	), nil
}
