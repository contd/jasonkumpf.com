---
title: Today I Learned Go Part 3
lang: en-US
tags:
  - go
---

# {{ $page.title }}

This is the third part of my notes from going through a [go-tooling-workshop](https://github.com/campoy/go-tooling-workshop/).

## `gofmt` versus `goimports`

Both tools `gofmt` and `goimports` will format your code to follow the code styling standard from the Go team at Google.  The `gofmt` tool can be used at the command line and also integrated into your code editor.  Many popular code editors have plugins or packages that bundle this with some other features for making developing in go more productive.

To use `gofmt` for a file and save the result in place you run:

```bash
gofmt -w main.go
```

You can also run it on all files for a package by using the `fmt` subcommand of the `go` tool:

```bash
go fmt .
```

And this will output all the files that were modified.

`goimports` will do the same as `gofmt` but it also will remove any import statements that aren't used and add missing import statements.  For example if a line of your code uses `fmt.Println` and your imports for that file do not include `"fmt"` then it will be added by `goimports`.

## Using `gomodifytags`

This is a handy tool that can take the pain and potential for error out of adding field tags to your structs.  For example, give the following struct in some code in your package:

```go
type Person struct {
	Name string
	Age		int
}
```
You can use `gomodifytags` to add them for you using the following command:

```bash
gomodifytags -w -file main.go -add-tags json -struct Person
```
Which will make the changes and write them to the file, and look like:

```go
type Person struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
}
```
