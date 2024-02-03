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
Example:
```go
age := 18
if age >= 18 {
    fmt.Println("Eligible")
}
```
You may or may not surround the condition with parentheses. Also, it is necessary to know that the `else if` or `else` keywords must start right after the previous block ends. This won't compile.
```go
if <condition> {

}
else if <condition> {

}
```
Go also provides a special option to initialize a variable before testing the condition as follows:

```go
age := 13
if age = 18; age >= 18 {
    fmt.Println("eligible") // Prints eligible
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
case <value>:
    // todo
default:
    // todo
}
```
Example:
```go
num := 3
switch num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three")
default:
    fmt.Println("Not one, two or three")
}
```
