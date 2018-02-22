package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"p_name" bson:"pName"`
	Age  int    `json:"p_age" bson:"pAge"`
}

func main() {
	f := &Person{"Tom", 30}
	describe(f)
}

func describe(f interface{}) {
	fmt.Println(reflect.TypeOf(f))
	val := reflect.TypeOf(f).Elem()
	fmt.Println(val)
	for i := 0; i < val.NumField(); i++ {
		structField := val.Field(i)
		fieldName := structField.Name
		jsonTag := structField.Tag.Get("json")
		bsonTag := structField.Tag.Get("bson")
		fmt.Printf("Field : %s jsonTag: %s bsonTag: %s\n",
			fieldName, jsonTag, bsonTag)
	}
}
