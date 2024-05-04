\part{Basics of Go}
# Basics of Go I
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

```go
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

Let's learn about variables in Go next.


## Variables

Variables help store data in our programs. They have a name and a value.

> **Note**: In Go, having unused variables in your code, unlike in other programming languages, is disallowed. Your code simply won't compile if you have a variable that hasn't been used somewhere. 

Let's go over the ways to create variables in Go.

### Declare first, assign later
You declare a variable in Go by using the following syntax:
```go
var <identifier> <type>
```
We make use of the `var` keyword to create variables. If you declare a variable and don't assign a value to it, Go will instantiate it with the default value of its datatype. For example, an unassigned `string` variable will be an empty string, and an unassigned numeric variable such as `int` will be 0. If you have a struct(more on this later) variable, which is simply a way of compositing multiple types, (similar to classes in Java), Go will recurse into the struct until it finds the basic types such as `int` and `string`, and will assign default values to it.

Example:
```go
package main

import "fmt"

func main() {
    var name string
    name = "Josh"
    fmt.Println(name)
}
```

### Declare and Assign
We can declare as well as assign variables like so:
```go
var age int = 32
```
This looks a bit verbose, doesn't it? To solve this problem, Go can also infer types from variables that are declared and assigned without a type as follows:
```go
var age = 12 // the type is int
```
There is an even better way to accomplish this, by making use of a special assignment operator.

### The `:=` operator
The `:=` operator enables us to assign to and declare variables simultaneously.
```go
hobby := "Playing video games"
```
Note that when creating variables using this operator, we don't have to use the `var` keyword. The type will be inferred from the value assigned. When using this operator, the left hand side of the assignment must be a new variable.

```go
var age = 31
age := 43 // This fails as age is already defined
```

### Declare multiple variables
We can declare multiple variables of the same type in Go using the following syntax.
```go
var <identifier>, <identifier>, ... <type>
```
Example:
```go
var height, weight, chest int
```

You can assign values too when declaring multiple variables. The number of values on the right hand side must be the same as the number of identifiers on left hand side, though.

Example:
```go
var foo, bar, baz = 1, "hello", true
```

Using the `:=` operator:
```go
foo, bar, baz := 1, "hello", true
```

### Custom types
In Go, it is possible to declare custom datatypes that are based on inbuilt datatypes internally by using the `type` keyword. 

Syntax:
```go
type <identifier> <datatype>
```

Example:
```go
package main

import "fmt"

func main() {
    type Number int
    var age Number = 24
    fmt.Println(age)
} // prints 24
```

Let's learn how to write conditional statements next.

## Conditional statements

### If-Else statements
The syntax to write an `if` statement is very similar to other programming languages:

```go
if <condition> {
    // todo
} else if <condition> {
    // todo
} else {
    // todo
}
```
Let's look at this example of a program that checks if a person is eligible to vote or not:
```go
age := 18
if age >= 18 {
    fmt.Println("Eligible")
} else {
    fmt.Println("Not Eligible")
}
```
You may or may not surround the condition with parentheses. You can also nest multiple `if-else` statements inside each other.


> **Note**: It is necessary to know that the `else if` or `else` keywords must start right after the last block's curly brace ends. This won't compile:
```go
if <condition> {

}
else if <condition> {

}
```
Go also provides a special option to initialize a variable before the condition as follows:

```go
age := 13
if age = 18; age >= 18 {
    fmt.Println("Eligible") // Prints Eligible
}
```
> **Note**: The value assigned to the variable will persist even after the scope of the conditional ends.

Go also supports logical operators such as `==` (is equal to), and `&&` (logical and).

### Switch statements

Here's the syntax of a switch statement in Go:
```go
switch <variable> {
case <value>:
    // todo
default:
    // todo
}
```

You can have as many cases as you want.

Example:
```go
num := 2
switch num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Not one or two")
}
```

## Functions
You define functions in Go using the `func` keyword.

Syntax:

```go
func <function-name>(<identifier1> <type1>, ...) <return-type> {
    // do something
    return <value>
}
```

Note that the datatype comes after the variable name or identifier. You may or may not have a return type to your function. But if you do have one, you must return an appropriate value pertaining to that return type. You also don't need to write the datatype for multiple parameters that have the same datatype. You can just write the parameter names by separating them with commas and adding the datatype at the end. For example:

```go
func sum(x, y int) int {
    return x + y
}
```

### Passing *n* parameters
If you're uncertain of the number of the parameters that the function must take, you can use the `...` operator. The resultant value of the passed parameter will be an array that you can iterate over.

Example:
```go
func getSum(nums ...int) int {
  sum := 0
  for _, val := range nums {
    sum += val
  }
  return sum
}
```
> *Note:* We will learn about `range` and arrays shortly.

### Anonymous functions
Anonymous functions are also called inline functions or lambda functions in some programming languages. They're useful if you're trying to create a short function body and assign it to a variable to perform some small redundant tasks.

Syntax: 
```go
<identifier> := func (<parameters>) <return-type> {
  return <val>
}
```

You don't need to have a function identifier or name because the variable name that you assign the function to will be used to invoke it.

Example:
```go
func main() {

    sum := func (x, y int) int {
      return x + y
    }

    fmt.Println(sum(7, 6)) // prints 13
    fmt.Println(sum(1, 1)) // prints 2
}
```
