package main

import (
	"fmt"
	"os"
	"strconv"
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const (
	neo4jURI = "bolt://neo4j:neo4jdb@localhost:7687"
)

func dbconnect() driver.Conn {
	conn, err := driver.NewDriver().OpenNeo(neo4jURI)
	if err != nil {
		fmt.Println("error connecting to neo4j:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to neo4j")
	return conn
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}
}

func createStatement(q string, con driver.Conn) driver.Stmt {
	st, err := con.PrepareNeo(q)
	handleError(err)
	return st
}

func createFleets(conn driver.Conn) int64 {

	query := `
	CREATE
		(f : Fleet{name : {name}, url : {url}}) 
	`
	statment := createStatement(query, conn)

	fleetName := "Fleet"
	url := "https://mojoreads.com/readfleet/"

	var numResult int64

	for i := 1; i <= 10; i++ {

		name := fleetName + strconv.Itoa(i)
		result, err := statment.ExecNeo(map[string]interface{}{
			"name": name,
			"url":  url + name,
		})
		handleError(err)

		numRows, err := result.RowsAffected()
		handleError(err)
		numResult += numRows
		fmt.Println(name, url+name)
	}

	// Closing the statment will also close the rows
	statment.Close()

	return numResult
}

func createUsersAndRelationsWithFleets(conn driver.Conn) int64 {

	query := `
	MATCH (f : Fleet) WHERE f.name={fleetname}
	CREATE
		(u : User{name : {name}, emailid : {emailid}, isAdmin : {isAdmin}}),
	(u)-[:MEMBEROF]->(f)
	`
	statement := createStatement(query, conn)

	userName := "User"
	emailID := "@gmail.com"
	isAdmin := 0

	num := 1
	var numResult int64

	for i := 0; i <= 99; i++ {

		isAdmin = 0

		if i!=0 && i%(num*10) == 0{
			num++
		}

		if i%11 == 0{
			isAdmin = 1
		}

		fleetName := "Fleet" + strconv.Itoa(num)
		name := userName + strconv.Itoa(i)

		result, err := statement.ExecNeo(map[string]interface	{}{
			"name":      name,
			"emailid":   name + emailID,
			"fleetname": fleetName,
			"isAdmin" : isAdmin,
		})
		handleError(err)

		numRows, err := result.RowsAffected()
		numResult += numRows
		handleError(err)

		fmt.Println(fleetName, name, name+emailID)
	}

	// Closing the statment will also close the rows
	statement.Close()

	return numResult
}

func createUserAndFleetRelations(conn driver.Conn) int64 {

	query := `
			MATCH (f : Fleet) WHERE f.name={fleetname}
			MATCH (u : User) WHERE u.name={name}
			CREATE (u)-[:MEMBEROF]->(f)
			`
	statement := createStatement(query, conn)

	num := 2
	userName := "User"
	var numResult int64

	for i := 3; i <= 99; i += 3 {

		fleetName := "Fleet" + strconv.Itoa(num)
		name := userName + strconv.Itoa(i)
		result, err := statement.ExecNeo(map[string]interface{}{
			"name":      name,
			"fleetname": fleetName,
		})
		handleError(err)

		numRows, err := result.RowsAffected()
		handleError(err)
		numResult += numRows

		fmt.Println(fleetName, name)

		if i%9 == 0 {
			num %= 10
			num++
		}
	}
	// Closing the statment will also close the rows
	statement.Close()

	return numResult
}

func main() {
	conn := dbconnect()
	defer conn.Close()

	// Creating the Fleet nodes
	numRowsEffected := createFleets(conn)
	fmt.Printf("Number of Fleets created: %d\n", numRowsEffected)

	// Creating the User nodes and relationships between Fleet and User nodes
	numRowsEffected = createUsersAndRelationsWithFleets(conn)
	fmt.Printf("Number of users created: %d\n", numRowsEffected)

	// Creating the User nodes and relationships between Fleet and User nodes
	numRowsEffected = createUserAndFleetRelations(conn)
	fmt.Printf("Number of extra user and fleet relations created: %d\n", numRowsEffected)

}
