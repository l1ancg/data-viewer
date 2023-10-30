package utils

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

func CreateObject(obj interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "resource",
		Fields: graphql.BindFields(obj),
	})
}

func CreateArguments(obj interface{}, tags ...string) graphql.FieldConfigArgument {
	return graphql.BindArg(obj, tags...)
}

func CreateSaveResolve(s interface{}, saveFunc func(p interface{})) func(params graphql.ResolveParams) (interface{}, error) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("expected struct, got %s", t.Kind()))
	}
	return func(p graphql.ResolveParams) (interface{}, error) {
		r := reflect.New(t)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Name
			objValue := reflect.ValueOf(r).Elem()
			fieldValue := objValue.FieldByName(fieldName)
			if fieldValue.IsValid() {
				val := reflect.ValueOf(p.Args[field.Name]).Convert(field.Type)
				fieldValue.Set(val)
			}
		}
		saveFunc(r)
		return r, nil
	}
}

func CreateGetResolve(s interface{}, getFunc func(dest interface{}, id int)) func(params graphql.ResolveParams) (interface{}, error) {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, isOK := params.Args["id"].(int)
		if isOK {
			dest := reflect.New(reflect.TypeOf(s))
			getFunc(dest, id)
		}
		return nil, errors.New("请指定ID参数")
	}
}

func CreateListResolve(s interface{}, listFunc func(interface{})) func(p graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		tp := reflect.TypeOf(s)
		sli := reflect.MakeSlice(reflect.SliceOf(tp), 0, 0)
		listFunc(sli)
		return sli, nil
	}
}
