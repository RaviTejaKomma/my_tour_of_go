package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

func main() {

	var user User

	fmt.Scanf("%s %s %b", &user.Name, &user.Email, &user.IsAdmin)

	var user1 map[string]interface{}
	temp, _ := json.Marshal(&user)
	json.Unmarshal(temp, &user1)

	fmt.Println(temp)
	fmt.Println(user)
	fmt.Println(user1)
}
