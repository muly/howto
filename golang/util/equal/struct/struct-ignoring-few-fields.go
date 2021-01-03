// inspired by a solution in https://stackoverflow.com/questions/47134293/compare-structs-except-one-field-golang/47134781

// need to extend this to cover
//  pointers
// 	nested structs
// 	other data types like slice and maps

package main

import (
	"fmt"
	"reflect"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Parent struct {
	Name    string
	NameP   *string
	Date1   time.Time
	Date2   time.Time
	Family  []string
	Child   child
	Members map[string]string
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

		// if diff := cmp.Diff(value.Interface(), otherValue.Interface()); diff != "" {
		// 	return false
		// }
		// or cmp.Equal(x, y, opt) // opt can be cmp.Comparer(func(x, y float64) bool {}
		// or simply
		return cmp.Equal(value.Interface(), otherValue.Interface())
	}
	return true
}

func main() {
	n1 := "NNN"
	n2 := "NNN"
	p1 := "PPP"
	p2 := "PPP2"

	f1 := &Parent{
		Name:   n1,
		NameP:  &p1,
		Date1:  time.Now(),
		Date2:  time.Now(),
		Family: []string{"1", "2"},
		Child:  child{Address: "AAA"},
		Members: map[string]string{
			"brothers": "1",
			"sisters":  "2",
		},
	}

	f2 := &Parent{
		Name:   n2,
		NameP:  &p2,
		Date1:  time.Now(),
		Date2:  time.Now(),
		Family: []string{"1", "2"},
		Child:  child{Address: "AAA"},
		Members: map[string]string{
			"brothers": "1",
			"sisters":  "2",
		},
	}

	fmt.Println(EqualExcept(f1, f2, mapset.NewSet("Date1", "Date2")))                                   // except time variables
	fmt.Println(EqualExcept(f1, f2, mapset.NewSet("Name")))                                             // except normal variable
	fmt.Println(EqualExcept(f1, f2, mapset.NewSet("NameP", "Date1", "Date2")))                          // except pointer and time variables
	fmt.Println(EqualExcept(f1, f2, mapset.NewSet("NameP", "Date1", "Date2", "Members[\"brothers\"]"))) // except pointer and time variables
}
