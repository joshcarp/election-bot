package electionbot

import (
	"context"
	"log"
	"testing"
)

func TestWhatever(t*testing.T){
	ser, err := server()
	if err != nil {
		return
	}
	if err := ser.Respond(context.Background(), ser.Database.Client); err != nil {
		log.Println(err)
	}
}
