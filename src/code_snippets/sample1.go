/*
This is to demonstrate how to connect to the Neo4j Database using Bolt driver from Go.
*/
package main

import (
	"log"
	"fmt"
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const (
	neo4jURI = "bolt://neo4j:neo4jdb@localhost:7687"
)

var db driver.Conn

func main(){
	var err error
	db, err = driver.NewDriver().OpenNeo(neo4jURI)
	fmt.Printf("%T \n",db)
	if err!= nil{
		log.Println("error connecting to neo4j:", err)
		return
	}
	defer db.Close()
	log.Println("Successfully connected to neo4j")
}