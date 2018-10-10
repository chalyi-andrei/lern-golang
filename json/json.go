package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	ID   int      `json:"id"`
	Data []string `json:"data"`
}

func main() {
	instance := &response{
		ID:   1,
		Data: []string{"apple", "peach", "pear"},
	}
	instanceJSON, _ := json.Marshal(instance)
	fmt.Println(string(instanceJSON))

	// Unmarshal JSON
	str := `{"id": 1, "data": ["apple", "peach"]}`
	res := response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
}
