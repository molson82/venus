package main

import (
	"fmt"
	"net/http"
	"os"

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

	fmt.Printf("Server is running on port %q", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
