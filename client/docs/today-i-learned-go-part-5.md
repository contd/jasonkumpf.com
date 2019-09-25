---
title: Today I Learned Go Part 5
lang: en-US
tags:
  - go
---

# {{ $page.title }}

## Concurrency with Go Routines and Channels

The example code here comes straight from this article which insipred this post.  I understood conceptually how go does concurrency and how channels help with communiction.  The following example code, I think, does a great job illustrating this simply.  To start, here is the code to accomplish the same result but without concurrency:

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generateKey() int {
    fmt.Println("Generating key")
    // Super-secret algorithm!
    keys := []int{3, 5, 7, 11}
    key := keys[rand.Intn(len(keys))]
    // It's kinda slow!
    time.Sleep(3 * time.Second)
    fmt.Println("Done generating")
    return key
}

func main() {
    rand.Seed(time.Now().Unix())
    // Call generateKey 3 times.
    for i := 0; i < 3; i++ {
        fmt.Println(generateKey())
    }
    fmt.Println("All done!")
}
```

This takes aproximatly 9 seconds to complete.  Using go routines with channels, the code is rewritten as:

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Update this function to accept a channel parameter,
// and remove the return value.
func generateKey(channel chan int) {
    fmt.Println("Generating key")
    keys := []int{3, 5, 7, 11}
    key := keys[rand.Intn(len(keys))]
    time.Sleep(3 * time.Second)
    fmt.Println("Done generating")
    // Write the key to the channel instead of returning.
    channel <- key
}

func main() {
    rand.Seed(time.Now().Unix())
    // Create a channel.
    channel := make(chan int)
    // Create 3 more goroutines.
    for i := 0; i < 3; i++ {
        go generateKey(channel)
    }
    // Read and print keys from the channel.
    // This also causes the program to wait until 3
    // keys have been read.
    for i := 0; i < 3; i++ {
        fmt.Println(<-channel)
    }
    fmt.Println("All done!")
}
```

Now this takes just over 3 seconds, much better!!
