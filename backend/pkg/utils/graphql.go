package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

func CreateObject(name string, obj interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
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
		r := reflect.New(t).Interface()
		mapstructure.Decode(p.Args, r)
		saveFunc(r)
		return r, nil
	}
}

func CreateGetResolve(s interface{}, getFunc func(dest interface{}, id int)) func(params graphql.ResolveParams) (interface{}, error) {
	tp := reflect.TypeOf(s)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, isOK := params.Args["id"].(int)
		if isOK {
			dest := reflect.New(tp).Interface()
			getFunc(dest, id)
			return dest, nil
		}
		return nil, errors.New("请指定ID参数")
	}
}

func CreateListResolve(s interface{}, listFunc func(interface{})) func(p graphql.ResolveParams) (interface{}, error) {
	tp := reflect.TypeOf(s)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return func(p graphql.ResolveParams) (interface{}, error) {
		sliceType := reflect.SliceOf(tp)
		slice := reflect.New(sliceType).Elem()
		listFunc(slice.Addr().Interface())
		return slice.Interface(), nil
	}
}
