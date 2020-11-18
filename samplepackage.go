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