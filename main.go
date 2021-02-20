package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/molson82/venus/mutations"
	"github.com/molson82/venus/queries"
)

var userSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queries.UserQueryType,
		Mutation: mutations.UserMutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	// fsClient := config.SetupFirebaseConfig()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), userSchema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)

}
