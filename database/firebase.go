package database

import (
	"context"


	"cloud.google.com/go/firestore"
)

type Firestore struct {
	*firestore.Client
}

func NewFirestore(projectID string) (Firestore, error) {
	a, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		return Firestore{}, err
	}
	return Firestore{Client: a}, nil
}

func (f Firestore) Get(collection, name string, i interface{}) error {
	snap, err := f.Collection(collection).Doc(name).Get(context.Background())
	if err != nil {
		return err
	}
	return snap.DataTo(i)
}

func (f Firestore) Set(collection, name string, i interface{}) error {
	_, err := f.Collection(collection).Doc(name).Set(context.Background(), i)
	return err
}
