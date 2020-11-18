package samplepackage

import (
	"fmt"
)

type Connection struct {
	Host	string
	Port	string
}

func (c *Connection) SetConnection() {
	fmt.Println("yoooo")
}

func New() *Connection {
	fmt.Println("ntaaap")
	return &Connection {
		Host: "BIBI"
		Port: "BUBU"
	}
}