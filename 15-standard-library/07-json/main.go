package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	var product Product
	if err := json.Unmarshal([]byte(`{"name":"tea","price":12}`), &product); err != nil {
		fmt.Println("invalid JSON:", err)
		return
	}
	fmt.Println("product:", product.Name, product.Price)

	encoded, err := json.Marshal(product)
	if err != nil {
		fmt.Println("cannot encode product:", err)
		return
	}
	fmt.Println("json:", string(encoded))
}
