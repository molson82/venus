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
		"update": &graphql.Field{
			Type:        models.UserType,
			Description: "Update user by id",
			Args: graphql.FieldConfigArgument{
				"id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"firstName": &graphql.ArgumentConfig{Type: graphql.String},
				"lastName":  &graphql.ArgumentConfig{Type: graphql.String},
				"phone":     &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				firstName, fnOk := params.Args["firstName"].(string)
				lastName, lnOk := params.Args["lastName"].(string)
				phone, pOk := params.Args["phone"].(string)
				updatedUser := models.User{}
				for i, u := range queries.MockUserData {
					if int64(id) == u.ID {
						if fnOk {
							queries.MockUserData[i].Firstname = firstName
						}
						if lnOk {
							queries.MockUserData[i].LastName = lastName
						}
						if pOk {
							queries.MockUserData[i].Phone = phone
						}
						updatedUser = queries.MockUserData[i]
						break
					}
				}
				return updatedUser, nil
			},
		},
		"delete": &graphql.Field{
			Type:        models.UserType,
			Description: "Delete user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				deletedUser := models.User{}
				for i, u := range queries.MockUserData {
					if int64(id) == u.ID {
						deletedUser = queries.MockUserData[i]
						queries.MockUserData = append(queries.MockUserData[:i], queries.MockUserData[i+1:]...)
					}
				}
				return deletedUser, nil
			},
		},
	},
})
