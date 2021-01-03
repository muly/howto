// inspired by a solution in https://stackoverflow.com/questions/47134293/compare-structs-except-one-field-golang/47134781

// need to extend this to cover
// 	nested structs
// 	other data types like slice and maps

// see go-cmp-except.go for better solution

package main

import (
	"fmt"
	"reflect"
	"time"

	mapset "github.com/deckarep/golang-set"
)

type Parent struct {
	Name string
	// NameP *string
	Date1 time.Time
	Date2 time.Time
	Child child
}

type child struct {
	Address string
}

func EqualExcept(f *Parent, other *Parent, ExceptField mapset.Set) bool {
	val := reflect.ValueOf(f).Elem()
	otherFields := reflect.Indirect(reflect.ValueOf(other))

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if ExceptField.Contains(typeField.Name) {
			continue
		}

		value := val.Field(i)
		otherValue := otherFields.FieldByName(typeField.Name)

		if value.Interface() != otherValue.Interface() {
			return false
		}
	}
	return true
}

func main() {
	n1 := "NNN"
	n2 := "NNN"
	// p1 := "PPP"
	// p2 := "PPP"

	f1 := &Parent{
		Name: n1,
		// &p1,
		Date1: time.Now(),
		Date2: time.Now(),
		Child: child{Address: "AAA"},
	}

	f2 := &Parent{
		Name: n2,
		// &p2,
		Date1: time.Now(),
		Date2: time.Now(),
		Child: child{Address: "AAA"},
	}

	fmt.Println(EqualExcept(f1, f2, mapset.NewSet("Date1", "Date2")))
	fmt.Println(EqualExcept(f1, f2, mapset.NewSet("Name")))

}

