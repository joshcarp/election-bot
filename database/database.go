package database

type Database interface {
	Get(collection, name string, i interface{}) error
	Set(collection, name string, i interface{}) error
	Delete(collection, name string) error
}
