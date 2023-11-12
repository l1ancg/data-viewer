package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

func CreateObject(name string, s interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: graphql.BindFields(s),
	})
}

func CreateArguments(s interface{}, tags ...string) graphql.FieldConfigArgument {
	t := typeOf(s)
	var fieldNames []string
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		name = strings.ToLower(name[0:1]) + name[1:]
		fieldNames = append(fieldNames, name)
	}
	fmt.Println(fieldNames)
	return graphql.BindArg(s, fieldNames...)
}

func CreateSaveResolve(s interface{}, saveFunc func(p interface{})) func(params graphql.ResolveParams) (interface{}, error) {
	t := typeOf(s)
	return func(p graphql.ResolveParams) (interface{}, error) {
		r := reflect.New(t).Interface()
		mapstructure.Decode(p.Args, r)
		saveFunc(r)
		return r, nil
	}
}

func CreateGetResolve(s interface{}, getFunc func(dest interface{}, id int)) func(params graphql.ResolveParams) (interface{}, error) {
	tp := typeOf(s)
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
	tp := typeOf(s)
	return func(p graphql.ResolveParams) (interface{}, error) {
		sliceType := reflect.SliceOf(tp)
		slice := reflect.New(sliceType).Elem()
		listFunc(slice.Addr().Interface())
		return slice.Interface(), nil
	}
}

func typeOf(obj interface{}) reflect.Type {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("expected struct, got %s", t.Kind()))
	}
	return t
}
