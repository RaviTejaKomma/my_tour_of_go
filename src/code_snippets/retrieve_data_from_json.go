package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// type Cover struct {
// 	Key string `json:"key"`
// }

// // // Structre that contains information regarding Publisher
// type Publisher struct {
// 	Name string `json:"name"`
// }

// // Structre that contains information regarding Category of the Book
// type Categories struct {
// 	Description string `json:"description"`
// 	Name		string `json:"name"`
// }

// type Book struct {
// 	Title   string   `json:"title"`
// 	Authors string `json:"authors"`
// 	Category []Categories `json:"categories"`
// 	Type      string `json:"type"`
// 	Pages     int    `json:"pages"`
// 	Publisher Publisher `json:"publisher"`
// 	Cover Cover `json:"cover"`
// }

// type Books struct{
// 	Book Book `json:"_source"`
// }

// type Hit struct {
// 	Hits []Books `json:"hits"`
// }

// type JsonData struct {
// 	Hits Hit `json:"hits"`
// }

type JsonData struct {
	Hits     struct {
		Hits     []struct {
			Book struct {
				Categories []struct {
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"categories"`
				Title     string `json:"title"`
				Type      string `json:"type"`
				Pages     int    `json:"pages"`
				Authors   string `json:"authors"`
				Publisher struct {
					Name string `json:"name"`
				} `json:"publisher"`
				Cover struct {
					Key string `json:"key"`
				} `json:"cover"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://es.data.staging.mojoreads.com:9200/work-eng/_search?q=publisher.id:915953&size=100&from=200", nil)
	handleError(err)

	resp, err := client.Do(req)
	handleError(err)

	defer resp.Body.Close()

	books := JsonData{}

	body, _ := ioutil.ReadAll(resp.Body)

	bs := string(body)

	err = json.Unmarshal([]byte(bs), &books)
	handleError(err)

	for _, book := range books.Hits.Hits{
		fmt.Printf("Title: %s\nAuthor: %s\nPublisher: %s\nCover: %s\n-------------------------------------------\n",
		book.Book.Title, book.Book.Authors,book.Book.Publisher.Name, book.Book.Cover.Key)
	}

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}
}