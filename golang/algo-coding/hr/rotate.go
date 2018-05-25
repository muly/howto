//HR - Data Structures - Array - Left Rotation
//https://www.hackerrank.com/challenges/array-left-rotation/problem

//TODO: WIP code. rotate() is incomplete and not tested

package main

import "fmt"

func main() {
	cmds, err := getInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	size := cmds[0]
	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, cmds[i+2])
	}

	for _, v := range rotate(a, cmds[1]) {
		fmt.Printf("%v ", v)
	}

}

func getInput() (cmds []int, err error) {
	input := 0

	for {
		if _, err = fmt.Scanf("%d", &input); err != nil {
			if err.Error() == "EOF" {
				err = nil
				break
			} else {
				return nil, err
			}
		}
		cmds = append(cmds, input)
	}
	return
}
func rotate(a []int, n int) []int {
	out := make([]int, 0, len(a))
	for i, _ := range a {
		j := i + n // TODO: need to handle the overflow
		out[i] = a[j]
	}

	return a
}
