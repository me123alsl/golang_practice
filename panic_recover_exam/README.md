# 예외 처리

- [예외 처리](#예외-처리)
  - [panic](#panic)
  - [defer](#defer)
  - [recover](#recover)

## panic

panic은 함수를 비정상적으로 종료시키는 함수이다.

```go
func main() {
    panic("Panic!")
}
```

## defer

defer는 함수가 종료될 때 실행되는 함수이다. (finally와 비슷)

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

## recover

recover는 panic으로 인해 종료된 함수를 복구시키는 함수이다.

```go
func main() {
    
    // 함수 종료 시점에 실행
    defer func() {
        // panic이 발생했을 경우, recover를 통해 복구
        if r := recover(); r != nil {
            fmt.Println("Recovered", r)
        }
    }()
    // panic 발생
    panic("Panic!")
}
```
