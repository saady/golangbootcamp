### Control Flow

### If statement

if statement example:

```go
if answer != 42 {
	return "Wrong answer"
}
```

if with a short statement

```go
if err := foo(); err != nil {
	panic(err)
}
```

### For loop

for loop simple example

```go
sum := 0
for i := 0; i < 10; i++ {
	sum += 1
}
```

for loop without pre/post statement


```go
sum := 1

for ; sum < 1000; {
	sum += sum
}
```

for loop as while loop


```go
sum := 1
for sum < 1000 {
	sum += sum
}
```

### Switch case statement


```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	 
	}
	
}
```

You can execute all the following statements after a match using the `fallthrough` statement:

```go
package main

import (
	"fmt"
)

func main() {
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	    fallthrough
    case 1:
    	fmt.Println("is <= 1")
	    fallthrough
	    
    case 2:
        fmt.Println("is <= 2")
        fallthrough
    case 3:
        fmt.Println("is <= 3")
        fallthrough
    case 4:
        fmt.Println("is <= 4")
        fallthrough
    case 5:
        fmt.Println("is <= 5")
        fallthrough
    case 6:
        fmt.Println("is <= 6")
        fallthrough
    case 7:
        fmt.Println("is <= 7")
        fallthrough
    case 8:
        fmt.Println("is <= 8")
        fallthrough
    default:
        fmt.Println("Try again!")
	}
}
```