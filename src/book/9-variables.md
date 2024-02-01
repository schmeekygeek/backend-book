# Variables

Variables help store data in our programs. They have a name and a value.

In Go, having unused variables in your code, unlike in other programming languages, is disallowed. Your code simply won't compile if you have a variable that hasn't been used somewhere.

You declare a variable in Go by using the `var` keyword.
```go
var <identifier> <type>
```
Here are a few datatypes in Go:

1. `bool`: A boolean (`true` or `false`)
2. `string`: String of characters (`"Hello World"`)
3. `int`: Integer (`34`) - They're 32 bit on 32 bit systems and 64 bit on 64 bit systems.
4. `uint`: Integer (`34`) - Unsigned integer.
5. `float32 float64`: Floating point number (`3.2`)
6. `byte`: Alias for `uint8`

In Go, you can set an `int` to 8, 16, 32 or 64 bits, signed or unsigned.
