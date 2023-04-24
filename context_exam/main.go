package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func main() {
	ContextWithValueTest()
	// Calculate_ContextInValue()
	// Context_Example()
}

/*
@ Context with Value Example
*/
type User struct {
	Id   string
	Name string
	Age  int
}
type UserKey string

func ContextWithValueTest() {
	id := "userId"
	userKey := UserKey(id)
	userValue := User{Id: id, Name: "test", Age: 10}
	ctx := context.WithValue(context.Background(), userKey, userValue)
	fmt.Printf("Key: %+v, KeyType: %+v\nValue: %+v, ValueTpye: %+v\n",
		userKey, reflect.TypeOf(userKey),
		ctx.Value(userKey), reflect.TypeOf(ctx.Value(userKey)))
	go ShowContextValue(ctx, id)
	// Check Context Value
	time.Sleep(1 * time.Second)
}
func ShowContextValue(ctx context.Context, id string) {
	fmt.Printf("%+v", ctx.Value(UserKey(id)))
	defer ctx.Done()
}

/*
@ Context Calculate Example
*/
func Calculate_ContextInValue() {
	var connectionId int = 0
	// ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	go testRutine(ctx, &connectionId)
	time.Sleep(3 * time.Second)
	cancel()
	println("connectionId : ", connectionId)
}

func testRutine(testContext context.Context, connectionId *int) {
	println("Start DateTime : ", time.Now().Format("2006-01-02 15:04:05"))
	for {
		select {
		case <-testContext.Done():
			fmt.Println("time out")
			println("End DateTime : ", time.Now().Format("2006-01-02 15:04:05"))
			return
		default:
			println("### DateTime : ", time.Now().Format("2006-01-02 15:04:05"))
		}
		*connectionId += 1
		time.Sleep(500 * time.Millisecond)
	}

}

/*
ctx.Done : Context가 취소되었을 때, Done 채널에 신호를 보낸다.
ctx.Err : Context가 취소된 이유를 반환한다.
ctx.cancel : Context를 취소한다.
*/
func Context_Example() {
	fmt.Println("############################################")
	fmt.Println("Context Example")
	fmt.Println("############################################")

	// Timeout 설정이 포함된 Context 생성
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Context 값 추가
	ctx = enrichContext(ctx)

	// 쓰레드 실행
	go doSomethingCool(ctx)

	// 2초 후 취소
	select {
	case <-ctx.Done():
		fmt.Println("exceed deadline")
		fmt.Println(ctx.Err())
	}

	time.Sleep(2 * time.Second)
}

// enrichContext : Context에 값을 추가한다.
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

// 쓰레드 : 0.5초 마다 doing something cool 출력
func doSomethingCool(ctx context.Context) {
	rId := ctx.Value("request-id")
	fmt.Println(rId)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("Doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}

}
