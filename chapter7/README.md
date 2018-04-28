### Interfaces

An interface type is defined by a set of methods. A value of interface type can hold any value
that implements those methods.

Here is a refactored version of our earlier exampel. This time we made the greeting feature
more generic by defining a function called `Greet` which takes a param of interface type
`Namer`. `Namer` is a new interface we defined which only defines one method: `Name()`. so 
`Greet()` will accept as param any value which has a `Name()` method defined.

To make our `User` struct implement the interface, we defined a `Name()` method. We can
now call `Greet` and pass our pointer to `User` type.


```go
package main

import "fmt"

type User struct{
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Namer interface{
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main(){
	u := &User{"SAADY", "Yousef"}
	fmt.Println(Greet(u))
}

```

We could now define a new type that would implement the same interface and our Greet function would still work.

```go
package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))
}
```


### Errors

An error is anything that can describe itself as an error string. The idea is captured by the predefined, built-in interface type, error, with its single method, Error, returning a string:

```go
type error interface {
	Error() string
}


```

the fmt package's various print routines automatically know to call the method when asked
to print an error.

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did'nt work",
	}
}


func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```