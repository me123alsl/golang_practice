package main

import "fmt"

func main() {
	/* panic (founded exception) */
	CallPanic()

	/* finally */
	defer func() {

		/* if panic then recover(catch) */
		if err := recover(); err != nil {
			fmt.Println("Recovered from", err)
		}
	}()
}

func CallPanic() {
	panic("Panic called")
}
