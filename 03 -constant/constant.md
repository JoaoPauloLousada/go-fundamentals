```go
const a = 42 // implicitly typed
const b string = "hello, world" // explicitly typed
const c = a // one constant can be assigned to another
const ( // group of constants
  d = true
  e = 3.14
)
const f = 2*5 // 10 - constant expression
const g = "hello, "+"world"// must be calculable at compile time
const h = someFunction()// does NOT work - can't be evaluated at compile time
const (
  i = "foo"
  j // "foo - unassigned constants receive previous value"
)
const k = iota //0 - iota is related to position in constant group
const (
  l = iota //0
  m // 1
  n = 3 * iota // 6
)
const (
  o = iota // 0 - iota resets to zero with each group
)
```
