# Context

- [Context](#context)
  - [Context Interface](#context-interface)
  - [Methods](#methods)
    - [WithCancel()](#withcancel)
    - [WithDeadline()](#withdeadline)
    - [WithTimeout()](#withtimeout)
    - [WithValue()](#withvalue)

## Context Interface

```go
// Context Interface
type Context interface {
    // Done() 채널을 통해 Context가 취소되었는지 확인한다.
    Done() <-chan struct{}
    // Err() 함수를 통해 Context가 취소된 이유를 확인한다.
    Err() error
    // Deadline() 함수를 통해 Context의 Deadline을 확인한다.
    Deadline() (deadline time.Time, ok bool)
    // Value() 함수를 통해 Context에 저장된 Value를 확인한다.
    Value(key interface{}) interface{}
}
```

## Methods

### WithCancel()

>Context 취소기능을 갖는 Context 객체를 생성한다.

```go
// WithCancel() 함수 원형- [Context](#context)
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// Example
// ctx, cancel 객체 생성하여, ctx.Done() 채널을 통해 취소 여부를 확인한다.
ctx, cancel := context.WithCancel(context.Background())
go func(context Context.context) {
   select {
    case <- context.Done():
        fmt.Println("Context Done")
    default :
        fmt.Println("Context Not Done")
        time.Sleep(1 * time.Second)
    }
}(ctx)

// ctx에 취소 요청을 보낸다. (Done 채널에 값이 전달된다.)
cancel()
```

---

### WithDeadline()

>Context Deadline이(Time) 포함된 Context 객체를 생성한다.

```go
// WithDeadline() 함수 원형
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

// Example
// ctx, cancel 객체 생성하여, 현재시간 2초뒤 채널을 Done() 채널을 통해 취소 여부를 확인한다.
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
go func(context Context.context) {
   select {
    case <- context.Done():
        fmt.Println("Context Done")
    default :
        fmt.Println("Context Not Done")
        time.Sleep(1 * time.Second)
    }
}(ctx)
```

---

### WithTimeout()

>Context Timeout이 포함된 Context 객체를 생성한다.

```go
// WithTimeout() 함수 원형
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

// Example
// ctx, cancel 객체 생성하여, 3초 후 채널을 Done() 채널을 통해 취소 여부를 확인한다.
ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
go func(context Context.context) {
   select {
    case <- context.Done():
        fmt.Println("Context Done")
    default :
        fmt.Println("Context Not Done")
        time.Sleep(1 * time.Second)
    }
}(ctx)

```

### WithValue()

> Context에 Key,Value를 추가한다.

```go
// WithValue() 함수 원형
func WithValue(parent Context, key, val interface{}) Context

// Example
// ctx에 Key, Value를 함께 전달한다.
ctx := context.WithValue(context.Background(), "key", "value")
// ctx에서 Key를 통해 Value를 가져온다.
value := ctx.Value("key")
```