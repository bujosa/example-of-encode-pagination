package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Filter struct {
	Field    string      `json:"Field"`
	Operator string      `json:"Operator"`
	Value    interface{} `json:"Value"`
}

type Pagination struct {
	Offset  int      `json:"Offset"`
	Limit   int      `json:"Limit"`
	Search  string   `json:"Search"`
	Sort    string   `json:"Sort"`
	Order   string   `json:"Order"`
	Next    string   `json:"Next"`
	Prev    string   `json:"Prev"`
	Filters []Filter `json:"Filters"`
}

func main() {
	// p := types.Pagination{Offset: 10, Limit: 20, Search: "test", Sort: "name", Order: "asc", Prev: "", Next: ""}
	data, _ := json.Marshal(Pagination{
		Offset: 0,
		Limit:  10,
		Search: "John",
		Sort:   "age",
		Order:  "asc",
		Next:   "",
		Prev:   "",
		Filters: []Filter{
			{
				Field:    "name",
				Operator: "eq",
				Value:    "John",
			},
			{
				Field:    "age",
				Operator: "in",
				Value:    []int{18, 19, 20},
			},
		},
	})
	result := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Encoded:", result)

	// Decode data from URL-safe Base64
	decodedData, _ := base64.URLEncoding.DecodeString(result)
	var p2 Pagination
	json.Unmarshal(decodedData, &p2)
	fmt.Println("Decoded:", p2)

	// Output:
	fmt.Print("Decoded:", p2.Limit, p2.Offset, p2.Search, p2.Sort, p2.Order, p2.Prev, p2.Next)
}
