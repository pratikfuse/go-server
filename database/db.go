package database

import (
	"os"

	"gopkg.in/mgo.v2"
)

var Conn *Connection

type Connection struct {
	Db *mgo.Database
}

func SetDbConnection(connection *Connection) {
	Conn = connection
}

func (c *Connection) Connect() error {
	config := mgo.DialInfo{
		Addrs:    []string{os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")},
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	server, err := mgo.DialWithInfo(&config)

	if err != nil {
		return err
	}

	c.Db = server.DB(config.Database)
	return nil
}
