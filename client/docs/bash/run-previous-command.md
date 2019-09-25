---
title: Run Previous Command
lang: en-US
tags:
  - bash
---

# {{ $page.title }}

Previously run commands can be viewed with the `history` command.

```bash
$ history
10048  git checkout master
10049  gpr
10050  rake
```

With this list, you can rerun any command using `!n`:

```bash
$ !10048
Already on 'master'
```

The command `!!` prints the last command you ran, then runs it. Here is an example:

```bash
$ ls
LICENSE.md README.md  bash       cucumber   rails
$ !!
ls
LICENSE.md README.md  bash       cucumber   rails
```

Replace the second `!` with the first few letters of a command you have previously run, and bash will search for, print, and run the most recent instance.

```bash
$ !rsp
rspec spec/models/user.rb
...
```
