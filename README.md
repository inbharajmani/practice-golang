# Quick notes

- [Quick notes](#quick-notes)
  - [Initialization](#initialization)
  - [variables](#variables)
  - [Constants](#constants)
  - [Rune](#rune)
  - [Array](#array)
  - [Slice](#slice)
  - [Map](#map)
  - [Struct](#struct)
  - [Blank identifier](#blank-identifier)
  - [Package level variables](#package-level-variables)
  - [Go routine](#go-routine)
  - [Gaurd clauses](#gaurd-clauses)
  - [Functions](#functions)
  - [Anonymous functions](#anonymous-functions)
  - [Closures](#closures)
  - [Conditionals](#conditionals)
  - [Pass by value and pass by reference](#pass-by-value-and-pass-by-reference)
  - [For loops](#for-loops)
  - [Switch case](#switch-case)
  - [Print statements](#print-statements)
  - [Panic and Recover](#panic-and-recover)

## Initialization

```golang
go mod init [module_nme]
```

To download all dependencies

```golang
go mod download
```

## variables

- Declaration
  - `var i int = 3`
  - It is possible to just define a var without a value but must include a datatype
  - Variables has to be in camelCase
  - int8, int16, int32, int64 -> stores both -ve and +ve integers. Hence for int16 = 65536 values = -32768 to 32767 values can be stored
  - int defaults to 32 or 64 bit depending on system architecture
  - uint for positive integers
  - byte = uint8
  - float32, float64 - no default type size
  - arithmetic operations of different data types are not possible
  - `var myString string = "string"` or \``string`\`
  - `var myBoolean bool = true`
  - defaults values of ints, unsigned ints, floats, rune = 0
  - bool defaults to false
  - string defaults to ""
  - `i := "hello"` -> shorthand / syntactic sugar. Only applies to variables and not const
  - Multivariable initialization `var i, h = 1, 2`
  - IMP: variables are scoped to the closed blocks

## Constants

- Values can't be modified once created
- const always need to be initialized
- const only applies to basic data types - int, float, string, bool

 `const variableA = "constant"`

## Rune

- Values enclosed in single quotes are rune literals
- rune = int32

## Array

- A list of items of same datatype
- Has to specify a static size

```golang
var arr [5]int
arr[0] = 10
arr[1] = 20
```

- Array initialization

```golang
var arr = [5]int{10,20,30,40,50}
```

- Partial initialization

```golang
var arr = [5]int{10,20}
```

 Other values default to 0

- Size can be inferred

```golang
var arr = [...]int{10,20,30,40,50}
```

## Slice

- Array with dynamic sizing

```golang
var slice []int
slice = append(slice, 10)
slice = append(slice, 20)
```

- Slice initialization

```golang
var slice = []int{10,20,30}
```

- Slicing an array

```golang
var arr = [5]int{10,20,30,40,50}
var slice = arr[1:4] // 20,30,40
```

- Slicing a slice

```golang
var slice1 = []int{10,20,30,40,50}
var slice2 = slice1[2:4] // 30,40
```

- Slicing syntax - inclusive of start index and exclusive of end index
- Length of slice - len(slice)
- Capacity of slice - cap(slice) - total elements that can be stored in the underlying array (or) length from the start index to the end of underlying array
- Create slice with make

```golang
var slice = make([]int, length, capacity)
```

- If capacity is not specified, it is set to length
- When appending elements, if capacity is exceeded, a new underlying array is created with double the capacity and elements are copied over
- Copying slices

```golang
var sourceSlice = []int{10,20,30}
var destSlice = make([]int, len(sourceSlice))
copy(destSlice, sourceSlice)
```

- Slice of structs

```golang
type person struct {
    name string
    age  int
}
var people []person
people = append(people, person{name: "Alice", age: 30}, person{name: "Bob", age: 25}, person{name: "Charlie", age: 35})
```

```golang
s := []struct {
    x int
    y string
}{
    {1, "a"},
    {2, "b"},
}
```

- Append is a variadic function. It can take multiple arguments

```golang
slice = append(slice, 30,40,50)
```

## Map

- Key value pairs
`var hashMap = make(map[string]string)`
- Cannot mix diff datatypes in map keys and in values
- Direct initialization or Literal syntax

```golang
var hashMap = map[string]int{
    "key1": 10,
    "key2": 20,
}
```

- Never use var to declare and initialize a map. Always use make or literal syntax
- Adding key value pair
`hashMap["key3"] = 30`
- Deleting key value pair
`delete(hashMap, "key2")`
- Check if key exists

```golang
value, exists := hashMap["key1"]
if exists {
    // key exists
} else {
    // key does not exist
}
```

## Struct

- Key, value of diff data types

```golang
type data struct {
    x int
    y string
    z bool
    something unint
}
```

Usage:

```golang
var structData = data {
    x: 10,
    y: "something sting",
    z: true,
    something: 10
}
```

- You cannot loop over struct fields like in python dict

## Blank identifier

- Used to identify unused variable _

## Package level variables

- All package level variables are available to use across files in same package

- PascalCase variables are accessible outisde a package
- Functions has to be PascalCase for other packages to import it

## Go routine

- green thread - abstraction of actual thread - a builtin thread
- It is not a OS threads like in java
- Also goroutines has builtin communication using "channel"

## Gaurd clauses

- Early returns

## Functions

```golang
func some(x int, y int) (sum int) {
    sum = x + y
    return sum
}
```

Explicit return is good to read and it also initializes variable "sum"

- defer keyword is used to execute any script just before any return statement in a function.
- Useful to close file descriptors, db connections - similar to With(context mangers) in python

 `defer fmt.Println("Exiting the program")`

## Anonymous functions

```golang
func() {
    fmt.Println("Anonymous function")
}()
```

- assign anonymous function to a variable

```golang
sumFunc := func(x int, y int) int {
    return x + y
}
result := sumFunc(10, 20)
fmt.Println("Sum:", result)
```

## Closures

- Function defined inside another function and it can access variables of outer function

```golang
func outerFunction(x int) func(int) int {
    return func(y int) int {    
        return x + y
    }
}
```

## Conditionals

- Syntactic sugar. y scope is limited to the if condition

```golang
if y := f(1,20); y > 3 {

}
else if y < 3 {

}
else {

}
```

## Pass by value and pass by reference

- Pass by value - int, float, bool, struct, array, string
- Zero value of value types - 0, "", false
- Reference types - pointers, slices, maps, channels, interfaces, functions
- Zero value of reference types - nil
- One can use make to create reference types
- make helps to allocate capacity in memory beforehand so appending elements does not create new underlying array every time
- Slices behave bit differently - sometimes as pointers and sometimes as values

## For loops

```golang
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

while loop

```golang
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

for loop with range

```golang
arr := []int{10,20,30,40,50}
for index, value := range arr {
    fmt.Println("Index:", index, "Value:", value)
}
```

Only need index

```golang
for index := range arr {
    fmt.Println("Index:", index)
}
```

Only need value

```golang
for _, value := range arr {
    fmt.Println("Value:", value)
}
```

## Switch case

```golang
switch variable := expression; variable {
case value1:
    // do something
case value2:
    // do something
default:
    // do something
}
```

- Replacement for multiple if else statements

```golang
switch {
case condition1:
    // do something
case condition2:    
    // do something
default:
    // do something
}
```

## Print statements

- fmt.Printf("String: %s, Int: %d, Float: %.2f", strVar, intVar, floatVar)
- fmt.Sprintf - returns formatted string instead of printing
- fmt.Println - prints with space separation

## Panic and Recover

- Panic is similar to raising an exception in other languages
- Recover is used to catch a panic and continue execution of the program

```golang
func safeDivision(num1, num2 int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    result := num1 / num2
    fmt.Println("Result:", result)
}
func main() {
    safeDivision(10, 0)
    fmt.Println("Program continues after recovery.")
}
```
