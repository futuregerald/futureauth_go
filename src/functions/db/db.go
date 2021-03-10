package db

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO create method on user to validate password

func Connect(uri string) error {
	return mgm.SetDefaultConfig(nil, "local_dev", options.Client().ApplyURI(uri))
}

func New() (*Client, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI != "" {
		return &Client{}, Connect(mongoURI)
	}
	log.Print("No mongoDB URI provided. Api starting without a data store")
	return &Client{}, nil
}

func (client *Client) Save(m mgm.Model) error {
	return mgm.Coll(m).Create(m)
}
