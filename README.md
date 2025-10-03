# Quick notes
- [Quick notes](#quick-notes)
  - [Initialization](#initialization)
  - [Basic commands](#basic-commands)
  - [Defer](#defer)
  - [Variables](#variables)
  - [Declaration](#declaration)
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
  - [Conditionals](#conditionals)
  - [Pass by value and pass by reference](#pass-by-value-and-pass-by-reference)


## Initialization

```bash
go mod init [module_nme]
```

## Basic commands

```bash
go run [file_name].go
go build [file_name].go
go clean
```

## Defer
- defer happens when function exits irrespective of how it exits - normal return, panic, error
- defer happens in LIFO order
- defer is useful to close file descriptors, db connections etc
- defer works at function level and not at block 

## Variables

## Declaration
    - `var i int = 3`
    - It is possible to just define a var without a value but must include a datatype
    - Variables has to be in camelCase
    - int8, int16, int32, int64 -> stores both -ve and +ve integers. Hence for int16 = 65536 values = -32768 to 32767 values can be stored
    - int defaults to 32 or 64 bit depending on system architecture
    - uint for positive integers
    - float32, float64 - no default type size
    - arithmetic operations of different data types are not possible
    - `var myString string = "string"` or \``string`\`
    - `var myBoolean bool = true`
    - defaults values of ints, unsigned ints, floats, rune = 0
    - bool defaults to false
    - string defaults to ""
    - `i := "hello"` -> shorthand / syntactic sugar. Only applies to variables and not const
    - Multivariable intialization `var i, h = 1, 2`
    - IMP: variables are scoped to the closed blocks

## Constants
    - Values can't be modified once created
    - const always need to be initialized
    
 `const variableA = "constant"`

## Rune

- Values enclosed in single quotes are rune literals
- Strings are UTF-8 encoded while runes are UTF-32 encoded
- Rune is an alias for int32

## Array
  - A list of items of same datatype
  - Has to specify a static size

## Slice
- Array with dynamic sizing
- Slice has 3 components - pointer to array, length, capacity
- Capacity is the max size of the slice
- When creating a slice from an array, the capacity is from the starting index to the end of the array
- When appending to a slice, if the capacity is exceeded, a new array is created
- Slices are reference types - changes to the slice will affect the underlying array

## Map
- Key value pairs
`var hashMap = make(map[string]string)`
- Cannot mix diff datatypes in map keys and in values

## Struct
- Key, value of diff data types
```
type data struct {
    x int
    y string
    z bool
    something unint
}
```

Usage:

var structData = data {
    x: 10,
    y: "something sting",
    z: true,
    something: 10
}

##  Blank identifier
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

func some(x int, y int) (sum int) {
    sum = x + y
    return sum
}

Explicit return is good to read and it also initializes variable "sum"

- defer keyword is used to execute any script just before any return statement in a function.
 - Useful to close file descriptors, db connections - similar to With(context mangers) in python

 `defer fmt.Println("Exiting the program")`

## Conditionals

- Syntactic sugar. y scope is limited to the if condition

if y := f(1,20); y > 3 {

}

## Pass by value and pass by reference
- Pass by value - int, float, bool, struct, array, string
- Pass by reference - mutex, map
- Slices behave bit differently - sometimes as pointers and sometimes as values