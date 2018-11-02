package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Author struct{
	AuthorID string		
	AuthorName string	
}

type Book struct{
	BookID string   
	Title string	
	Authors []Author
}

type Books struct{
	Books []Book
}

const (
	neo4jURI = "bolt://neo4j:neo4jdb@localhost:7687"
)

var db driver.Conn

func dbconnect() driver.Conn {
	con, err := driver.NewDriver().OpenNeo(neo4jURI)
	if err != nil {
		fmt.Println("error connecting to neo4j:",err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to neo4j")
	return con
}


func main(){

	jsonFile, err := os.Open("books_data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened books_Data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	var booksData Books

	json.Unmarshal([]byte(byteValue), &booksData)

	query := `MERGE (b:Book{bookId:{BookID})`

	conn := dbconnect()
	defer conn.Close()

	st := createStatement(query, conn)

	for i:=0 ; i<len(booksData.Books); i++{
		fmt.Println(booksData.Books[i].BookID)
	}

	st.Close()
}

func createStatement(q string, con driver.Conn) driver.Stmt {
	st, err := con.PrepareNeo(q)
	handleError(err)
	return st
}

func handleError(err error){
	if err != nil{
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}	
}