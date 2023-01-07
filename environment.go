package simplifi

import (
	"context"
	"crypto/tls"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/richterk/simplifi/properties"
)

// Environment - Type representing the Simplifi environment
type Environment struct {
	DB         *mongo.Session
	Properties *properties.Properties
}

// IgnitionEngineEnvironment - Type representing a specific Simplifi environment implementation
type IgnitionEngineEnvironment struct {
	DB         *mongo.Session
	Properties *properties.IgnitionEngineProperties
}

// InitDB - DB Stuff
func (env *IgnitionEngineEnvironment) InitDB(props *properties.IgnitionEngineProperties) {
	appProperties := props
	//		session, err := mgo.Dial(appProperties.DB.Host)
	// NEW DB LOGIC
	uri := strings.Replace(appProperties.DB.Host, "?ssl=true&", "?", 1)
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	// END NEW DB LOGIC
	session, err := client.StartSession()

	if err != nil {
		panic(err)
	}
	env.DB = &session
}

// InitDB - DB Stuff
func (env *Environment) InitDB(props *properties.Properties) {
	appProperties := props
	//		session, err := mgo.Dial(appProperties.DB.Host)
	// NEW DB LOGIC
	uri := strings.Replace(appProperties.DB.Host, "?ssl=true&", "?", 1)
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	// END NEW DB LOGIC
	session, err := client.StartSession()

	if err != nil {
		panic(err)
	}
	env.DB = &session
}
