package db

import (
	"encoding/json"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PasswordConfig is used to generate the argon2 password hash
type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Email            string             `json:"name" bson:"name"`
	Tenant           primitive.ObjectID `json:"tenantID" bson:"tenantID,omitempty"`
	password         string             `bson:"password"`
	AppMetaData      bson.M             `json:"appMetaData" bson:"appMetaData"`
	UserMetaData     bson.M             `json:"userMetaData" bson:"appMetaData"`
	Confirmed        bool               `json:"confirmed" bson:"confirmed"`
	IsAdmin          bool               `json:"isAdmin" bson:"isAdmin"`
	Disabled         bool               `json:"disabled" bson:"disabled"`
	Roles            []string           `json:"roles" bson:"roles"`
}

// UserJSON is what's sent to the signup endpoint and used to create the User model
type UserJSON struct {
	ID           string             `json:"id,omitempty"`
	Email        string             `json:"name" bson:"name"`
	Tenant       primitive.ObjectID `json:"tenantID" bson:"tenantID,omitempty"`
	Password     string             `json:"password" bson:"password"`
	AppMetaData  json.RawMessage    `json:"appMetaData" bson:"appMetaData"`
	UserMetaData json.RawMessage    `json:"userMetaData" bson:"appMetaData"`
	Confirmed    bool               `json:"confirmed" bson:"confirmed"`
	IsAdmin      bool               `json:"isAdmin" bson:"isAdmin"`
	Disabled     bool               `json:"disabled" bson:"disabled"`
	Roles        []string           `json:"roles" bson:"roles"`
}
