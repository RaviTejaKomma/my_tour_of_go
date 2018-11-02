package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)


const (
	neo4jURI = "bolt://neo4j:neo4jdb@localhost:7687"
)

var db driver.Conn

type book struct {
	ISBN13      string `json:"ISBN13"`
	ISBNType    string `json:"ISBNType"`
	Title       string `json:"title"`
	Authors     string `json:"authors"`
}


type editions struct {
	Books []book `json:"editions"`
	Time     int    `json:"time"`
	Total    int    `json:"total"`
}

func dbconnect() driver.Conn {
	con, err := driver.NewDriver().OpenNeo(neo4jURI)
	if err != nil {
		fmt.Println("error connecting to neo4j:",err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to neo4j")
	return con
}

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.mojoreads.com/work/v2/edition/search?query=java&category=59&limit=5&skip=0", nil)
	handleError(err)
	req.Header.Add("Authorization", "Bearer pERWDV4spn9I1XibuikjiG08qIOiXgmbNUlBRH18d5GSM4KMUHwSH7gQ_p-6jGxC")

	resp, err := client.Do(req)
	handleError(err)

	defer resp.Body.Close()

	ed := new(editions)

	json.NewDecoder(resp.Body).Decode(ed)

	conn := dbconnect()
	defer conn.Close()


	//query := `MERGE (a:Book{ISBN13:{isbn}, Title: {title}}) 
	//		 	ON CREATE SET a.ISBN13 = {isbn}`

	query :=`MATCH (b:Book) WHERE b.Title="Java" RETURN b`

	st := createStatement(query, conn)

	// for i:=0 ; i < len(ed.Books); i++{
	// 	numResult := executeStatement(ed.Books[i], st)	
	// 	fmt.Printf("CREATED ROWS: %d\n", numResult) 
	// }

	rows := queryStatement(st)
	consumeRows(rows,st)
	// Closing the statment will also close the rows
	st.Close()
}

func createStatement(q string, con driver.Conn) driver.Stmt {
	st, err := con.PrepareNeo(q)
	handleError(err)
	return st
}

func queryStatement(st driver.Stmt) driver.Rows {
	// Even once I get the rows, if I do not consume them and close the
	// rows, Neo will discard and not send the data
	rows, err := st.QueryNeo(nil)
	handleError(err)
	return rows
}

func executeStatement(data book, st driver.Stmt) int64{

	result, err := conn.ExecNeo(st,map[string]interface{}{"isbn": data.ISBN13, "title": data.Title})
	handleError(err)

	numResult, err := result.RowsAffected()
	handleError(err)
	return numResult
}

func consumeRows(rows driver.Rows, st driver.Stmt) {
	// This interface allows you to consume rows one-by-one, as they
	// come off the bolt stream. This is more efficient especially
	// if you're only looking for a particular row/set of rows, as
	// you don't need to load up the entire dataset into memory
	data, _, err := rows.NextNeo()
	handleError(err)

	// This query only returns 1 row, so once it's done, it will return
	// the metadata associated with the query completion, along with
	// io.EOF as the error
	_, _, err = rows.NextNeo()
	handleError(err)
	fmt.Printf("FIELDS: %d %f\n", data)
}


func handleError(err error){
	if err != nil{
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}	
}