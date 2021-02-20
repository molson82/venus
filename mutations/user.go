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
				"phone":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
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
