---
title: Today I Learned Go Part 7
lang: en-US
tags:
  - go
---

# {{ $page.title }}

## Replace The Current Process With An External Command

Go's `syscall.Exec` function can be used to execute an external program. Instead of forking a child process though, it runs the external command in place of the current process. You need to give the function three pieces of information: the location of the binary, the pieces of the command to be executed, and relevant environment. Here is a simple example.

```go
package main

import "fmt"
import "os"
import "syscall"

func main() {
    // get the system's environment variables
    environment := os.Environ()

    // get a slice of the pieces of the command
    command := []string{"tmux", "new-session", "-s", "burrito"}

    err := syscall.Exec("/usr/local/bin/tmux", command, environment)
    if err != nil {
        fmt.Printf("%v", err)
    }
}
```

When this program is executed, it will replace itself with a new tmux session named *burrito*.


## Seeding Golang's Rand

'Random' numbers in Go don't always seem random. This is because the `rand` package defaults to a seed of 1.

That's great if you need a bunch of random numbers at the start of your program. Not great if you expect a different outcome each time you run the program.

A solution is to seed `rand` with Unix time. Try it in the `init()` function:

```go
package main

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}
...
```

## Sleep For A Duration

Many languages allow you to sleep for a certain number of milliseconds. In those languages, you can give `500` or `1000` to the sleep function to sleep for half a second and a second respectively. In Go, the duration of a call to [`time.Sleep`](https://golang.org/pkg/time/#Sleep) is in nanoseconds. Fortunately, there are constants that make it easy to sleep in
terms of milliseconds.  For example, you can sleep for a half a second (500 milliseconds) like so:

```go
package main

import (
    "time"
)

func main() {
    time.Sleep(500 * time.Millisecond)
}
```

Other available time constants are `Nanosecond`, `Microsecond`, `Second`, `Minute`, `Hour`.
