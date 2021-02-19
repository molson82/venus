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

// UserMockData : returns mock user data
func UserMockData() []User {
	users := []User{
		{
			ID:        1,
			Firstname: "John",
			LastName:  "Smith",
			Address: Address{
				Street1: "123 Maple Street",
				Street2: "",
				City:    "Chicago",
				State:   "IL",
				ZipCode: 34988,
			},
			Phone: "111-867-3099",
		},
		{
			ID:        2,
			Firstname: "Mike",
			LastName:  "Johnson",
			Address: Address{
				Street1: "8 Cactus Lane",
				Street2: "Apt 44",
				City:    "Austin",
				State:   "TX",
				ZipCode: 44444,
			},
			Phone: "987-354-6687",
		},
		{
			ID:        3,
			Firstname: "Sarah",
			LastName:  "Claire",
			Address: Address{
				Street1: "98 Nova Circle",
				Street2: "Apt 35",
				City:    "New York",
				State:   "NY",
				ZipCode: 23877,
			},
			Phone: "713-999-4453",
		},
		{
			ID:        4,
			Firstname: "Brady",
			LastName:  "Bunch",
			Address: Address{
				Street1: "8615 Broadway",
				Street2: "",
				City:    "Hollywood",
				State:   "CA",
				ZipCode: 90210,
			},
			Phone: "444-908-1234",
		},
	}

	return users
}
