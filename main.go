package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/educolog7/packages/types"
)

func main() {
	p := types.Pagination{Offset: 10, Limit: 20, Search: "test", Sort: "name", Order: "asc", Prev: "", Next: ""}
	data, _ := json.Marshal(p)
	result := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Encoded:", result)

	// Decode data from URL-safe Base64
	decodedData, _ := base64.URLEncoding.DecodeString(result)
	var p2 types.Pagination
	json.Unmarshal(decodedData, &p2)
	fmt.Println("Decoded:", p2)

	// Output:
	fmt.Print("Decoded:", p2.Limit, p2.Offset, p2.Search, p2.Sort, p2.Order, p2.Prev, p2.Next)
}
