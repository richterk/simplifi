package simplifi

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"github.com/richterk/simplifi/properties"
)

//Environment - Type representing the Simplifi environment
type Environment struct {
	DB         *mgo.Session
	Properties *properties.Properties
}

//IgnitionEngineEnvironment - Type representing a specific Simplifi environment implementation
type IgnitionEngineEnvironment struct {
	DB         *mgo.Session
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

	dialInfo, err := mgo.ParseURL(uri)

	if err != nil {
		panic(err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("Failed to connect: ", err)
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	// END NEW DB LOGIC

	env.DB = session
}

// InitDB - DB Stuff
func (env *Environment) InitDB(props *properties.Properties) {
	appProperties := props
	//		session, err := mgo.Dial(appProperties.DB.Host)
	// NEW DB LOGIC
	uri := strings.Replace(appProperties.DB.Host, "?ssl=true&", "?", 1)
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInfo, err := mgo.ParseURL(uri)

	if err != nil {
		panic(err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("Failed to connect: ", err)
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	// END NEW DB LOGIC

	env.DB = session
}
