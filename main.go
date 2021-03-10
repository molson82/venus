package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/molson82/venus/mutations"
	"github.com/molson82/venus/queries"
)

var userSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queries.UserQueryType,
		Mutation: mutations.UserMutationType,
	},
)

func main() {
	userHandler := handler.New(&handler.Config{
		Schema:   &userSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/user", userHandler)

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)

}
