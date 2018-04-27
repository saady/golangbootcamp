## Types

### Structs

a struct is a collection of fields/properties. You can define new types as structs or interfaces.

You don't need to define getters or setters on struct fields, they can be accessed automatically. however, note that only exported fields (Capitalized) can be accessed from outside the package.

A struct literal sets a newly allocated struct value by listing the values of its fields. You can list just a subset of fields by using the `"Name":` `syntax` (the order if named fields is irrelevant when using this syntax). The special prefix `&` constructs a pointer to a newly allocated struct.

```go
package main

import (
    "fmt"
    "time"
)

type Bootcamp struct {
    Lat float64
    Lon float64
    Date time.Time
}

func main() {
    fmt.Println(Bootcamp{Lat: 34.012836, Lon:-118.495338, Date: time.Now()})
}
```

output

```shell
{124545 24577 2017-11-29 13:33:28.883516231 +0000 WET}
```

Declaration of struct literals:

```go
package main

import "fmt"

type Point struct {
    x, y int
}

var (
    p = Point{1, 2}
    q = &Point{1, 2}
    r = Point{x: 1}
    s = Point{}
)

func main(){
    fmt.Println(p, q, r, s)
}
```
Output:

```shell
{1 3} &{1 3} {1 0} {0 0}
```

Accessing fields using the dot notation:

```go
package main

import (
    "fmt"
    "time"
)

type Bootcamp struct {
    Lat, Lon float64
    Date time.Time
}

func main() {
    event := Bootcamp{
        Lat: 34.2567,
        Lon: -112.25236
    }
    event.Date = time.Now()
    fmt.Prinf("Event on %s, location (%f, %f)", event.Date, event.Lon, event.Lat)
}
```

Output

```shell
event on 2017-11-29 13:46:07.912421294 +0000 WET, location (-134.265300, 23.345550)%
```

### Initializing

Go supports the `new` expression to allocate a zeroed value of the requested type and to return a pointer to it.

```go
x := new(int)
```

a common way to "initialize" a variable containing a struct or a reference to one, is to create a struct literal. Another option is to create a constructor. This is usually done when the zero value isn't good enough and you need to set some default field value instance.

```go
package main

import "fmt"

type Bootcamp struct {
    Lat float64
    Lon float64
}

func main() {
    x := new(Bootcamp)
    y := &Bootcamp{}
    fmt.Println(*x == *y)
}
```

Output

```shell
true

```

Note that slices, maps and channels are usually allocated using `make` so the data structure these types are built upon can be initialized.

### Composition vs Inheritance

Composition is a well understood concept for most OOP programmers and Go supports it, here is an example of the problem it's solving:

```go
package main

import "fmt"

type User struct {
    Id       int
    Name     string
    Location string
}

type Player struct {
    Id       int
    Name     string
    Location string
    GameId   int
}

func main() {
    p := Player{}
    p.Id = 42
    p.Name = "Yousef"
    p.Location = "Casablanca"
    p.GameId = 252
    fmt.Printf("%+v", p)
}
```

Output

```shell
{Id:42 Location:Casablanca Name:Youssef GameId:2335}%
```

The above example demonstrates a classic OOP challenge, our `Player` struct has the same fields as the `User` struct but it also has a `GameId` field. Having to duplicate the field names is not a big deal, but it can be simplified by composing our struct.

```go

type User struct {
    Id             int
    Name, Location string
}

type Player struct {
    User
    GameId int
}

```

we can initialize a new variable of type `Player` two different ways.

Using the dot notation to set the fields:

```go

package main

import "fmt"

type User struct {
    Id             int
    Name, Location string
}

type Player struct {
    User
    GameId int
}

func main() {
    p := Player{
        User{Id: 42, Name: "Youssef", Location: "Casablanca"}, 2345
        }
        // directly set a field defined on the player struct
        p.Id = 33
        fmt.Printf("%+v", p)
}
```
