package models

import "github.com/graphql-go/graphql"

// User struct that holds an Address as well
type User struct {
	ID        int64   `json:"id"`
	Firstname string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Address   Address `json:"address"`
	Phone     string  `json:"phone"`
}

// UserType : User GraphQL Schema
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.Int},
			"firstName": &graphql.Field{Type: graphql.String},
			"lastName":  &graphql.Field{Type: graphql.String},
			"address":   &graphql.Field{Type: AddressType},
			"phone":     &graphql.Field{Type: graphql.String},
		},
	},
)
