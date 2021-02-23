package models

import "github.com/graphql-go/graphql"

// User struct that holds an Address as well
type User struct {
	ID        string `firestore:"id,omitempty" json:"id"`
	Firstname string `firestore:"firstName,omitempty" json:"firstName"`
	LastName  string `firestore:"lastName,omitempty" json:"lastName"`
	Phone     string `firestore:"phone,omitempty" json:"phone"`
}

// UserType : User GraphQL Schema
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.String},
			"firstName": &graphql.Field{Type: graphql.String},
			"lastName":  &graphql.Field{Type: graphql.String},
			"phone":     &graphql.Field{Type: graphql.String},
		},
	},
)
