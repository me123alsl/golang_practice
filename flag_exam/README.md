# golang Flag

- [golang Flag](#golang-flag)
  - [Flag 란](#flag-란)
  - [예제](#예제)
  - [결과](#결과)

## Flag 란

> Command-line의 argument를 파싱하는 패키지

## 예제

```go
func main() {
    // args1 아규먼트 파싱, default 값은 default1, 설명은 description1
    args1 := flag.String("args1", "default1", "description1")

    // args2 아규먼트 파싱, default 값은 default2, 설명은 description2
    args2 := flag.String("args2", "default2", "description2")

    // 파싱
    flag.Parse()

    // 파싱 결과 출력 (룩업사용)
    fmt.Println("args1:", flag.Lookup("args1").Value)
    fmt.Println("args2:", flag.Lookup("args2").Value)

    // 파싱 결과 출력 (지정된 변수로 출력, 반드시 포인터로 받아야 함)
    // fmt.Println("args1:", *args1)
    // fmt.Println("args2:", *args2)

    // 파싱 외 아규먼트 출력
    fmt.Printf("%+v", flag.Args())
}

```

## 결과

> 아규먼트를 지정하지 않으면 사용법이 출력됨

``` bash
$ go run main.go

Usage of C:\Users\admin\AppData\Local\Temp\go-build2782829117\b001\exe\main.exe:
  -args1 string
        description1 (default "default1")
  -args2 string
        description2 (default "default2")

```

> args1에 "TEST"  
> args2는 지정하지 않음  
> 그 외 아규먼트는 "TEST2", "TEST3"으로 지정

``` bash
$ go run main.go --args1 "TEST" "TEST2" "TEST3"

args1: TEST
args2: default2
[TEST2 TEST3]
```
