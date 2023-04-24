package main

import "fmt"

/* Generic */
func print1[T any](x T) {
	fmt.Println(x)
}

/* Empty Interface */
func print2(x interface{}) {
	fmt.Println(x)
}

/* Type 제한자 */
type Integer interface {
	~int | ~int16 | ~int32 | ~int64 | ~uint | ~uint16 | ~uint32 | ~uint64
}
type Float interface {
	float32 | float64
}

type Numeric interface {
	Integer | Float
}

type MyInt int

func main() {

	print1(1)
	print1("Hello")
	print1(1.2)

	print2(1)
	print2("Hello")
	print2(1.2)

	fmt.Println(min_1(1, 2))
	fmt.Println(min_2(3, 4))
	fmt.Println(min_3(5, 6))

	var myInt1, myInt2 MyInt = 7, 8
	fmt.Println(min_2(myInt1, myInt2))

}

/* Generic exam1 */
func min_1[T int | int16 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

/* Generic exam1 */
func min_2[T Integer | Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

/* Generic exam1 */
func min_3[T Numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}
