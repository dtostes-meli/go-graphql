package schema

import (
	"github.com/graphql-go/graphql"
)

// SellerList List
var SellerList []Seller

// Seller Type
type Seller struct {
	ID   int64  `"json:id"`
	Name string `"json:name"`
}

var sellerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Seller",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"sellerById": &graphql.Field{
			Type:        sellerType,
			Description: "List of Sellers",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				id64 := int64(id.(int))
				for _, seller := range SellerList {
					found := seller.ID == id64
					if found {
						return seller, nil
					}
				}
				return Seller{}, nil
			},
		},
		"sellerList": &graphql.Field{
			Type:        graphql.NewList(sellerType),
			Description: "List of Sellers",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return SellerList, nil
			},
		},
	},
})

// SellerSchema is the GraphQL schema for seller
var SellerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
