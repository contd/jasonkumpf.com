---
title: Today I Learned Go Part 6
lang: en-US
tags:
  - go
---

# {{ $page.title }}

## Go iota

Go has an interesting feature called `iota`. When declaring a list of constants, this keyword represents successive untyped integer constants.

```go
const (
	foo = iota  // foo == 0
	bar = iota  // bar == 1
	baz = iota  // baz == 2
)
```

Anytime `const` is invoked, the counter resets.

```go
const foo = iota  // foo == 0
const bar = iota  // bar == 0
```

This is a cool way to quickly define a list of integer constants, such as 'true' and 'false', for later use.


## Not So Random

Go's `rand` package makes it easy to generate all sorts of pseudo-random numbers. So if you write a program like so:

```go
package main

import "fmt"
import "math/rand"

func main() {
    stuff := []string{
        "one",
        "two",
        "three",
        "four",
    }
    fmt.Println(stuff[rand.Intn(len(stuff))])
}
```

and then run it, you will get output like:

```
three
```

and any subsequent runs of the program will continue to produce `three`. This is because the default seed for global functions in `math/rand` is [specified](https://golang.org/pkg/math/rand/#Seed) as `1`.

If you want your program to be a little less predictable, you will want to seed it yourself, perhaps with the current time, instead of `1`. Try adding the following to the beginning of the `main` function:

```go
rand.Seed(time.Now().UTC().UnixNano())
```

You'll also want to import the `time` package. Things should *appear* to be a bit more random now.
