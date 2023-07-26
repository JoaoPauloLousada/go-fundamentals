```go
var i int = 32
var f float32
f = i // error! - Go doesn't support implicit conversions
f = float32(i)// type conversions allow explicit conversion
```
