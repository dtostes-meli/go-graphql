package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dtostes-meli/go-graphql/schema"
	"github.com/graphql-go/graphql"
)

func init() {
	seller1 := schema.Seller{ID: 1, Name: "Mr Daniels"}
	seller2 := schema.Seller{ID: 2, Name: "Mr Cabral"}
	seller3 := schema.Seller{ID: 3, Name: "Ms Clarissa"}
	seller4 := schema.Seller{ID: 4, Name: "Mr Celsius"}
	schema.SellerList = append(schema.SellerList, seller1, seller2, seller3, seller4)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.SellerSchema)
		json.NewEncoder(w).Encode(result)
	})
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get single seller: curl -g 'http://localhost:8080/graphql?query={todo(id:\"1\"){id,text}}'")
	http.ListenAndServe(":8080", nil)
}
