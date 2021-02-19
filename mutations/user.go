package mutations

import (
	"log"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/molson82/venus/models"
	"github.com/molson82/venus/queries"
)

// UserMutationType : User Graphql schema mutations. Create Update Delete
var UserMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        models.UserType,
			Description: "Create a new User",
			Args: graphql.FieldConfigArgument{
				"firstName": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"lastName":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"street1":   &graphql.ArgumentConfig{Type: graphql.String},
				"street2":   &graphql.ArgumentConfig{Type: graphql.String},
				"city":      &graphql.ArgumentConfig{Type: graphql.String},
				"state":     &graphql.ArgumentConfig{Type: graphql.String},
				"zipcode":   &graphql.ArgumentConfig{Type: graphql.Int},
				"phone":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				// address := models.Address{
				// 	Street1: params.Args["street1"].(string),
				// 	Street2: params.Args["street2"].(string),
				// 	City:    params.Args["city"].(string),
				// 	State:   params.Args["state"].(string),
				// 	ZipCode: params.Args["zipcode"].(int64),
				// }
				user := models.User{
					ID:        int64(rand.Intn(100000)),
					Firstname: params.Args["firstName"].(string),
					LastName:  params.Args["lastName"].(string),
					Phone:     params.Args["phone"].(string),
				}
				log.Printf("new User logged: %v\n", user)
				queries.MockUserData = append(queries.MockUserData, user)
				return user, nil
			},
		},
	},
})
