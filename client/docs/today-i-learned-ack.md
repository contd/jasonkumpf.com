---
title: Today I Learned Ack
lang: en-US
tags:
    - nix
---

# {{ $page.title }}

This is the first post in a series where I've decided to go through a trove of other people's *Today I Learned* or *TILs* and write up my selections for each topic.

To kick things off I will start with a nice new tool called `ack`.  It's [beyond grep](https://beyondgrep.com/) which so far has lived up to its claim of being better for programmers and very fast.

## ack --bar

The [`ack`](https://beyondgrep.com/) utility has a fun Easter egg that dumps a Star Wars meme to the command line. Give it a try.

```bash
ack --bar
```

See `man ack` for more details.

## Case-Insensitive Search

Use the `-i` flag to perform a case-insensitive search with `ack`.

```bash
ack -i easter

ack/ack-bar.md
3:The [`ack`](https://beyondgrep.com/) utility has a fun Easter egg that dumps

postgres/configure-the-timezone.md
18:Eastern time.
```

If you are a Vim user, you may be familiar with `smart-case`. The `--smart-case` option is a related Ack feature worth checking out.

See `man ack` for more details.

## List Available File Types

The `ack` utility allows you to filter the searched files based on file type. If you'd like to know all of the file types available, you can use the `--help=types` flag. This will include file types you've specified in your `.ackrc` file.

Here is a sample of some of the output.

```bash
ack --help=types
    ...
    --[no]css          .css .less .scss
    --[no]dart         .dart
    --[no]delphi       .pas .int .dfm .nfm .dof .dpk .dproj .groupproj .bdsgroup .bdsproj
    --[no]elisp        .el
    --[no]elixir       .ex .exs
    --[no]erlang       .erl .hrl
    --[no]fortran      .f .f77 .f90 .f95 .f03 .for .ftn .fpp
    --[no]go           .go
    --[no]groovy       .groovy .gtmpl .gpp .grunit .gradle
    --[no]haskell      .hs .lhs
    --[no]hh           .h
    --[no]html         .html .mustache .handlebars .tmpl
    --[no]jade         .jade
    --[no]java         .java .properties
    --[no]js           .js
    ...
```

See `man ack` for more details.

## Search Files Specific to a Language

The `ack` command makes it easy to narrow the set of searched files to those of a specific programming language. For instance, if you have a rails project and only want to search the ruby files, use the `--ruby` flag with your `ack` command.

```bash
ack --ruby Active
```

With the `--ruby` flag, I get a manageable number of results. Without it, not so much.

```bash
ack --ruby Active | wc -l
      26
ack Active | wc -l
    5253
```

See `man ack` for more details.
