package framework

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

type Environment struct {
	DB         *mgo.Session
	Properties *Properties
}

// InitDB - DB Stuff
func (env *Environment) InitDB(props *Properties) {
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
