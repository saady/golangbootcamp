## Collection Types

### Arrays

The type `[n]T` is an array of `n` values of type `T`

The expression:

```go
var a [10]int
```
declares a variable `a` as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but 
don't worry; Go provides a convenient way of working with arrays.

```go
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "hello"
    a[1] = "world"
    fmt.Println(a[0], a[1])
    fmt.Println(a)
}
```

Output

```shell
hello world
[hello world]
```

You can also set the array entries as you declare the array

```go

package main

import "fmt"

func main() {
    a := [2]string{"hello", "world"}
    fmt.Printf("%q", a)
}
```

Output

```shell
["hello" "world"]%
```

You can use an illepsis to use an implicit length when you pass the values:

```go

package main

import "fmt"

func main() {
    a := [...]string{"hello", "world"}
    fmt.Printf("%q", a)
}
```

Output

```shell

["hello" "world"]%
```

### Multi-dim arrays

You can also create multi-dim arrays:


```go

package main

import "fmt"

func main() {
    var a [2][3]string
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            a[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
        }
    }
    fmt.Printf("%q", a)

}

```

Output

```shell
[["row 1 - column 1" "row 1 - column 2" "row 1 - column 3"] ["row 2 - column 1" "row 2 - column 2" "row 2 - column 3"]]%
```

### Slices

Slices wrap arrays to give a more general, powerful, and convinient interface to sequences of data.
Except for items with explicit dimension such as transformation matrices, most array programming in GO is done
with slices rather then simple arrays.

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.

if a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous a pointer to the underlying array.

A slice points to an array of values and also includes a length. Slices can be resized since they are just a wrapper on top of another data structure.

`[]T` is a slice with element of type `T`.



```go
package main
import fmt

func main() {
    p := []int{2, 3, 5, 7, 7}
}
fmt.Prinln(p)
```

Output

```
[2 3 5 7 7]
```

### Slicing a slice

Slice can re-sliced, creating a new slice value that points to the same array.

The expression

```shell
s[lo:hi]
```

evaluates to a slice of the element from `lo` through `hi-1`, inclusive. Thus

```shell
s[lo:lo]
```

is empty and 

```
s[lo:lo+1]
```
has one element

Note: `lo` and `hi` would be integers representing indexes.


```go
package main

import "fmt"

func main(){
    mySlice := []int{1, 2, 3, 4, 6}
    fmt.Println(mySlice)

    fmt.Println(mySlice[1:4])
    fmt.Println(mySlice[:3])
    fmt.Println(mySlice[:4])
}
```

### Making slice


Besides creating slices by passing the value right away (slice lateral), you can use `make`.

You create an empty slice of a specific length and then populate each entry:

```go
package main

import fmt

func main() {
    cities := make([]string, 3)
    cities[0] = "Casablanca"
    cities[1] = "Agadir"
    cities[2] = "Rabat"
    fmt.Println("%q", cities)
}
```


### Appending to a slice

Note however, that you would get a runtime error if you were to do that

```shell
cities := []string{}
cities[0] = "Santa Monica"
```

as explained above, a slice is seating on top of an array, in this case, the array is empty and the slice can't set a value in the referred array. There is a way to do that though, and that is by using the `append` function:


```go
package main

import "fmt"

func main() {
    cities := []string{}
    cities = append(cities, "Casablanca")
    fmt.Println(cities)
}
```

You can append more than one entry to a slice:


```go
package amin

import "fmt"

func main() {

    cities := []string{}
    cities = append(cities, "Agadir", "Rabat")
    fmt.Println(cities)
}
```

And you can also append a slice to another using an ellipsis:

```go
package main

import "fmt"

func main() {
	cities := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Santa Monica", "Venice"}
	cities = append(cities, otherCities...)
	fmt.Printf("%q", cities)
	// ["San Diego" "Mountain View" "Santa Monica" "Venice"]
}
```


### Length

at any time, you can check the length of a slice by using `len`:

```go

package main

import "fmt"

func main() {
    cities := []string{
        "Casablanca",
        "Rabat",
        "Agadir"
    }
    fmt.Println(len(cities))
    countries := make([]string, 42)
    fmt.Println(len(countries))
}
```

### Nil slice

The zero value of a slice is nil. A nil slice has a length and capacity of 0.

```go

package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }
}
```

### Range

The `range` form of the for loop iterates over a `slice` or a `map`.
Being able to iterate over all the elements of a data structure is very useful and `range`
simplifies the iteration.

```go
package main
import "fmt"
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func main() {
    for i, v := range pow {
        fmt.Println("2**%d = %d\n", i, v)
    }
}
```

You can skip the index or value by assigning to _. If you only want the index, drop the ", value" entirely.

```go
package main
import "fmt"
func main() {
    pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
```


### Break and continue

as if you were using normal loop, you can stop the iteration anytime by using `break`

```go
package main

import "fmt"

func main() {
    pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i)
        if pow[i] >= 16 {
            break
        }
    }
    fmt.Println(pow)
}
```

### Range and maps

range can also be used on maps, in that case, the first parameter is not an incremental integer but the map key.

```go
package main
import "fmt"

func main() {
    cities := map[string] int{
        "New York": 8336697,
        "Los Angeles": 3857799,
        "Chicago": 3857799,
    }
    for key, value := range cities {
        fmt.Printf("%s has %d inhabitants \n", key, value)
    }
}
```

### Maps

Maps are somewhat similar to what other languages call `dictionnaries` or `hashes`.

A map maps keys to values. Here we are mapping string keys (actor names) to an integer value (age).

```go
package main

import "fmt"

func main() {
    celebs := map[string]int{
        "Nicolas cage": 50,
        "Selena gomez": 21,
        "Jude law": 41,
        "Scarlett Johansson": 29,
    }
    fmt.Printf("%#v", celebs)
}
```

When not using map literals like above, maps must be created with make (not new) before use.
The nil map is empty and cannot be assigned to.

Assignement follow the Go convention and can be observed in the exmaple below.

```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex
func main(){ 
    m = make(map[string]Vertex)
    m["Bell labs"] = Vertex{40.68433, -74.39967}
    fmt.Println(m["Bell labs"])
}
```
