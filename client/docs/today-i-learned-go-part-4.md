---
title: Today I Learned Go Part 4
lang: en-US
tags:
  - go
---

# {{ $page.title }}

This is the fourth part of my notes from going through a [go-tooling-workshop](https://github.com/campoy/go-tooling-workshop/).

## Testing and Code Coverage

After doing quite a bit of writing test code I found that there are some additional flags or parameters to the `go test` command that are helpful.  I'll use an example from some real code I wrote to illustrate these.  I have a project I created (github.com/contd/links) with the following structure:

```bash
├── app.go
├── main.go
├── main_test.go
├── model.go
├── README.md
```
Testing this project by simply running the following (with -v for verbose output) in the project directory:

```bash
go test -v .
```
And the output looks like:

```bash
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.02s)
=== RUN   TestGetNonExistentLink
--- PASS: TestGetNonExistentLink (0.01s)
=== RUN   TestCreateLink
--- PASS: TestCreateLink (0.03s)
=== RUN   TestGetLinks
--- PASS: TestGetLinks (0.05s)
=== RUN   TestGetLink
--- PASS: TestGetLink (0.04s)
=== RUN   TestUpdateLink
--- PASS: TestUpdateLink (0.05s)
=== RUN   TestDeleteLink
--- PASS: TestDeleteLink (0.05s)
PASS
ok  	github.com/contd/links	0.256s
```

The `main_test.go` is my test file that I've already run and all tests pass using the `go test` command.  But getting code coverage can simply be done with the following:

```bash
go test -cover .
ok  	github.com/contd/links	0.245s	coverage: 67.6% of statements
```
Also to get more detail about the coverage you can use the `-coverprofile` flag to generate a coverage file.  The basic version will just show how many times and line was executed but using the `-html` flag will actually open a web browser with a more friendly visual representation of the coverage.  Here are the two commands and a snapshot of part of the web page coverage profile:

```bash
go test -coverprofile coverage.out .
go tool cover  -html=coverage.out
```
![](./screenshot-001.png)

To not only see *whether* a line of code was executed but also how often it was executed the additional `-covermode=count` will add that to the `coverage.out`file and then the html coverage fill add coloring from dark green to bright green for least to most executed lines of code:

```bash
go test -covermode=count -coverprofile coverage.out .
go tool cover  -html=coverage.out
```

![](./screenshot-002.png)
