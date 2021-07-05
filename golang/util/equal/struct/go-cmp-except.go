// compare 2 structs data while ignoring
// [x] few fields: Date1, Date2
// 		cmpopts.IgnoreFields(Parent{}, "Date1", "Date2")
// [x] nested struct:
//		cmpopts.IgnoreFields(Parent{}, "Child")
// [x] few fields in a nested struct: child.Address2
//		cmpopts.IgnoreFields(child{}, "Address", "Address2")))
// [x] pointer fields: NameP
// 		cmpopts.IgnoreFields(Parent{}, "NameP")
// [x] few keys in a map
//		cmp.FilterPath(isIgnoredKey, cmp.Ignore()),
// [x] fields of other data types like slice and maps
// 		cmpopts.IgnoreFields(Parent{}, "Family", "Members")
// [] zero values: string, int, pointers, pointer to struct, slices, slices of pointers, etc
//
// [] 

package main

import (
	"fmt"
	"time"

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
	Address  string
	Address2 string
}

func main() {
	n1 := "NNN"
	n2 := "NNN"
	p1 := "PPP"
	p2 := "PP"

	f1 := &Parent{
		Name:   n1,
		NameP:  &p1,
		Date1:  time.Now(),
		Date2:  time.Now(),
		Family: []string{"1", "2"},
		Child:  child{Address: "AAA", Address2: "BBB"},
		Members: map[string]string{
			"brothers":    "1",
			"sisters":     "2",
			"created_dt":  "some date 1",
			"modified_dt": "some date 2",
		},
	}

	f2 := &Parent{
		Name:   n2,
		NameP:  &p2,
		Date1:  time.Now(),
		Date2:  time.Now(),
		Family: []string{"2", "2"},
		Child:  child{Address: "AAA", Address2: "BB"},
		Members: map[string]string{
			"brothers":    "1",
			"sisters":     "2",
			"created_dt":  "some date 3",
			"modified_dt": "some date 4",
		},
	}

	fmt.Println(cmp.Equal(f1, f2, cmpopts.IgnoreFields(Parent{}, "Date1", "Date2", "NameP", "Child", "Family", "Members")))
	fmt.Println(cmp.Equal(f1, f2,
		cmpopts.IgnoreFields(Parent{}, "Date1", "Date2", "NameP", "Family"),
		cmpopts.IgnoreFields(child{}, "Address2"),
		cmp.FilterPath(isIgnoredKey, cmp.Ignore()),
	))

}

func isIgnoredKey(p cmp.Path) bool {
	step, ok := p[len(p)-1].(cmp.MapIndex)
	return ok && (step.Key().String() == "created_dt" || step.Key().String() == "modified_dt")
}

// Notes:
// 		file names must be exported
// 		in case of nested struct, to ignore the complete struct, pass the field name of the nested struct.
//
