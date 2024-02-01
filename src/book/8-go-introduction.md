\part{Tutorial to Go}
# Introduction to Go
In this part, we'll start with learning the basics of the Go programming language. We won't delve into every little detail there is about programming in Go, but we will cover those that are crucial going forward in this book.

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.[^2] Attributes such as high performance, low memory usage, easy-to-use concurrency model, simplicity and readability, makes Go one of the best choices for backend applications.

There are many guides online for installing Go, so we won't be covering that here. You simply need an editor such as Vim or VSCode, along with a terminal to run your Go files/project from. After successful installation, you should be able to run Go commands in your terminal like so.
```bash
$ go version
go version go1.21.5 linux/amd64
```

Let's write our first Hello World application using Go next.

## Hello World in Go

Type and save the following code in a file called `main.go`, or any filename of your choice.

```{.go .numberLines}
// main.go
package main // 1.

import "fmt" // 2.

func main() { // 3.
    fmt.Println("Hello, World!") // 4.
}
```

And run it like so:
```bash
$ go run main.go
Hello, World!
```
Congratulations, you just wrote your first program in Go!

___

Let's break it down, piece by piece.

```go
1. package main
```
The name of the package the file is in. When dealing with a single Go source file or when the file in the root of the project, the package name will be `main`.

```go
2. import "fmt"
```
This is the import statement, we use it to import packages from our project and other Go modules. When importing more than one package we surround it with parentheses like so:
```go
import (
    "fmt"
    "math"
)
```
```go
3. func main() {...}
```
This is the entry point of your Go application.

```go
4. fmt.Println("...")
```
We call `Println` from the `fmt` package to write 'Hello World' to the `stdout`.

[^2]: https://en.wikipedia.org/wiki/Go_(programming_language)

Let's learn about variables in go next.
