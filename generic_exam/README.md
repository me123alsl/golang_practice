# Generic

- [Generic](#generic)
  - [Generic Function](#generic-function)
    - [Generic Function 정의](#generic-function-정의)
    - [타입 지정자](#타입-지정자)
  - [Example](#example)

## Generic Function

### Generic Function 정의

> 함수 정의 시 '[ ]' 안에 타입을 지정한다.

```go
func Method[T type](t T) {
    // ...
}
```

### 타입 지정자  

> interface 타입의 지정할 타입을 | (or) 로 구분하여 지정한다.

```go
type Integer interface {
 ~int | ~int16 | ~int32 | ~int64 | ~uint | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
 float32 | float64
}

type Numeric interface {
 Integer | Float
}
```

## Example

```go
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
```
