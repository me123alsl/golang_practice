package main

import (
	"flag"
	"fmt"
)

func main() {
	// args1 아규먼트 파싱, default 값은 default1, 설명은 description1
	args1 := flag.String("args1", "default1", "description1")

	// args2 아규먼트 파싱, default 값은 default2, 설명은 description2
	args2 := flag.String("args2", "default2", "description2")

	// 파싱
	flag.Parse()

	// 파싱 결과 출력 (룩업)
	fmt.Println("args1:", flag.Lookup("args1").Value)
	fmt.Println("args2:", flag.Lookup("args2").Value)

	// 파싱 결과 출력 (지정된 변수로 출력)
	fmt.Println("args1:", *args1)
	fmt.Println("args2:", *args2)

	// 파싱 외 아규먼트 출력
	fmt.Printf("%+v", flag.Args())
}
