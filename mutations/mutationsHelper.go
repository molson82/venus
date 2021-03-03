package mutations

import (
	"bytes"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/graphql-go/graphql"
)

func createUpdateList(p graphql.ResolveParams, obj interface{}) []firestore.Update {
	updateList := []firestore.Update{}
	t := reflect.TypeOf(obj)
	// go over fields from struct
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "ID" {
			continue
		}
		v, ok := p.Args[makeFirstLowerCase(t.Field(i).Name)]
		if ok {
			updateList = append(updateList, firestore.Update{Path: makeFirstLowerCase(t.Field(i).Name), Value: v})
		}
	}

	return updateList
}

func makeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
