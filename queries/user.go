package queries

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/molson82/venus/config"
	"github.com/molson82/venus/models"
	"google.golang.org/api/iterator"
)

// UserQueryType : the GraphQL query type for User schema
var UserQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        models.UserType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var user models.User
					id, ok := p.Args["id"].(string)
					if ok {
						// id exists
						dataSnap, err := config.FSClient.DBClient.Collection("user").Doc(id).Get(context.Background())
						if err != nil {
							return nil, err
						}
						dataSnap.DataTo(&user)
					}

					return user, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(models.UserType),
				Description: "Get user list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var users []models.User
					limit, ok := params.Args["limit"].(int)
					if !ok {
						limit = 100
					}
					docIter := config.FSClient.DBClient.Collection("user").Limit(limit).Documents(context.Background())
					for {
						doc, err := docIter.Next()
						if err == iterator.Done {
							break
						}
						if err != nil {
							return nil, err
						}
						var u models.User
						doc.DataTo(&u)
						users = append(users, u)
					}

					return users, nil
				},
			},
		},
	},
)
