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
			"street1": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					add := p.Source.(User)
					return add.Address.Street1, nil
				},
			},
			"street2": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					add := p.Source.(User)
					return add.Address.Street2, nil
				},
			},
			"city": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					add := p.Source.(User)
					return add.Address.City, nil
				},
			},
			"state": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					add := p.Source.(User)
					return add.Address.State, nil
				},
			},
			"zipcode": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					add := p.Source.(User)
					return add.Address.ZipCode, nil
				},
			},
			"phone": &graphql.Field{Type: graphql.String},
		},
	},
)
