## The basics

### Variables & inferred typing

The `var` statement declare a list of variables with the type declared last

```go

var (
    name string
    age int
    location string
)
```

or even

```go
var (
    name, location string
    age int
)
```

variables can also be declared one by one

```go

var name string
var location string
var age int

```

a var declaration can include initializers, one per variable

```go

var (
    name string     = "saady"
    location string = "Casablanca"
    age int         = 25
)

```

if an initializer is present, the type can be omitted, the variable will take the type of the initializer (inferred typing)

```go

var (
    
    name     = "saady"
    location = "Casablanca"
    age      = 25
)
```

We can also initialize the variables on the same line

```go

var (
    name, location, age =   "saady", "Casablanca", 25
)
```

Inside a function, the `:=` short assignement can be used in place of a var declaration with implicit type.

```go

func main() {
    name, location := "SAADY", "Casablanca"
    age := 25
    fmt.Printf("%s (%d) %s", name, age, location)
}

```

To run the code above:

```shell
$ go fmt main_short_assingnement.go
$ go run main_short_assingnement.go
````

output:

```shell
SAADY (25) Casablanca
```

A variable can contain `any type` including functions


```go

func main() {
    action := func() {
        // doing something
    }
    action()
}
```

Outside a function, every construct begins with a keyword (`var`, `func`, and so on) and the `:=` construct is not available.


### Constants

Constants are declared like variables, but with the `const` keyword.

Constants can only be `character`, `string`, `boolean`, or `numeric` values and cannot be declared with using the `:=` syntax.
An untype contants takes the type needed by its context.

```go

const Pi = 3.14
const (
    StatusOk        = 200
    StatusCreated   = 201
    StatusAccepted  = 202
)
```

```go

package main

import (
    "fmt"
)

const (
    Pi = 3.14
    Truth = false
    Big = 1 << 62
    Small = Big >> 61
)

func main() {
    const Greeting = "nihao"
    fmt.Println(Greeting)
    fmt.Println(Pi)
    fmt.Println(Truth)
    fmt.Println(Big)
}
```

### Printing Constants and Variables

While you can print the value of a variable or constant using the builtin `print` and  `println` functions, the more idiomatic and 
felxible way is to use the `fmt` package

```go

func main() {
    cylonModel := 6
    fmt.Println(cylonModel)
}
```

`fmt.Println` prints the passed in variables values and appends a newline. `fmt.Printf` is used when you want to print one or multiple values
using a defined format specifier.

```go

func main() {
    name := "Yousef"
    aka := fmt.Sprintf("yous(%d)f", 3)
    fmt.Println("%s is know as %s", name, aka)

}
```


### Packages and Imports


Every Go program is make up of packages, Programs start running in package main

```go

package main

func main() {
    print("hello world! \n")
}
```

if you are wrinting an executable code (versus library), then you need to define a `main`
package and a main function which will be the entry point to your software.

By convention the package name is the same as the last element of import path. For instance the `math/rand` package comprises 
files that begin with the statement package rand.

```go

import "fmt"
import "math/rand"

```

or Grouped

```go

import (
    "fmt"
    "math/rand"
)
```

Usually, non standard lib packages are namespaced using a web url.

```go

import "github.com/mattetti/goRailsYourself/crypto"

```

### Code location

The snippet above basically tells the compiler to import the crypto package available at the 
`github.com/mattetti/goRailsYourself/crypto` path. it does not mean that the compiler will automatically pull down the repository.

You need to pull down the code yourself. The easiest way is to use the `go get` command
provided by go.


```shell

go get github.com/mattetti/goRailsYourself/crypto

```

this command will pull down the code and put it into your GO path. when installing Go, we set the `GOPATH` env var and that is what's used to
store binaries and libs.

That's also where you should store your code (your workspace).

```shell

$ ls $GOPATH/src

bin pkg src

```

`bin` folder will contain go compiled binaries.
`pkg` folder contains the complied versions of the available libs so the compiler can link against them without 
recompiling them.
`src` folder contains all the Go source code organized by import path:

```shell
$ ls $GOPATH/src
bitbucket.org	code.google.com	github.com	launchpad.net

