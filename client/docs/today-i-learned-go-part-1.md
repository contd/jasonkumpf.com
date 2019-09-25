---
title: Today I Learned Go Part 1
lang: en-US
tags:
  - go
---

# {{ $page.title }}

This post is mostly my notes from going through a [go-tooling-workshop](https://github.com/campoy/go-tooling-workshop/).  This set of tutorials I found to be very helpful even though most of it I knew.  Its has a good babalnce of explanation and brevity.  Here I just wanted to put down some notes on things I had fogotten about or didn't know mostly  for my reference.

## Advanced `go get`

When using `go get` these are some helpful flags:

- `go get -d`: down the code, but do not compile anything.
- `go get -u`: download the latest version even if its already in you `GOPATH`.
- `go get -v`: enable verbose mode.

Also, if a url contains more than one package in sub directories of that url, instead of running `go get` one-by-one you can install all of them in one command like so:

```bash
go get github.com/campoy/go-web-workshop/...
```

## Using `go list`

`go list` allows you to obtain information about your workspace and the packages stored in it.

**NOTE:** in order to list all the packages in the standard library you can simply run `go list std`.

```bash
go list github.com/golang/example/...    # remember that ... means "and everything below"
github.com/golang/example/appengine-hello
github.com/golang/example/gotypes
github.com/golang/example/gotypes/defsuses
github.com/golang/example/gotypes/doc
github.com/golang/example/gotypes/hello
github.com/golang/example/gotypes/hugeparam
github.com/golang/example/gotypes/implements
github.com/golang/example/gotypes/lookup
github.com/golang/example/gotypes/nilfunc
github.com/golang/example/gotypes/pkginfo
github.com/golang/example/gotypes/skeleton
github.com/golang/example/gotypes/typeandvalue
github.com/golang/example/hello
github.com/golang/example/outyet
github.com/golang/example/stringutil
github.com/golang/example/template
```

## Advanced Using `go list`

List and count how many `.go` files are under `github.com/olang/example` using `go list`:

```bash
go list -f '{{join .GoFiles "\n"}}' github.com/golang/example/...
app.go
weave.go
main.go
main.go
hello.go
main.go
main.go
lookup.go
main.go
main.go
main.go
main.go
hello.go
main.go
reverse.go
main.go
```
Now to get the count, you only need to pipe the above output to `wc -l`:

```bash
go list -f '{{join .GoFiles "\n"}}' github.com/golang/example/... | wc -l
16
```
