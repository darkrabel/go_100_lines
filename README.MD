# go in 1oo lines
In the following 100 lines of code, we’ll go through several examples illustrating Go’s features. We’ll cover how to declare variables, understand Go’s built-in types, work with arrays and slices, cover maps, and touch on the flow of control. From there, we’ll go beyond 100 lines and also look at pointers, structs, and Go’s built-in support for concurrency.
# About go.
Go is an open-source programming language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is often described as “C for the 21st century”, however, it borrows important ideas from several other languages like ALGOL, Pascal, Modula-2, Oberon, CSP, and others. At its core, Go relies on simplicity, reliability, and efficiency to overcome the shortcomings of its ancestors. Go has garbage collection, a package system, first-class functions, lexical scope, immutable strings that rely on UTF-8, and an awesome concurrency model.

As a compiled language, Go is generally faster than interpreted languages and is safer from crashes thanks to its built-in type system. That being said, there is a good balance between expressiveness and safety that gives programmers the benefits of a strong type system without the burden of complex workflows.

The use cases of the language vary from network servers and distributed systems to CLI’s, web and mobile development, scalable databases, and cloud-native applications.
# what you will learn in this repo
*  variable
* types
* array 
* slices 
* maps
* control flow
#  prequest 
Need to know a programing languge.
# discalimer
This is not beginer friedly guide.
# variable 
When writing Go programs, variables must be declared before they can be used. The example below shows how to declare single variables or a group of variables. In the interest of space, the output is displayed as an in-line comment.
* source code file-1 -https://github.com/darkrabel/go_100_lines/blob/main/variabble.go
* source code file-2- https://github.com/darkrabel/go_100_lines/blob/main/variable2.go
# Types
Go offers a rich collection of types, including numericas, booleans, strings, error, and the ability to create custom types. Strings are are a sequence of UTF-8 characters enclosed in double-quotes. Numerical types are the most versatile, with 8, 16, 32, and 64-bit variants for both signed (int) and unsigned (uint) integers.

A byte is an alias for uint8. A rune is an alias for int32. Floats (or floating-point numbers) are either float32 or float64. Complex numbers are also supported and can be represented as complex128 or complex64.

When a variable is declared it is assigned to the natural “null” value of the corresponding type. For example, in var k int, k has the value 0. In var s string, s has the value "“. The example below shows the difference between user-specified types and the default types assigned with a short variable declaration.
* types source code file - https://github.com/darkrabel/go_100_lines/blob/main/types.go
# Arrays
Storing a number of elements in a list can be achieved using arrays, slices, and maps (Go’s version of hash-maps). We’ll consider all three in the examples below. Arrays are defined by their fixed size and a common data type for all elements. Interestingly, the size of the array is part of the type, meaning arrays cannot grow or shrink, otherwise, they would have a different type. Array elements are accessed using square brackets. The example below shows how to declare an array containing strings and how to loop through its elements.
