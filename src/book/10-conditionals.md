# Conditional statements

## If-Else statements
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

## Switch statements

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
