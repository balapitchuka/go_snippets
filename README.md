## go

#### udemy course reference link:
> https://github.com/StephenGrider/GoCasts

integer types : int8, int16, int32, int64; and uint8, uint16, uint32, and uint64
float types : float32(6 decimal precision) and float64(15 decimal precision)
complex types : complex64 and complex128

'/' operation has to be performed between similar types
```go
x := 12 
k := 5 
fmt.Println(x) 
fmt.Printf("Type of x: %T\n", x)    // Type of x: int

div := x / k 
fmt.Println("div", div) // div 2  (integer division)
```