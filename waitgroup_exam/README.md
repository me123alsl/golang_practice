# Waitgroup

- [Waitgroup](#waitgroup)
  - [WaitGroup 이란](#waitgroup-이란)
  - [WaitGroup 사용법](#waitgroup-사용법)
  - [Example](#example)

## WaitGroup 이란

- WaitGroup은 고루틴이 종료될 때까지 대기하는 기능을 제공한다.

## WaitGroup 사용법

- WaitGroup은 3가지 메서드를 제공한다.

  - Add(int) : 고루틴의 개수를 지정한다.
  - Done() : 고루틴이 종료될 때 호출한다.  (고루틴 종료 시 Add(-1) 과 동일)
  - Wait() : 모든 고루틴이 종료될 때까지 대기한다.

## Example

```go
func main() {
    // 변수 설정
    var playTime int = 5 // Sleep Time (Second)
    var wgCount int = 5  // Wait Group Count (Go Routine Count)
    
    // Wait Group 생성
    wg := sync.WaitGroup{}

    // Wait Group Count 증가
    wg.Add(wgCount) 

    // Go Routine 생성
    for i := 1; i <= wgCount; i++ {
        go newFunction(playTime, strconv.Itoa(i), &wg)
    }

    // Defer
    // Wait Group 대기
    defer func() {
        wg.Wait()
        fmt.Println("FINISH")
    }()

}

func newFunction(sleepSecond int, msg string, wg *sync.WaitGroup) {
    fmt.Println(msg)
    time.Sleep(time.Duration(sleepSecond) * time.Second)

    // Wait Group Count 감소
    defer wg.Done()
}
```
