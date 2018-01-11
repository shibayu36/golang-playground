package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	type Human struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Sex string `json:"sex"`
	}

	// data := `{"first_name": "Yuki", "last_name": "Shibazaki", "sex": "male"}`
	data := `{"first_name": "Yuki", "last_name": "Shibazaki", "sex": "male"}`

	human := &Human{}
	json.Unmarshal([]byte(data), human)
	fmt.Printf("%v\n", human)
}
