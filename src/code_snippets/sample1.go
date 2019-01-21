/*
This is to demonstrate how to connect to the Neo4j Database using Bolt driver from Go.
*/
package main

import (
	"fmt"
	"os"

	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

const (
	neo4jURI = "bolt://neo4j:neo4jdb@localhost:7687"
	//neo4jURI = "bolt://neo4j:xaDEiYB4grt7@35.207.184.186:7687"
)

type User struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db driver.Conn

func dbconnect() driver.Conn {
	con, err := driver.NewDriver().OpenNeo(neo4jURI)
	if err != nil {
		fmt.Println("error connecting to neo4j:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to neo4j")
	return con
}

func main() {

	var user User
	user.Name = "abc"
	user.Email = "abc@gmail.com"
	user.UID = "123456"

	conn := dbconnect()
	defer conn.Close()

	var userCreated User

	// CreateUserQuery := `
	// CREATE
	// 	(u : User{ uid : {uid}, name : {name}, email : {email} })
	// RETURN
	// 	u.uid as userID, u.name as name, u.email as email
	// `

	// CreateUserQuery := `
	// CREATE
	// 	(u : User{ uid : {uid}, name : {name}, email : {email} })
	// RETURN
	// 	u
	// `

	CreateUserQuery := `
	CREATE
		(u : User{ uid, name, email} )
	RETURN+
		u
	`

	stmt, err := conn.PrepareNeo(CreateUserQuery)
	defer stmt.Close()
	if err != nil {
		fmt.Println("Error occurred in CreateUser:UserService is:", err)
		return
	}

	rows, err := stmt.QueryNeo(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"uid":   user.UID,
	})
	if err != nil {
		fmt.Println("Error occurred in CreateUser:UserService is:", err)
		return
	}

	fmt.Println("Rows is:", rows)

	data, temp, err := rows.NextNeo()
	if err != nil {
		fmt.Println("Error occurred in CreateUser:UserService is:", err)
		return
	}

	// fmt.Println("Temp is:", temp)
	fmt.Println("Data is:", data)
	// fmt.Println("Data is:", data[0].)

	// fmt.Println("Data is:", data[0].email, data[0].name, data[0].uid)

	userCreated.UID = data[0].(string)
	userCreated.Name = data[1].(string)
	userCreated.Email = data[2].(string)

	fmt.Println("User created is:", userCreated)
}

/* user := datamodels.User{
	UUID : "7680f153-a716-49e6-bab1-433abd4f680d",
	Name : "sai",
	Email : "sai@gmail.com",
}

fleet := []datamodels.Fleet{}

db.Model(&user).Association("Fleets").Find(&fleet)

fmt.Println(fleet)

var user datamodels.User

db.Raw("select * from users where uuid = ?", "7680f153-a716-49e6-bab1-433abd4f680d").Scan(&user)

var t []FleetInfo

query := `
	SELECT user_fleets.is_admin, user_fleets.status, fleets."name", fleets.uuid,  fleets.url FROM fleets
	INNER JOIN user_fleets
	ON fleets.uuid = user_fleets.fleet_uuid where user_fleets.user_uuid = '7680f153-a716-49e6-bab1-433abd4f680d'
`
db.Raw(query).Scan(&t)

fmt.Println(user)
fmt.Println(t)

var fleets []datamodels.Fleet
db.Model(&user).Related(&fleets, "Fleets")
fmt.Println("Fleets is:", fleets)

user := datamodels.User{
	UUID : "123456",
	Name : "ravi",
	Email : "ravi@gmail.com",
}

db.Create(&user)

fleet := datamodels.Fleet{
	UUID : "73046406",
	Name : "mojoreads",
	URL : "logosvision",
	CreatedBy : user.UUID,
}

fleet.Users = make([]datamodels.User, 1)
fleet.Users[0] = user

db.Create(&fleet)

var userRetrieved datamodels.User

db.Table("users").Where("uuid = ?", "123456").First(&userRetrieved)
fmt.Println(userRetrieved)
db.Model(&userRetrieved).Related(&userRetrieved.Fleets, "Fleets")

fmt.Println(userRetrieved)

var temp []datamodels.Fleet;

db.Set("gorm:auto_preload", true).Find(&temp)

fmt.Println(temp)

fleets := []datamodels.Fleet{}

db.Model(&user).Association("Fleets").Find(&fleets)

fmt.Println(fleets)
*/
