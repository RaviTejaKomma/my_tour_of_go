package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

type Cover struct {
	Key string `json:"key"`
}

// // Structre that contains information regarding Publisher
type Publisher struct {
	Name string `json:"name"`
}

// Structre that contains information regarding Category of the Book
type Categories struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Book struct {
	Title     string       `json:"title"`
	ISBN13    string       `json:"ISBN13"`
	Authors   string       `json:"authors"`
	Category  []Categories `json:"categories,omitempty"`
	Type      string       `json:"type"`
	Pages     int          `json:"pages"`
	Publisher Publisher    `json:"publisher"`
	Cover     Cover        `json:"cover"`
}

type Books struct {
	Book Book `json:"_source"`
}

type Hit struct {
	Hits []Books `json:"hits"`
}

type JsonData struct {
	Hits Hit `json:"hits"`
}

// type JsonData struct {
// 	Hits     struct {
// 		Hits     []struct {
// 			Book struct {
// 				Categories []struct {
// 					Description string `json:"description"`
// 					Name        string `json:"name"`
// 				} `json:"categories"`
// 				Title     string `json:"title"`
// 				ISBN13    string  `json:"ISBN13"`
// 				Type      string `json:"type"`
// 				Pages     int    `json:"pages"`
// 				Authors   string `json:"authors"`
// 				Publisher struct {
// 					Name string `json:"name"`
// 				} `json:"publisher"`
// 				Cover struct {
// 					Key string `json:"key"`
// 				} `json:"cover"`
// 			} `json:"_source"`
// 		} `json:"hits"`
// 	} `json:"hits"`
// }

const (
	neo4jURI = "bolt://neo4j:neo4jdb@localhost:7687"
)

func dbconnect() driver.Conn {
	con, err := driver.NewDriver().OpenNeo(neo4jURI)
	if err != nil {
		fmt.Println("error connecting to neo4j:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to neo4j")
	return con
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}
}

func executeStatement(book Books, st driver.Stmt) int64 {

	result, err := st.ExecNeo(map[string]interface{}{
		"title":     book.Book.Title,
		"publisher": book.Book.Publisher.Name,
		"cover":     book.Book.Cover.Key,
		"category":  book.Book.Category[0].Name,
		"isbn13":    book.Book.ISBN13,
		"authors":   book.Book.Authors,
	})
	handleError(err)

	numResult, err := result.RowsAffected()
	handleError(err)

	return numResult
}

func createStatement(q string, con driver.Conn) driver.Stmt {
	st, err := con.PrepareNeo(q)
	handleError(err)
	return st
}

func insertDataInNeo(conn driver.Conn, books JsonData) int64 {
	query := `
	MERGE 
		(b : Book{authors : {authors}, title : {title}, publisher : {publisher}, cover : {cover}, category : {category}}) 
	ON CREATE SET 
		b.ISBN13 = {isbn13}
	`

	st := createStatement(query, conn)
	var numResult int64

	for _, book := range books.Hits.Hits {
		numResult += executeStatement(book, st)
	}

	// Closing the statment will also close the rows
	st.Close()

	return numResult
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	client := &http.Client{}

	conn := dbconnect()
	defer conn.Close()

	var numberOfBooks int
	fmt.Println("Enter the number of books to retrieved:")
	fmt.Scanf("%d\n", &numberOfBooks)

	URL := "http://es.data.staging.mojoreads.com:8881/work-eng/_search?q=publisher.id:915953&size=" + strconv.Itoa(numberOfBooks) + "&from=200"

	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth("logos", "HDxvf1fDRzec4LaKrB72dfozyqJj800u"))

	handleError(err)

	resp, err := client.Do(req)
	handleError(err)

	defer resp.Body.Close()

	books := JsonData{}

	body, _ := ioutil.ReadAll(resp.Body)

	bs := string(body)

	err = json.Unmarshal([]byte(bs), &books)
	handleError(err)

	// for _, book := range books.Hits.Hits {
	// 	fmt.Printf("Title: %s\nAuthor: %s\nPublisher: %s\nCover: %s\nCategory: %s\n-------------------------------------------\n",
	// 		book.Book.Title, book.Book.Authors, book.Book.Publisher.Name, book.Book.Cover.Key, book.Book.Category[0].Name)
	// }

	numRowsEffected := insertDataInNeo(conn, books)
	fmt.Printf("Number of books created: %d\n", numRowsEffected)

}
