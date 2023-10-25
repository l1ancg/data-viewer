package utils

import "github.com/graphql-go/graphql"

func CreateObject(obj interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "resource",
		Fields: graphql.BindFields(obj),
	})
}

func CreateArguments(obj interface{}, tags ...string) graphql.FieldConfigArgument {
	return graphql.BindArg(obj, tags...)
}
