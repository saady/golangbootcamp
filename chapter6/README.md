### Methods

While technically Go isn't an OOP language, types and methods allow for an OO style of programming.
The big difference is that Go does not support type inheritance but instead has a concept of interface.

In this chapter, we will focus on Go's use of methods and interfaces.

`Note`: A faq is `What is the difference between a function and a method ?`. 
A method is a function that has a defined receiver, in OOP terms, a method is a function
on an instance of an object.

Go does not have classes. However, you can define methods on struct types.

The `the method receiver` appears in its own argument list between the `func` keyword and the method name.
Here is an example with a `User` struct containing two fields: `FirstName` and `LastName` of type string.

```go
package  main

import "fmt"

type User struct{
	FistName, LastName string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FistName, u.FistName)
}


func main(){
	u := User{"Yousef", "SAADY"}
	fmt.Println(u.Greeting())
}
```


Note how the methods are being created outside of the struct, if you have been writing OO code for a while,
you might find that a bit odd at first. The method on the `User` type could be defined anywhere in the package.


### Code organization

Methods can be defined on any file in the package, but my recommendation is to organize the code as shown bellow
:

```go
package models

// list of imports 

import (
	"fmt"
)

// list of constants
const (
	ConstExample = "const before vars"
)

// list of variables

var (
	ExportedVar = 42
	nonExportedVar = "so say we all"
)

// Main type(s) for the file
// try to keep the lowest amount of structs per file when possible

type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}

// List of functions
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA",
		},
	}
}

// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
```

In fact, you can define a method on any type you define in your package, not just structs. You cannot define a method on a type from another package, or on a basic type.


### Type aliasing

To define methods on a type you don't own, you need to define an alias for the type you
want to extend:

```go
package main

import  (
	"fmt"
	"strings"
)

type MyStr string

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyStr("test").Uppercase())
}

```

### Methods receivers

Methods can be associated with a named type (`User` for instance) or a pointer to a named
type (`*User`). In the two type aliasing examples above, methods were defined on the value
types (`MyStr`)

There are two reasons to use a pointer receiver. First, to avoid copying the value on each
method call(more efficient if the value type is a large struct). The above example would have 
been better written as follows.

```go
package main

import (
	"fmt"
)

type User struct{
	FirstName, LastName string
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"SAADY", "Yousef"}
	fmt.Println(u.Greeting())
}

```

Remember that Go passes everything by value, meaning that when `Greeting()` is defined
on the value type, everytime you call `Greeting()`, you are copying the `User` struct.
Instead when using a pointer, only the pointer is copied (cheap).

The other reason why you might want to use a pointer is so that the method can modify the 
value that its receiver points to.

```go
package  main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,  Y float64
	
}

func (v *Vertex) Scale (f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

```