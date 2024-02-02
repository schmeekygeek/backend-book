# Variables

Variables help store data in our programs. They have a name and a value.

In Go, having unused variables in your code, unlike in other programming languages, is disallowed. Your code simply won't compile if you have a variable that hasn't been used somewhere. 

Let's go over the ways to create variables in Go.

## Declare first, Assign later
You declare a variable in Go by using the following syntax:
```go
var <identifier> <type>
```
We make use of the `var` keyword to create variables. If you declare a variable and don't assign a value to it, Go will instantiate it with the default value of its datatype. For example, an unassigned `string` variable will be an empty string, and an unassigned numeric variable such as `int` will be 0. If you have a struct(more on this later) variable, which is simply a way of compositing multiple types, (similar to classes in Java), Go will recurse into the struct until it finds the basic types such as `int` and `string`, and will assign default values to it.

Example:
```go
func main() {
    var name string
    name = "Josh"
    fmt.Println(name)
}
```

## Declare and Assign
We can declare as well as assign variables like so:
```go
var age int = 32
```
This looks a bit verbose, doesn't it? To solve this problem, Go can also infer types from variables that are declared and assigned without a type as follows:
```go
var age = 12 // the type is int
```
There is an even better way to accomplish this, by making use of a special assignment operator.

## The `:=` operator
The `:=` operator enables us to assign to and declare variables simultaneously.
```go
hobby := "Playing video games"
```
Note that when creating variables using this operator, we don't have to use the `var` keyword. The type will be inferred from the value assigned. When using this operator, the left hand side of the assignment must be a new variable.

```go
var age = 31
age := 43 // This fails as age is already defined
```

## Declare multiple variables
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