```

```shell
$ ls $GOPATH/src/github.com/mattetti
goblin			goRailsYourself		jet

```

when starting a new program or lib, it is recommended to do so inside the `src` folder, using a fully qualified path
(for instance `github.com/saady/myprojectname`)


### Exported names

After importing a package, you can refer to the names it exports(meaning variables, methods and functions that are variable from outside package)
In Go, a name is exported if it begins with a Capital letter. `Foo` is an exported name, as is `FOO`. The name `foo` is not exported.

See the difference between:

```go

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.pi)
}

```

and 


```go

func main() {
    fmt.Println(math.Pi)
}

```

`Pi` is exported and can be accessed from outside the page, while `pi` isn't available.

```shell
cannot refer to unexported name math.pi

```

### Functions, signatures, return values, named results

A function can take zero or more typed arguments. The type comes after the variable  name.
Functions can be defined to return any number of values that are always typed.

```go

package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}

```

Instead of declaring the type of each parameter, we only declare one type that applies to both


```go

package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Prinln(add(4, 5))
}

```


In the following example, the `location` function returns two string values.

```go

func location(city string) (string, string) {
    var region string
    var continent string

    switch city {
    case "Agadir", "Ouled Teima", "Taroudant":
        region, continent = "Souss massa", "Africa"
    case "New Yorl", "NYC":
        region, continent = "New York", "North America"
    default:
        region, continent = "Unknow", "Unknown"
    }
    return region, continent
}

func main() {
    region, continent := location("Agadir")
    fmt.Printf("SAADY lives in %s, %s", region, continent)
}

```

Functions take parameters. in Go, functions can return multiple "result parameters", not
just a single value. They can be named and act just like variables.

If the result parameters are named, a return statement without arguments returns the current values of the results.

```go

func location(name, city string) (region, continent string) {
    switch city {
    case "Agadir":
        continent = "Africa"
    default:
        continent = "Unknown"
    }
    return
}

func main() {
    region, continent := location("SAADY", "Agadir")
    fmt.Printf("%s lives in %s", region, continent)
}
```

I personally recommend against using named return parameters because they often cause more confusion than they save time or help clarify your code.

### Pointers

Go has pointers, but no pointer arithmetic. Struct fields can be accessed through a struct
pointer. The indirection through the pointer is transparent (you can directly call fields and methods on a pointer).

Note that by default Go passes arguments by value (copying the arguments), if you want to
pass the arguments by reference, you need to pass the pointers (or use a structure using reference values like slices and maps).

To get the pointer of a value, use the `&` symbol in front of the value; to dereference a pointer,
use the symbol `*`.

Mathods are often defined on pointers and not values (although they can be defined on both), so you will often store a pointer in a variable
as in the example below:

```go

client := &http.Client{}
resp, err := client.Get("http://google.fr")

```

### Mutability

in Go, only contants are immutable. However because arguments are passed by value, a
function receiving a value argument and mutating it, won't mutate the original value.


```go

package main

import "fmt"

type Artist struct {
    Name, Genre string
    Songs       int
}

func newRelease(a Artist) int {
    a.Songs++
    return a.Songs
}

func main(){
    me := Artist{Name: "SAADY", Genre: "rock", Songs: 1}
    fmt.Printf("%s released their dth song \n", me.Name, newRelease(me))
    fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
```

Output

```shell

SAADY released their 2'th song
SAADY has a total of 1 songs%

```

As you can see the total amount of songs on the `me` variable's value wasn't changed. To
mutate the passed value, we need to pass it by reference, using a pointer.

```go
package main

import "fmt"

type Artist struct {
    Name, Genre string
    Songs       int
}

func newRelease(a Artist) int {
    a.Songs++
    return a.Songs
}

func main(){
    me := &Artist{Name: "SAADY", Genre: "rock", Songs: 1}
    fmt.Printf("%s released their dth song \n", me.Name, newRelease(me))
    fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
```

Output

```shell
SAADY released their 2th song
SAADY has a total of 2 songs%
```

---- END OF CHAPTER 2 ----