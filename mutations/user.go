package mutations

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/molson82/venus/config"
	"github.com/molson82/venus/models"
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
				newUser := models.User{
					FirstName: params.Args["firstName"].(string),
					LastName:  params.Args["lastName"].(string),
					Phone:     params.Args["phone"].(string),
				}
				data, _, err := config.FSClient.DBClient.Collection("user").Add(context.Background(), newUser)
				if err != nil {
					return nil, err
				}
				newUser.ID = data.ID

				return newUser, nil
			},
		},
		"update": &graphql.Field{
			Type:        models.UserType,
			Description: "Update user by id",
			Args: graphql.FieldConfigArgument{
				"id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"firstName": &graphql.ArgumentConfig{Type: graphql.String},
				"lastName":  &graphql.ArgumentConfig{Type: graphql.String},
				"phone":     &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(string)
				updatedUser := createUpdateList(params, models.User{})
				_, err := config.FSClient.DBClient.Collection("user").Doc(id).Update(context.Background(), updatedUser)
				if err != nil {
					return nil, err
				}
				dataSnap, geterr := config.FSClient.DBClient.Collection("user").Doc(id).Get(context.Background())
				if geterr != nil {
					return nil, geterr
				}
				var user models.User
				dataSnap.DataTo(&user)
				user.ID = id

				return user, nil
			},
		},
		"delete": &graphql.Field{
			Type:        models.UserType,
			Description: "Delete user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// id, _ := params.Args["id"].(string)
				// deletedUser := models.User{}
				// for i, u := range queries.MockUserData {
				// 	if id == u.ID {
				// 		deletedUser = queries.MockUserData[i]
				// 		queries.MockUserData = append(queries.MockUserData[:i], queries.MockUserData[i+1:]...)
				// 	}
				// }
				return nil, nil
			},
		},
	},
})
