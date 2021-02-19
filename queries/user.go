package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/molson82/venus/models"
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
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// id exists
						for _, user := range models.UserMockData() {
							if int(user.ID) == id {
								return user, nil
							}
						}
					}

					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(models.UserType),
				Description: "Get user list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return models.UserMockData(), nil
				},
			},
		},
	},
)
