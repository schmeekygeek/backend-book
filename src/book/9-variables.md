# Variables

Variables help store data in our programs. They have a name and a value.

In Go, having unused variables in your code, unlike in other programming languages, is disallowed. Your code simply won't compile if you have a variable that hasn't been used somewhere. 

Let's go over some ways to create variables in Go.

## Declare first, Assign later
You declare a variable in Go by using the following syntax:
```go
var <identifier> <type>
```
We make use of the `var` keyword to create variables. If you declare a variable and don't assign a value to it, Go will instantiate it with the default value of its datatype. For example, an unassigned `string` variable will be an empty string, and an unassigned numeric variable such as `int` will be 0. If you have a struct(more on this later) variable, which is simply a way of composing multiple types, (similar to classes in Java), Go will recurse into the struct until it finds the basic types such as `int` and `string`, and will assign default values to it.

## Declare and Assign

## `:=` Operator
