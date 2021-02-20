package models

import "github.com/graphql-go/graphql"

// Address struct that holds info for User
type Address struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int64  `json:"zipcode"`
}

// AddressType :  GraphQL Schema
var AddressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"street1": &graphql.Field{Type: graphql.String},
			"street2": &graphql.Field{Type: graphql.String},
			"city":    &graphql.Field{Type: graphql.String},
			"state":   &graphql.Field{Type: graphql.String},
			"zipcode": &graphql.Field{Type: graphql.Int},
		},
	},
)
