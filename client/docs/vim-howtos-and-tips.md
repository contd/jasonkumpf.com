---
title: Vim How-Tos and Tips
lang: en-US
tags:
    - vim
---

# {{ $page.title }}

This is more of a scratchpad and reference for me on various things I've found useful at some point and wanted to organize it all in one place.  Most are just composed of pieces copied and modified a little from where I've found them.

- [Absolute and Relative Line Numbers](#absolute-and-relative-line-numbers)
- [Add a File without Loading It](#add-a-file-without-loading-it)
- [Add Custom Dictionary](#add-custom-dictionary)
- [Backspace Options](#backspace-options)
- [Beginning and End of Previous Change](#beginning-and-end-of-previous-change)
- [Breaking the Undo Sequence](#breaking-the-undo-sequence)
- [Buffer Time Travel](#buffer-time-travel)
- [Call a Vimscript Method in Vim](#call-a-vimscript-method-in-vim)
- [Center The Cursor](#center-the-cursor)
- [Change Inner Tag Block](#change-inner-tag-block)
- [Check Your Current Color Scheme](#check-your-current-color-scheme)
- [Close a File](#close-a-file)
- [Close All Other Splits](#close-all-other-splits)
- [Close All Other Windows](#close-all-other-windows)
- [Close the Current Buffer](#close-the-current-buffer)
- [Coerce The Current Filetype](#coerce-the-current-filetype)
- [Count Links in a Markdown File](#count-links-in-a-markdown-file)
- [Count the Number of Matches](#count-the-number-of-matches)
- [Create a New Directory in netrw](#create-a-new-directory-in-netrw)
- [Create a New File in a New Directory](#create-a-new-file-in-a-new-directory)
- [Current Value of a Setting](#current-value-of-a-setting)
- [Default netrw to Tree Liststyle](#default-netrw-to-tree-liststyle)
- [Delete a Line From Another Line](#delete-a-line-from-another-line)
- [Delete Comments](#delete-comments)
- [Delete Every Other Line](#delete-every-other-line)
- [Delete Lines that Match a Pattern](#delete-lines-that-match-a-pattern)
- [Delete to the End of the Line](#delete-to-the-end-of-the-line)
- [Deleting Directories Of Files From netrw](#deleting-directories-of-files-from-netrw)
- [Difference Between :wq and :x](#difference-between-wq-and-x)
- [Display Word Count Stats](#display-word-count-stats)
- [Edges of the Selection](#edges-of-the-selection)
- [Edit the Current File Always](#edit-the-current-file-always)
- [End of the Word](#end-of-the-word)
- [Explore Buffers with BufExplorer](#explore-buffers-with-bufexplorer)
- [Filter Lines Through an External Program](#filter-lines-through-an-external-program)
- [Find and Replace Across Files](#find-and-replace-across-files)
- [Format Long Lines to Text Width](#format-long-lines-to-text-width)
- [Get the pid of the Session](#get-the-pid-of-the-session)
- [Grepping Through the vim Help Files](#grepping-through-the-vim-help-files)
- [Head of File Name](#head-of-file-name)
- [Help For Non-Normal Mode Features](#help-for-non-normal-mode-features)
- [Highlighting Search Matches](#highlighting-search-matches)
- [Horizontal to Vertical and Back Again](#horizontal-to-vertical-and-back-again)
- [Increment All the Numbers](#increment-all-the-numbers)
- [Incremental Searching](#incremental-searching)
- [Increment and Decrement Numbers](#increment-and-decrement-numbers)
- [Interact with the Alternate File](#interact-with-the-alternate-file)
- [Jump to Matching Pair](#jump-to-matching-pair)
- [Jump to the First Non-Blank Character](#jump-to-the-first-non-blank-character)
- [Jump to the Next Misspelling](#jump-to-the-next-misspelling)
- [List All Buffers](#list-all-buffers)
- [List of Plugins](#list-of-plugins)
- [Load a Directory of Files Into the Buffer List](#load-a-directory-of-files-into-the-buffer-list)
- [Marks Across Vim Sessions](#marks-across-vim-sessions)
- [Match the Beginning and End of Words](#match-the-beginning-and-end-of-words)
- [Moving to a Specific Line](#moving-to-a-specific-line)
- [Navigating by Blank Lines](#navigating-by-blank-lines)
- [NETRW Listing Styles](#netrw-listing-styles)
- [Next Modified Buffer](#next-modified-buffer)
- [Open an Unnamed Buffer](#open-an-unnamed-buffer)
- [Open a Tag in a Split Window](#open-a-tag-in-a-split-window)
- [Opening a URL](#opening-a-url)
- [Opening man Pages in Vim](#opening-man-pages-in-vim)
- [Open Vim to a Tag Definition](#open-vim-to-a-tag-definition)
- [Override Vim's Filetype](#override-vims-filetype)
- [Paste a Register From Insert Mode](#paste-a-register-from-insert-mode)
- [Preventing Typos with Abbreviations](#preventing-typos-with-abbreviations)
- [Previous Buffer](#previous-buffer)
- [Previous Visual Selection](#previous-visual-selection)
- [Print Version Information](#print-version-information)
- [Quick man Pages](#quick-man-pages)
- [Re-indenting Your Code](#re-indenting-your-code)
- [Rename Current File](#rename-current-file)
- [Repeating Characters](#repeating-characters)
- [Repeat the Previous Change](#repeat-the-previous-change)
- [Replace a Character](#replace-a-character)
- [Reset Target tslime Pane](#reset-target-tslime-pane)
- [Reverse A Group Of Lines](#reverse-a-group-of-lines)
- [Rotate Everything By 13 Letters](#rotate-everything-by-13-letters)
- [Scrolling Relative to the Cursor](#scrolling-relative-to-the-cursor)
- [Searching For Hex Digits](#searching-for-hex-digits)
- [Set End of Line Markers](#set-end-of-line-markers)
- [Setting Filetype with Modelines](#setting-filetype-with-modelines)
- [Set Your Color Scheme](#set-your-color-scheme)
- [Show All Syntax Highlighting Rules](#show-all-syntax-highlighting-rules)
- [Show Matching Entries For Help](#show-matching-entries-for-help)
- [Sort Alphabetically](#sort-alphabetically)
- [Split Different](#split-different)
- [Splitting For New Files](#splitting-for-new-files)
- [Swap Occurrences of Two Words](#swap-occurrences-of-two-words)
- [Swapping Split Windows](#swapping-split-windows)
- [Tabs to Spaces](#tabs-to-spaces)
- [The Black Hole Register](#the-black-hole-register)
- [The Vim Info File](#the-vim-info-file)
- [Tmux Copy Mode](#tmux-copy-mode)
- [Toggle Absolute and Relative Paths in BufExplorer](#toggle-absolute-and-relative-paths-in-bufexplorer)
- [Toggling Syntax Highlighting](#toggling-syntax-highlighting)
- [Turning Off Search Highlighting](#turning-off-search-highlighting)
- [Unloading a Buffer](#unloading-a-buffer)
- [Use Active Window with BufExplorer](#use-active-window-with-bufexplorer)
- [Verbose Commits with Fugitive](#verbose-commits-with-fugitive)
- [View Commit History of a File](#view-commit-history-of-a-file)
- [Viewing Man Pages with man.vim](#viewing-man-pages-with-manvim)
- [Vim without the Extras](#vim-without-the-extras)
- [What is on the Runtime Path?](#what-is-on-the-runtime-path)
- [Whole Line Auto-Completion](#whole-line-auto-completion)
- [Wrap with Some Room](#wrap-with-some-room)
- [Vim Plugins](#vim-plugins)
	- [Amend Commits with Fugitive](#amend-commits-with-fugitive)
	- [Blank Lines Above and Below](#blank-lines-above-and-below)
	- [Build and Install a Go Program](#build-and-install-a-go-program)
	- [Case-Aware Substitution with vim-abolish](#case-aware-substitution-with-vim-abolish)
	- [Coercing Casing With vim-abolish](#coercing-casing-with-vim-abolish)
	- [Creating Non-Existent Directories](#creating-non-existent-directories)
	- [Deleting Buffers in BufExplorer](#deleting-buffers-in-bufexplorer)
	- [Interactive Buffer List](#interactive-buffer-list)
	- [Quick Quickfix List Navigation](#quick-quickfix-list-navigation)


## Absolute and Relative Line Numbers

By default, vim uses absolute line numbering. This can be turned off with `set nonumber` or more concisely `set nonu`. Turn it back on with `set nu`. Get more details at `:h number`.

Vim also supports relative line numbers. If you'd rather use relative line numbers, first turn off absolute line numbers (`set nonu`) and then turn on relative line numbers with `set relativenumber`. Shave off some characters with `set rnu`. As you might expect, you can turn off relative numbering with `set nornu`.

See `:h relativenumber` for more details.

## Add a File without Loading It

Generally, when you interact with files (e.g. `:e some-file.txt`), you are both adding it to the buffer list and loading the contents of the file as a separate buffer. The `:bad` command allows you to add a file to the buffer list without loading it. For instance, you can add your `README.md` to the buffer list and leave the current buffer in focus with:

    :bad README.md

This command seems particularly useful for scripting the setup of an initial vim environment or preparing for a `:bufdo` command.

## Add Custom Dictionary

When editing a file with `spell` turned on, you may find vim highlighting some of your content in red. This red highlighting indicates a misspelling. Sure, these words technically aren't going to show up in something like the
Merriam-Webster dictionary, but as far as you are concerned, they are words. They are part of your internal, shared language. The word *admin* is a great example. Why not just tell vim once and for all that such words are valid.

You can do just that by moving your cursor over the *misspelled* word and hitting `zg`. That will add it to the `spellfile`, a list of words, which vim includes in its spell checking. This means vim will no longer highlight that word red.


## Backspace Options

The `backspace` option determines the behavior of pressing the backspace key (`<BS>`). By default, Vim's `backspace` option is set to an empty list. There are three values that can be added that each independently alter the behavior of the backspace key. These are `indent`, `eol`, and `start`.

When `indent` is included, you can backspace over indentation from `autoindent`. Without it, Vim will not allow you to backspace over indentation.

When `eol` is included, you can backspace over an end of line (eol) character. If the cursor is at the first position of a line and you hit backspace, it will essentially be joined with the line above it. Without `eol`, this won't happen.

When `start` is included, you can backspace past the position where you started Insert mode. Without `start`, you can enter Insert mode, type a bit, and then when backspacing, only delete back as far as the start of Insert mode.

The `backspace` default is absurd, you are going to want to add all of the above to your Vim settings.

See `:h 'backspace'` for more details.

## Beginning and End of Previous Change

You can jump to the beginning of the previous change with the `[` mark by hitting `'[` from normal mode. Similarly, you can jump to the end of the previous change with the `]` mark by hitting `']`.

Text that was just pasted is also considered a change. Thus, hitting `'[` and `']` will jump to the beginning and end, respectively, of the text that was just pasted into the buffer.

See `:h '[` and `:h ']` for more details.

## Breaking the Undo Sequence

Generally, the sequence of undo-able actions is segmented by command. When entering Insert mode, everything typed until exiting Insert mode is part of a single undo-able segment. If you are going to be typing in Insert mode for a while though, you may want to break it up a bit. Without leaving Insert mode, hit `ctrl-g u` to mark a break in the sequence of undos.

For example, starting in Normal mode and then typing `iabc<CTRL-G>udef<CTRL-G>ughi<ESC>` will leave the buffer with:

    abcdefghi

Hitting `u` once will leave the buffer with:

    abcdef

Hitting `u` again:

    abc

Hitting `ctrl-r`:

    abcdef

See `:h i_CTRL-G_u` for more details.
## Buffer Time Travel

Vim allows you to go to an *earlier* text state for a buffer with `:earlier`. For instance, if you want to see the state of the buffer from 10 minutes ago:

    :earlier 10m

Similarly, you can move back toward the present text state of the buffer
with `:later`. If 10 minutes earlier was too far, you can come back 5
minutes like so:

    :later 5m

I encountered these in [Nick Nisi's 'Vim +Tmux'](https://www.youtube.com/watch?v=5r6yzFEXajQ) talk.

## Call a Vimscript Method in Vim

To call a Vimscript method in Vim, first source the file.

    :source ~/path/to/vimscript.vim

Next, put the cursor over the method call, and run that line.

    :exec getline(".")


## Center The Cursor

If you've moved the cursor to a particular line towards the top or bottom of the screen, you may want to center the screen on that line. This can be achieved by hitting `zz`. It is a quick way to center the cursor and ensure that you have an equal amount of context on either side of the line of interest.

Take care to not mistake this for `ZZ` which will save and quit the buffer.

See `:h zz` and `:h z.` for more details.

## Change Inner Tag Block

Vim has excellent commands for changing the contents of parentheses (`ci(`), brackets (`ci[`), and squiggly braces (`ci{`) (as well as single quotes, double quotes, backticks, etc.). But what about tags?

Place your cursor anywhere inside the tags, type `cit` in Visual mode, and this:

```haml
<div>Begone</div>
```

Becomes:

```haml
<div></div>
```

`d` works too, but I prefer `c` because it puts you in Insert mode at the end of the opening tag, ready to type.

Check out `:h v_it` for more information.

## Check Your Current Color Scheme

Vim ships with a number of different color schemes. There is also a plethora of color schemes built by the open source community. If you'd like to know what color scheme you are currently using, you can check with

    :echo g:colors_name

So more details at both `:h g:colors_name` and `:h colorscheme`.

## Close a File

In Vim there are many ways to close a file. One of the best is `:x`, also know as `:xit`.

`:x` writes and quits a file when changes have been made, like `:wq`, but with one less keystroke.

**An important distinction:**

`:x` and `:wq` are similar, but not identical commands.

The difference is that `:wq` writes the current file and quits, as long as the file is not read-only and has a name, while `:x` only writes if changes have been made. So, if you were to generate a file with Vim, make no changes, and try to write and quit with `:x`, the file would not be written. Doing the same thing with `:wq` would write the file.

## Close All Other Splits

If you have a couple other splits open, but you only want the window in focus, you may find yourself doing some finger gymnastics to navigate to those other split windows to close out of them. There is an easier way.

The Ex command `:only` will close all other splits, excluding the window that is currently in focus.

See `:h :only` for more details.

## Close All Other Windows

Opening split windows can be useful in a number of circumstances. Eventually though, you are going to want to go back to just one window. Generally when this happens to me, I navigate to each of the other split windows that I don't want and execute `:q`. What I want to do is essentially close all the other split windows except for my current one. Vim provides a single command for doing this. By hitting

    <CTRL>w <CTRL>o

all other windows are closed leaving the current window as the only one on the screen.

If you want this command to be able to work with windows containing modified buffers, you are going to want to have the `hidden` option turned on.

See `:h CTRL-W_CTRL-O` for more details.

## Close the Current Buffer

There are a number of ways in Vim to close the current buffer. Obviously, `:q` will do the trick, but that kills all of your buffers which isn't ideal if you are still editing other files.

If you start digging through the Vim docs, you might come across both `:bd` (`:bdelete`) and `:bw` (`:bwipe`). At surface level, these seem like aliases of each other. Give them both a try and in both cases the current buffer will go away, dropping you into one of the other buffers you have open.

The difference between `:bd` and `:bw` is in the details, namely in the side-effects. The `:bd` command is sort of a soft delete that removes the file from the buffer list (do an `:ls` to check). If you have set marks on that buffer, you'll notice that they are still there (check with `:marks`). You may also notice that that buffer may still appear all over the jump list
(see `:jump`). The `:bd` command is going to leave traces of your file all over the place (which could either be really handy or really annoying depending on what you are doing). The `:bw` command on the other hand is going to *wipe out* all of this stuff, hence its name. The Vim docs for `:bw` warn us to only use it if we know what we are doing.

Something worth noting for both commands is that if the buffer is *dirty* (modified, but unsaved), then they won't work, unless you force them to with `:bd!` or `:bw!`.

## Coerce The Current Filetype

If Vim doesn't recognize the filetype of the currently edited file, I can tell Vim what filetype to use. Consider, for instance, that I have a draft of a markdown file with the name, `documentation.md.draft`. Vim will not recognize this as a markdown file and will, thus, not apply markdown syntax highlighting to that file. I can easily tell Vim to treat this as a markdown file by setting its filetype:

    :set filetype=markdown

Markdown syntax highlighting and other relevant options will now be applied to the current buffer.

See `:h filetype` for more details.

## Count Links in a Markdown File

Today I learned a great way to count the links in a README or link collection. Vim Regex!

    :%s/- \[//n

In Vimspeak:

> On every line, find all of the times where `- [` (the beginning of a Markdown bulleted hyperlink) occurs, count them, and report the number of matches.

See `:help substitute` in Vim for more information.

## Count the Number of Matches

You can use the substitution functionality in vim to count the number of matches for a given search term like so:

    :%s/transaction_id//n

You will see the result in the command tray like so:

    8 matches on 8 lines

If you want to find matches globally (that is, counting multiples per line), you can add the `g` flag:

    :%s/transaction_id//gn

for a response like:

    13 matches on 8 lines

The magic is in the `n` flag which tells vim to report a count of the matches and not actually perform the substitution. See `:h :s_flags` for more details. Also, check out `:h count-items`.
## Create a New Directory in netrw

With `netrw` open (try `vi .`), navigate to the parent directory you want to create a new directory in and then hit `d`. Type the name of the new directory in the provided prompt and then hit enter.

## Create a New File in a New Directory

From within a vim session, if you create a buffer for a new file in a directory that doesn't exist. For instance, let's say that `/features` doesn't exist and the new file is `my_latest_feature_spec.rb`:

    :e spec/features/my_latest_feature_spec.rb

Vim's command line will inform you that this is a buffer for a `[New DIRECTORY]`. If you then make some changes and subsequently try to save the file, Vim will present you with:

    "spec/features/my_latest_feature_spec.rb" E212: Can't open file for writing

This is because the containing directory doesn't exist. You can quickly create that directory with a combination of Vim filename shorthands and shelling out to the `mkdir` command.

    :!mkdir -p %:h

The `%` is shorthand for the qualified path of the current file. The `:h` is a filename modifier that returns the *head of the filename*, that is, it resolves to the path with everything except the name of the file. Thus, this command is essentially resolving to:

    :!mkdir -p spec/features/

Vim will shell out with this command making directories for all non-existent directories in the given path. Now you can happily save your new file.

## Current Value of a Setting

Check the value of any Vim setting by adding a `?` to the end of its name.

    ## Validate spelling is turned off
    :set spell?
    => nospell

    ## Validate incremental search is turned on
    :set incsearch?
    => :incsearch

## Default netrw to Tree Liststyle

The built-in `netrw` plugin is a great way to browse files and directories within a Vim session. `netrw` supports four ways of displaying files and directories. That is, there are four liststyles. You can toggle through these by hitting `i`.

I prefer the tree liststyle, which is not the default. I can set the tree liststyle as the default by adding the following line to my `.vimrc` file.

    let g:netrw_liststyle = 3

Now, every time I visit or revisit a `netrw` window, I'll see everything nicely displayed as a tree.
## Delete a Line From Another Line

Today I was cleaning up a test with an extra empty line at the top of the file, away from my Vim cursor. I wanted to delete it... without moving the cursor.

It seems like Vim almost lets you do this, with `:[range]d`. But it leaves the cursor on the deleted line, which isn't very magical. This is the hack we found:

    :[range]d

Then:

    ''

`''` returns the cursor to the last jump point.

## Delete Comments

`:g/^\s*#/d` will remove comment lines from a file with Vim.

## Delete Every Other Line

You can delete every other line in the current buffer using the following command.

There is a fairly elegant way in vim to delete every other line in the current buffer. Why would you want to do that? I don't know. Nevertheless, here it is:

    :g/^/+d

This will essentially delete all even numbered lines. If you'd like to delete all odd numbered lines, delete the first line in the file (`ggdd`) and then run the same command as above.

This syntax is a bit awkward, so you may be better off going straight for a macro (e.g. `qqjddq5@q` or `qqddjq5@q`).

[source](http://stackoverflow.com/questions/1946738/vim-how-to-delete-every-second-row)

## Delete Lines that Match a Pattern

The `:g` command can be used to execute an Ex command over the entire buffer for all lines that match a given pattern. By choosing `d` (delete) as the Ex command, all lines that match the given pattern will be deleted. For instance, if I want to remove all lines that contain `binding.pry`, I can execute the following command:

    :g/binding\.pry/d

See `:h :g` for more details.

## Delete to the End of the Line

There are a number of ways to delete from the cursor position to the end of the line. Generally when I am doing this, I want delete to the end of the line and then start typing something different. Perhaps the best way to do this is by hitting `C`. It deletes to the end of the line and then leaves you in insert mode. This also makes for easier repetition with the dot command.

This is synonymous with hitting `c$`.

See `:h C` for more details.

## Deleting Directories Of Files From netrw

In `netrw`, you can delete files and directories by navigating over the target of deletion and hitting `D`.

By default, `netrw` will use `rmdir` when deleting directories. This means that if a directory has files in it, then it won't be deleted. `rmdir` rightly gives an error when the target directory isn't empty.

Not to worry though, `netrw` can be configured to use `rm -r` instead of `rmdir` when deleting directories.

    :let g:netrw_localrmdir='rm -r'

[source](https://gist.github.com/KevinSjoberg/5068370)
## Difference Between :wq and :x

The `:wq` command is used in Vim to write and quit. The contents of the buffer are written to disk for the associated file and then the Vim session is terminated. So, what is the difference between this and the `:x` command. The Vim help files give the following description of the `:x` command:

> Like ":wq", but write only when changes have been made.

So, `:wq` writes the buffer to disk either way, whereas `:x` just exits if the buffer hasn't changed. Either way the contents of the resulting file are going to be the same. So what's the difference? Modification time.

If you `:x` a buffer that hasn't changed, the modification time will be untouched because the file isn't *re-saved*. The `:wq` command will alter the modification time no matter what.

This matters if the modification time is used by anything. For instance, a background process that monitors a directory for changed files based on modification times will get some false positives if you use `:wq` too liberally.

[source](http://docstore.mik.ua/orelly/unix3/vi/ch05_03.htm)

## Display Word Count Stats

You can display counts for the current file including line count, word count, and byte count by hitting `g CTRL-g`. This also displays the line, word, and byte that your cursor is currently at. The output looks something like the following:

    Col 1 of 0; Line 108 of 337; Word 397 of 1451; Byte 4571 of 18077

See `:h 12.5` for more details.

## Edges of the Selection

When you make a visual selection, Vim stores the position of the first character of the selection in the `<` mark and the position of the last character of the selection in the `>` mark.

Thus moving to the edges of your previous selection is easy. To move to the beginning of the selection, press

    `<

To move to the end, press

    `>

## Edit the Current File Always

The Vim command `:e` edits a file. Add a bang, `:e!`, to edit the current file 'always', discarding any changes to the current buffer.

This is useful if you rename a file with a new file extension (i.e. `.txt` to `.rb`). `:e!` reloads the file in Vim, picking up any customizations you have in your configuration files such as syntax highlighting.

## End of the Word

Word-based movement can serve as a quick way to get around locally in Vim. I most often use `w` and `b` for this kind of movement. While `w` and `b` move me to the beginning of the next and previous word, respectively, I find that
sometimes it would be more convenient if I were at the end of a word.

The `e` and `ge` commands serve this purpose well. `e` will move me to the end of the next word and `ge` will move me to the end of the previous word.

## Explore Buffers with BufExplorer

I'm a huge fan of Vim buffers. I try to have my buffer list roughly mirror the files I am currently holding in my brain.

The BufExplorer Vim plugin helps, and is included in the Hashrocket Dotmatrix. Today I learned to use the command `\bs`, which opens a colorful buffer list that you can navigate with Vim directions.

[https://github.com/jlanzarotta/bufexplorer](https://github.com/jlanzarotta/bufexplorer)
## Filter Lines Through an External Program

Vim allows you to filter lines from your current buffer through an external program. For instance, if you have some ugly looking json that you'd like to format in a readable way, you might want to filter it through an external json pretty printer program.

To filter the entire file through an external program, use

    :%!! <external-program>

Or you can make a visual selection and just filter that

    :'<,'>!! <external-program>

See `:h !!` for more details.

## Find and Replace Across Files

Vim can find and replace strings across files, just like other text editors. It's really (sort of) easy.

First load all the files you want to change into the buffer with a splatted directory.

    :args path/to/files/*/*

Then, execute the substitution.

    :argdo %s/old_string/new_string/ge | update

The `e` flag is important; it tells Vim not to issue an error if the search pattern fails. This will prevent `No Match` errors from breaking the chain.

## Format Long Lines to Text Width

Vim allows you to set the maximum width for text per line in a buffer with the `textwidth` setting. For example, you can set the maximum text width to 80 characters like so:

    :set textwidth=80

With this set, vim will automatically break on whitespace whenever you hit 80 characters. However, there are two places where this doesn't quite pan out. You will see this as soon as you open a file with lines of text that exceed 80 characters and when you paste long lines of text into your buffer. You can quickly remedy this with `gw`.

Make a visual selection of the lines that need formatting and then hit `gw`. All the lines should then we truncated to 80 or less characters.

See `:h gw` and `:h gq` for more details.## Get help with Pathogen

Today I found a great command in Pathogen.vim, `:Helptags`. This is a riff off `:helptags {dir}`, a Vim feature that builds helptags for directory `{dir}`.

`:Helptags` does the same for all the directories in your `runtimepath`, which defaults on Unix and Mac OS X to:

    "$HOME/.vim,
    $VIM/vimfiles,
    $VIMRUNTIME,
    $VIM/vimfiles/after,
    $HOME/.vim/after"

The use case here is when you've just loaded a new plugin into your Vim bundle, you open up Vim, and `:h {my_new_plugin}` isn't 'helping' you out.

## Get the pid of the Session

Your current Vim session is a process running on your machine. That means that this session is tied to a particular process id, _pid_.

Want to know what the pid is?

    :echo getpid()

This will echo the pid of your current Vim session. See `:h getpid()` for more details.

## Grepping Through the vim Help Files

Trying to look up the help file for a Vim feature, but you cannot quite remember the right keyword? Use `:helpgrep`. With `:helpgrep`, you can search across all of the Vim help files for not just the specific keywords, but any pattern of text. For instance, if you want to find where `substitution` is mentioned in the help files, try:

    :helpgrep substitution

It makes a list of all occurrences in the quick fix window and then opens up a split with the cursor on the line of the first occurrence. You can then hit `:copen` to see the rest of the entries.

See `:h helpgrep` for more details.
## Head of File Name

At Hashrocket, I kept seeing my coworkers type a variety of commands into vim command mode that included `%:h`. I finally decided to ask what was going on. It turns out that it produces the directory of the file in your current vim buffer.

The `%` represents the current file and `:h` is a filename modifier, *head of the filename*, that truncates the last component and any separators. So if you remove the file part of the current file (`%`), you are left with the (relative) directory of the current file. Your imagination and vim's flexibility can now take over.

A common use case is to use it to quickly edit another file that you know is in the same directory. Why type out a long pathname over and over throughout the day, when you can type:

    :e %:h<tab>

After hitting tab, the pathname will be auto-completed. Complete the rest of the filename as you do.

Or perhaps you aren't sure what file you want to edit and you'd rather just get a picture of the whole directory:

    :e %:h

You are now exploring the whole directory in netrw mode. Yay!  If you want to find out more about similar features, there is a section in the Vim documentation that talks all about [filename modifiers](http://vimdoc.sourceforge.net/htmldoc/cmdline.html#filename-modifiers).

## Help For Non-Normal Mode Features

The majority of your time in vim will be spent in normal mode. You will often look to the documentation for help on any number of commands and bindings available in normal mode. For instance, to find out more about *goto file*, you may do `:h gf`. And if you want to read more about *yanking* lines of code, you may do `:h y`.

But what about commands and bindings that aren't found in normal mode? What if you want to read about *yanking* text from visual mode? What if you want to get more details on insert's x-mode? Doing `:h y` and `:h ctrl-x`, respectively, won't do the trick because vim thinks you are talking about normal mode bindings.

The docs for these and other non-normal mode features can be found by prepending `i_` and `v_` to the binding in question.

To get at the docs for *yanking* from visual mode:

    :h v_y

And to read up on insert's x-mode:

    :h i_ctrl-x

## Highlighting Search Matches

Want to see all the places in the buffer that match a search pattern? Turn on `hlsearch` and Vim will highlight all the matches of the previous search.

Try turning it on with `:set hlsearch` and then search for some pattern using `/`.

If you no longer want to see all the highlighted matches, turn it off with `:set nohlsearch`.

See `:h hlsearch` for more details.

## Horizontal to Vertical and Back Again

If you have two Vim windows open with a horizontal split and you want to quickly switch them to be vertically split, you can type the following two key bindings:

    Ctrl-w t
    Ctrl-w H

To go from vertically split windows to horizontally split windows you can instead use:

    Ctrl-w t
    Ctrl-w K

[source](http://stackoverflow.com/questions/1269603/to-switch-from-vertical-split-to-horizontal-split-fast-in-vim)
## Increment All the Numbers

Vim's substitution feature can be used for more than simple static text replacement. Each replacement can be the result of some operation on the original text. For instance, what if we'd like to increment all numbers in the buffer? We can achieve this by searching for all numbers and then using `\=` with `submatch`. Whenever the replacement string of a substitution starts with `\=`, the remainder of the string is evaluated as an expression.

Given the following text in our buffer:

    1 2 a b c 45 123 1982

We can run the following substitution command:

    :%s/\d\+/\=submatch(0)+1/g

This will transform all digits in the buffer, resulting in:

    2 3 a b c 46 124 1983

Want to decrement all the numbers instead?

    :%s/\d\+/\=submatch(0)-1/g

See `:h sub-replace-expression` for more details.

## Incremental Searching

You can do a text-based search on the contents of your current buffer by hitting `/` and then beginning to type a pattern (including regex). The `incsearch` feature makes searching for text even easier. As you type your pattern, the first valid match will be located and highlighted. As you continue to type the pattern, it will continue to update the highlighted
match. Incremental searching makes it easy to see when you've made a typo in your pattern. By default `incsearch` is turned off in Vim. You can enable it with `:set incsearch`.

See `:h incsearch` for more details.

## Increment and Decrement Numbers

Vim has commands built-in to increment and decrement numbers: `CTRL-A` and `CTRL-X`, respectively.

Combining this with a macro was a technique we tried today to build a large database migration. We ended up finding a more efficient solution, but it's still a very magical key combination.

Check out `:help CTRL-A` for more info.

## Interact with the Alternate File

If you have a couple buffers going in a Vim session and you check out the buffer list with `:ls`, you'll notice that one of those buffers has a `#` indicator next to it. That means the file for this buffer is considered the alternate file of the current, visible buffer. In addition to hitting `CTRL-^` to switch to that buffer, you can reference it in other commands with `#`. This means you can quickly `:edit`, `:split`, `:vsplit`, and so forth the alternate file by just giving `#` as the argument to those commands.

Quickly open the alternate file in a vertical split with:

    :vsp #

See `:h alternate-file` for more details.## Joining Lines Together

You can quickly join a series of lines onto a single line using the `J` command. Simply hitting 'J' in normal mode will join the current line with the line below it leaving a space in between the two. If you want to join the next 5 lines, you can hit `5J`. Hitting '5J' on this:

    first,
    second,
    third,
    fourth,
    fifth

will transform it into this:

    first, second, third, fourth, fifth

Similarly, you can perform a line by line join on a visual selection by hitting `J` after making the desired visual selection.

## Jump to Matching Pair

If you are dealing with code or data that contains parentheses or brackets that are hard to follow, you can easily jump between them with `%`.

For example, if you move over a `[` and hit `%`, Vim will jump your cursor to the matching `]`. Hit `%` one more time and it will jump back.

See `:h %` for more details.

## Jump to the First Non-Blank Character

With Vim you can jump down any *n* amount of lines with `n + j`, and back up with `n + k`.

An alternate command is `n + +`, which jumps down to the first non-blank character, and `n + -`, which jumps back up to the first non-blank character.

## Jump to the Next Misspelling

If spelling is turned on (`:set spell`), you can jump back and forth between words that are misspelled. To jump to the next misspelling, hit `]s`. To jump to the previous misspelling, hit `[s`.

See `:h ]s` and `:h [s` for more details.
## List All Buffers

The `:ls` command will list the buffers you have open. What vim doesn't tell you though is that there are some unlisted buffers that it isn't displaying. To see *all* of the buffers, you can use `:ls!`. According to the vim help
file:

> When the [!] is included the list will show unlisted buffers (the term "unlisted" is a bit confusing then...).

This reveals buffers for `netrw`, `:help` files, etc. This helps explain the sometimes sporadic numbering that vim uses for buffers.

## List of Plugins

Get a list of all your plugins (and other sourced scripts) with

    :scriptnames

See `:h scriptnames` for more details.

## Load a Directory of Files Into the Buffer List

Consider the scenario where I want to go through all files in a directory to make a series of minor, related changes. After editing each file, I can type something like `:e path/to/next/file.md` to bring up the next file. That can get a bit tedious though. Instead, I can load up all the files in the directory with the `args` command:

    :args path/to/files/*.md

From there, I can use `:bnext` (or `:bn) and `:bprev` to more quickly jump through those files I want to edit.

I can also run `:ls` to see all the files loaded in to the buffer list at that point.

[source](http://superuser.com/questions/396481/how-to-load-multiple-files-in-multiple-subdirectories-into-vim-buffers)
## Marks Across Vim Sessions

Any mark set with a capital letter (that is, `A-Z`) is called a *file mark*. File marks can be used to move from file to file. File marks will also be written to the Vim Info file (`~/.viminfo`) which means that if you close Vim and open it again, it will still know about those file marks. This means that your file marks are persisted across vim sessions which will save you the trouble of having to redefine them each time.

A more detailed description of marks is at `:help marks`.

## Match the Beginning and End of Words

Often when doing a substitution for an exact word, say `user` to `admin`, I will include spaces on either end of the regex pattern to avoid unintentional replacements. For example, I may use something like

    :%s/ user / admin /

in order to avoid a substitution like `username` to `adminname`.

In this case, the spaces can be replaced with zero-width regex characters that match the beginning and end of a word. These are `\<` and `\>`, respectively. Utilizing these, the previous substitution can be achieved with

    :%s/\<user\>/admin/

See `:h /\<` and `:h /\>` for more details.

## Moving to a Specific Line

Often times when I open a new buffer, it is with the intention of moving to a particular line. For example, if I am trying to move to line 55 in the file, then I will hit `55j`\*. This works fine when I am dealing with a freshly opened buffer. That is, this works fine if I am starting from the top of the buffer.

In general, there is a better approach. I can move to an exact line number from normal mode with `:{N}` where `{N}` is the line number. So, if I want to get to line 55 regardless of where I am currently positioned in the buffer, I can simply hit `:55<cr>`.

*  *This actually is slightly inaccurate, it moves me to line 56, not 55. If I need to be precise, `55j` doesn't cut it.*

** *Also, you can apparently use `55G` to achieve the same thing ([source](https://twitter.com/rossnelson/status/591239951032983553)).*
## Navigating by Blank Lines

Use vim to open a file full of code (or text) that has some blank lines. Move the cursor to the middle of the file. Then start hitting `{` or `}`. You'll see that the cursor jumps from blank line to blank line.

Use `{` to jump to the closest blank line _behind_ the cursor. Use `}` to jump to the closest blank line _ahead_ of the cursor.  This may not seem like the most practical or obvious way to navigate, but it can help move you around a bit quicker than just tapping `k` and `j`.

## NETRW Listing Styles

When you edit a directory with vim (`vim .`), you are taken into netrw which allows you to explore the contents of that directory. By default, you will see a list of the files and directories in the target directory, like so:

```
" ============================================================================
" Netrw Directory Listing                                        (netrw v151)
"   /Users/jbranchaud/code/til
"   Sorted by      name
"   Sort sequence: [\/]$,\<core\%(\.\d\+\)\=\>,\.h$,\.c$,\.cpp$,\~\=\*$,*,\.o$,\.obj$,\.info$,\.swp$,\.bak$,\~$
"   Quick Help: <F1>:help  -:go up dir  D:delete  R:rename  s:sort-by  x:special
" ==============================================================================
../
./
.git/
git/
go/
postgres/
rails/
ruby/
vim/
zsh/
LICENSE
README.md
```

This (`thin`) is only one of a few listing styles that you can use to explore a directory. The other styles are `long`, `wide`, and `tree`. You can cycle between them by pressing `i`. For instance, if you cycle through to the `tree` format, you will be presented with navigable tree structure like so:

```
" ============================================================================
" Netrw Directory Listing                                        (netrw v151)
"   /Users/jbranchaud/code/til/git
"   Sorted by      name
"   Sort sequence: [\/]$,\<core\%(\.\d\+\)\=\>,\.h$,\.c$,\.cpp$,\~\=\*$,*,\.o$,\.obj$,\.info$,\.swp$,\.bak$,\~$
"   Quick Help: <F1>:help  -:go up dir  D:delete  R:rename  s:sort-by  x:special
" ==============================================================================
../
til/
| .git/
| git/
| | checkout-previous-branch.md
| | delete-all-untracked-files.md
| | dry-runs-in-git.md
| | intent-to-add.md
| | staging-changes-within-vim.md
| | stashing-untracked-files.md
| | verbose-commit-message.md
| go/
| postgres/
| rails/
| ruby/
| | create-an-array-of-stringed-numbers.md
| | limit-split.md
| | parallel-bundle-install.md
| | summing-collections.md
| vim/
| zsh/
| LICENSE
| README.md
```

[source](http://vimdoc.sourceforge.net/htmldoc/pi_netrw.html)

## Next Modified Buffer

After working for a while on a feature that involves looking at a number of files, I end up with a decent buffer list. I will have inevitably edited a few of those files and occasionally I'll inadvertently leave one of the buffers modified. Instead of opening
the buffer list (`:ls`), finding the modified buffer, and navigating to it, I can just jump straight to it. I can do this with `:bmodified` or just `:bm`. This jumps straight to the next modified buffer. If there is no modified buffer, it tells me *No modified buffer found*.

See `:h bmodified` for more details.
## Open an Unnamed Buffer

There are two ways (that I know of) to open an unnamed buffer.

The first is before vim has even been launched. You can simply execute `vim` from the command-line without any arguments. Follow that by invoking `:ls` to see that the current and only buffer has *no name*.

The second method is with a vim session that is already open. If you invoke `:new`, a new buffer will be created that, like the first method, has *no name*.

## Open a Tag in a Split Window

Using tags and `ctrl-]` is a quick way to jump from the use of a keyword to the declaration or definition whether in the same file or another file. Sometimes what you really want is that tag definition opened up in a (horizontal) split window. That way you can see the definition without losing context. This can be accomplished with `ctrl-w ]`.

The Vim help file gives the following definition of this command:

> Use identifier under cursor as a tag and jump to it in the new upper window.

See `:h CTRL-W_]` for more details.

## Opening a URL

Vim makes it easy to quickly open a URL that appears in a file. Simply move the cursor over the URL and hit `gx`. This will use your respective operating system's *open* command (e.g. `open` for Mac OS X) to open the URL.

One caveat is that the URL must contain the protocol/scheme. That is, `www.duckduckgo.com` won't work, but `https://www.duckduckgo.com` will.

You can also use `gx` to open files on your system.

## Opening man Pages in Vim

In [Quick Man Pages](quick-man-pages.md), I explained how you can quickly open man pages with `K`. For times when the particular command isn't in the buffer or the command contains a hyphen, you can instead use `:Man`. With the `ft-man-plugin` enabled, you can use `:Man` with the name of any command that has a manual page and the respective man page will be opened in a split buffer. For example, check out `git log` with:

    :Man git-log

If you don't want the first manual entry, provide a specific number. For instance, you can open the `echo(3)` man page with:

    :Man 3 echo

See `:h :Man` for more details.

## Open Vim to a Tag Definition

If you are using [ctags with Vim](https://andrew.stwrt.ca/posts/vim-ctags/), you can provide a tag name when opening Vim. This signals to Vim that it should open to the file and location of the tag's definition. For instance, if you have a Rails project going and you provide Vim with the `UsersController` tag name, it will open the `app/controllers/users_controller.rb`. Just use the `-t` flag like so:

```bash
$ vim -t UsersController
```

See `man vim` for more details.

## Override Vim's Filetype

Vim's filetype auto-detection is great for effortless syntax highlighting, but what if a certain kind of file (i.e. Ruby) contains lots of another kind of code (i.e. SQL)? The Ruby code will be highlighted and readable, the SQL a large monochrome blob. Hard to read and reason about. We can do better!

Override the automatic assignment with:

    :set ft=sql

This command with no assignment returns the current setting:

:set ft
=> filetype=lua

We can easily revert to the auto-detected extension as needed.
## Paste a Register From Insert Mode

Generally pasting a register is done from Normal mode using `p` or something like `'1p` (the former pasting from the default register, `0`, and the latter pasting from register `1`). Vim also allows you to paste from a register without leaving Insert mode. By hitting `CTRL-R` and then the name of the register, Vim will insert the contents of the register in front of
the cursor.

For example, to paste from the default register from Insert mode, hit `CTRL-R 0`.  Note, mappings and abbreviations will not be applied to the inserted text.

See `:h i_CTRL-R` for more details.

## Preventing Typos with Abbreviations

Are you are prone to mistyping *the* as *teh* or *function* as *funciton*? You can add one-line abbreviations to your `.vimrc` file to auto-correct these mistakes for you.

    abbr teh the
    abbr funciton function

By adding these (or others) to your vim configuration, whenever you type the misspelled version, vim will know to instantly replace it with the correct version. This can be handy, but use it sparingly.

## Previous Buffer

I often find myself needing to jump back and forth between two buffers. For instance, if I am iterating on a test and the implementation, there is a lot of switching from one to the other and then back again.

This quickest way to do this is to use the command for going to the previous buffer. The default binding for that is `ctrl-^`.

With that binding, it is fast and easy to toggle between two buffers.

[source](http://vimdoc.sourceforge.net/htmldoc/editing.html#CTRL-^)

## Previous Visual Selection

Typing `gv` in normal mode will re-select the previous visual selection.  This makes it easy to re-select a specific block of text. For instance, if you are performing a search and replace on a visual selection and you didn't get the regex quite right, you can quickly type `gv` and then edit the regex of your previous command.

## Print Version Information

Want to know what version of Vim you are using, plus a bunch of other information? Try entering

    :version

This will display the version including patches. It will tell you when it was compiled. A list of available and unavailable features is also included.## Quick File Info

If you are browsing a directory with vim (e.g. `vim .`) and you want to see information about a file such as the *last modification date* and *file size*, move your cursor to that file's name and type `qf`.

## Quick man Pages

Within Vim, if you encounter a command that has man pages (such as `grep`), you can move your cursor over that word in normal mode and press `K` (`shift+k`) to view the man pages for that command.

[source](https://twitter.com/vimchi/status/571438478027837440)

## Re-indenting Your Code

If you have pasted some poorly formatted code or you've written a portion of code in a way that mangled the indentation, you can quickly re-indent that code for you. Or at least do its best to try and indent it correctly.

Make a visual selection of the code that you want to re-indent and then hit `=`.

For instance, this ruby code

```ruby
if this_thing
p something
else
p nothing
end
```

will be indented by Vim as

```ruby
if this_thing
  p something
else
  p nothing
end
```

See `:h =` for more details on how vim decides what formatting and
indenting it will do.

## Rename Current File

Vim doesn't come with an easy way to rename the existing, current file. The closest you will get with a one-off command is `:saveas {new file name}` which renames the current buffer, but also leaves you with your old file.

Another approach that you can take is to chain two commands together. You can start be *deleting* the current file (don't worry, you've still got the file contents in your buffer) and then *saving* the buffer with the new name. Like so:

    :call delete(expand('%')) | saveas new-file-name.txt

While this seems to do the job well enough, your mileage may vary. Consider using a more robust plugin, such as
[vim-eunuch](https://github.com/tpope/vim-eunuch) or
[rename.vim](https://github.com/danro/rename.vim/blob/master/plugin/rename.vim).
## Repeating Characters

It's not common to need to type 39 space characters in a row, but when the moment arises, you don't want to be caught hitting the space bar 39 times. Vim makes it easy. From normal mode, type the number, then `i`, and then the character to be repeated followed by an escape.

For instance, for 39 spaces, hit:

    39i <esc>

or for 80 dashes, hit:

    80i-<esc>

## Repeat the Previous Change

If you have just performed a change and you realize you want to immediately do it again, you can hit `.` instead of retyping it. This is a good way to quickly repeat simple changes.

For instance, if you are adding indentation to a code block with `>`, you can estimate the number of indents that need to happen and type `3>` or you can indent once and then hit `.` for additional indentation until you reach the right level of indentation.

Check out the docs for more details: `:help .`

## Replace a Character

Throughout the day I'll often find myself deleting a single character and putting a different one in its place. I usually navigate over the target character and hit `s` which removes the character under the cursor and puts me in insert mode. From there I type the new character and hit escape to return to normal node. This isn't the best way to perform such an edit though. Vim has a command specifically for replacing a character. The `r` command. It does essentially the same thing as my current approach but instead of putting me in insert mode, it simply replaces the character and leaves me in normal node.

See `:h r` for more details.

## Reset Target tslime Pane

The combination of [`tslime`](https://github.com/jgdavey/tslime.vim) and [`turbux`](https://github.com/jgdavey/vim-turbux) makes running tests from Vim in a tmux session as easy as a single key binding. One problem that can arise from time to time is having `tslime` focused on an undesired tmux window/pane combination. There is no binding to tell `tslime` that you'd like to re-select the target window and pane.

I've often resorted to closing out of Vim in order to reset the prompt. There is a better way and it doesn't require you to wipe out your Vim session.

Just `unlet` the global Vim variable for the `tslime` plugin like so:

    :unlet g:tslime

The next time you invoke `turbux` it will see that `g:tslime` isn't set and will prompt you for a new window and pane combination.

## Reverse A Group Of Lines

The following command can be used to reverse the order of all lines in a file by doing a global move on all lines that match the beginning of line character starting at the theoretical 0th character:

:g/^/m 0

Reversing a range of lines is a little more work. Just as the previous example needs to be anchored against the 0th character, a specific range of lines needs to be anchored at the line just before the range. Thus reversing the lines 5 to 10 requires being anchored at line 4, like so:

:4,10g/^/m 4

See `:h 12.4` for more details on how this works.

[source](http://superuser.com/questions/189947/how-reverse-selected-lines-order-in-vim#)
## Rotate Everything By 13 Letters

For some inane reason, Vim comes with a [ROT-13](https://en.wikipedia.org/wiki/ROT13) feature. So, if you are ever in need of rotating the letters of some portion of the file by 13, you can do that with the simple `g?` binding.

For example, if you hit `g??` on the following line:

    Six dollar eggs

you will get

    Fvk qbyyne rttf

As you can see, casing is preserved.  The only practical uses of this are Vimgolf and convincing people at coffee
shops that you are a hacker.

See `:h g?` for more details.

## Scrolling Relative to the Cursor

If you hit `zt` while in normal mode, the window will be redrawn such that the line the cursor is positioned on is at the top of the window. Similarly, if you hit `zb`, the window will be redrawn such that the line the cursor is currently on will be at the bottom of the window.

The one that comes in really handy, though, is `zz` (note: this is not `ZZ`) which will reposition the current line to the center of the screen. This can come in really handy if you have scrolled to the top (or bottom) of the visible part of the buffer and you want to quickly view more of the context around the current line.

See `:h scroll-cursor` for more details and commands.

## Searching For Hex Digits

If you want to find sequences of hex digits (`A-F`, `a-f` and `0-9`) in a search, you can hack together something like:

    /[A-Fa-f0-9]\+

This is a bit verbose, though. Vim has a number of built in character classes that can be referenced in searches and substitutions. For hex digits, there is `\x`. Using this, the search above can be achieved with:

    /\x\+

See `:h \x` for more details and other character classes.

## Set End of Line Markers

Vim has a number of invisible characters that you can set. One of those characters is the end of line (`eol`) character.
Whatever character you set this to will appear at the end of each line in your file. This is great for highlighting extra whitespace at the end of a line that would otherwise appear invisible.

Set the `eol` invisible character like so

    :set listchars=eol:¬

or append it to the existing list of invisible characters like so

    :set listchars+=eol:¬

See `:h listchars` to see what other invisible characters you can set.
## Setting Filetype with Modelines

Vim and various plugins generally use known file extensions to determine the filetype of a file. This is important because it is how Vim decides which filetype-specific settings to enable, such as syntax highlighting.

If I am editing a file such as `build.boot`, Vim is not going to know that its filetype should be set to `clojure`. The `build.boot` file is full of clojure code though, so I'm losing out on syntax highlighting and so forth. I can settle for manually setting the filetype to clojure (e.g. `:set ft=clojure`) each time I open up the file.

Or I can use a modeline setting. By including a comment at the top or bottom of the file specifying the filetype setting, I can ensure that each time I go to edit the file, the appropriate filetype will be set.

That modeline comment will look something like:

```clojure
; vim: set ft=clojure:
```

See `:h modeline` for more details.

## Set Your Color Scheme

Vim ships with a number of standard color schemes for both light and dark backgrounds. You can also find and install many others. To set your color scheme, you can use the `:colorscheme` command. You can quickly see what color schemes are available by typing `:colorscheme` and then tabbing for completion. For instance, you can set the *delek* color scheme by entering

    :colorscheme delek

See `:h colorscheme` for more details.

## Show All Syntax Highlighting Rules

When Vim adds syntax highlighting to the contents of a buffer based on its `filetype`, it does so with a set of rules. These rules specify particular colors for each set of named tokens that match particular patterns. You can check out the syntax highlighting rules for the current `filetype` of the current buffer by running:

    :syntax

See `:h :syntax` for more details.

## Show Matching Entries For Help

Looking up things with Vim's `:help` command can be error prone if you don't know exactly how to format what you are looking up. Bypass some of the guesswork by hitting `Ctrl-d` after writing part of the `:help` command. This will populate a wildmenu of possible matches.

For instance, if you know there is a command containing `?`, but you aren't sure how to look it up, try the following:

    :help ?<Ctrl-d>

You can tab through to the one you want and hit enter to read up on it. Who knew there were so many Vim bindings involving a `?`.

See `:h help-context` for more details.
## Sort Alphabetically

One way to make a list nicer to read is to sort it. Vim comes with a command built in for just this purpose.

Here is a snippet from a Gemfile:

```
gem 'coffee-rails', '~> 4.1.0'
gem 'uglifier', '>= 1.3.0'
gem 'sass-rails', '~> 5.0'
gem 'puma'
gem 'gravatar_image_tag'
gem 'authem'
gem 'jquery-rails'
gem 'pg'
gem 'redcarpet'
gem 'rails_12factor', group: :production
gem 'sdoc', '~> 0.4.0', group: :doc
```

These gems might have been added during the development process. To sort them, enter visual mode, highlight the desired range, and enter `:sort`. Here's the result:

```
gem 'authem'
gem 'coffee-rails', '~> 4.1.0'
gem 'gravatar_image_tag'
gem 'jquery-rails'
gem 'pg'
gem 'puma'
gem 'rails_12factor', group: :production
gem 'redcarpet'
gem 'sass-rails', '~> 5.0'
gem 'sdoc', '~> 0.4.0', group: :doc
gem 'uglifier', '>= 1.3.0'
```

## Split Different

Vim's defaults for `:split` and `:vsplit` are to open the splits above and to the left, respectively. I prefer for them to the split below and to the right, respectively. To get `:split` to split below:

    :set splitbelow

and to get `:vsplit` to split to the right:

    :set splitright

Add those two lines to your `.vimrc` to have them set that way all the time.  If you have them set as such in `.vimrc` and want to temporarily revert the setting for the current file, type:

    :set nosplitbelow

or

    :set nosplitright

as commands.

## Splitting For New Files

Let's assume you already have a vim session open. You now want to open a new file. You can open a new horizontally split window with the `:new` command. Alternatively, if you'd like the new window to open with a vertical split, you can use the `:vnew` command.

See `:new` and `:vnew` for more details.
## Swap Occurrences of Two Words

Imagine I have a file with `foo` and `bar` all over the place. The tables have turned and now I want all occurrences of `foo` to be `bar` and all occurrences of `bar` to be `foo`.

Reaching for a simple substitution won't work because after flipping all the occurrences of `foo` to `bar`. I will no longer be able to distinguish between the new `bar`s and the `bar`s that need to be flipped.

[Abolish.vim](https://github.com/tpope/vim-abolish) enhances Vim's substitution capabilities making it easy to flip these strings in one relatively simple command.

    :%S/{foo,bar}/{bar,foo}/g

Notice the uppercase `S` as well as the ordering of `foo` and `bar` in the before and after sequences.

## Swapping Split Windows

Consider a scenario where you have a vim window that has been split horizontally into two viewports. You'd prefer the top one to be on bottom and the bottom one to be on top. You want to swap them.

If you are currently focused on the top viewport, then tell vim to move that viewport down with `CTRL-w J`. As you might guess, moving the bottom viewport up can be done with `CTRL-w K`. `J` and `K` mean down and up in other contexts; vim is consistent with their meaning here.

Viewports of a vertical split can be swapped in the same sort of way. Use `CTRL-w L` to move the left viewport to the right. Use `CTRL-w H` to move the right viewport to the left.

## Tabs to Spaces

If you prefer spaces over tabs in your files, then opening up a file full of tabbed indentation is not ideal. You can quickly convert all tabs in the current buffer to spaces using:

    :retab

This assumes that you have `expandtab` set. See `:help :retab` for more details.

## The Black Hole Register

Vim has a variety of registers for storing and moving around text. Vim also has a special register called the *black hole register*. This black hole register is associated with the `_` character.

> When writing to this register, nothing happens.  This can be used to delete text without affecting the normal registers.  When reading from this register, nothing is returned.

As stated in the docs, if you don't want to overwrite the unnamed register or some other register when deleting text, you can use the black hole register.  For instance, deleting the current line without any register side-effects
looks like this:

    "_dd

See `:h registers` for more info on Vim's registers.

## The Vim Info File

Vim serializes a bunch of useful information as you edit files, jump around, and execute commands. This is so that vim can recall this information in between sessions. Vim creates a Vim Info file (`~/.viminfo`) in your home directory which it uses to enhance your long-term experience with the editor. File marks, registers, command and search history, and jump history are some of the more interesting things that vim stores there.

Read more about it at `:help viminfo` or just take a look at the file with `vim ~/.viminfo`.
## Tmux Copy Mode

Last week we encountered a breaking change in our Tmux configuration, and briefly lost the ability to scroll with the mouse. It became an opportunity to re-learn Tmux's copy mode, and liberate ourselves from the mouse a little more.

Enter copy mode with `<prefix> [`. From here, you can navigate with standard Vim commands.

`ctrl-u` moves you a half-page up, perfect for wading through test failure backtraces.

## Toggle Absolute and Relative Paths in BufExplorer

After opening [BufExplorer](https://github.com/jlanzarotta/bufexplorer) using `<leader>bs`, you will see both files and the paths to those files. By default you will see absolute paths to the files. You can use `R` to toggle between relative and absolute paths.

For relative paths, the path will be relative to the current working directory of the Vim session (`:pwd`).

## Toggling Syntax Highlighting

Syntax highlighting in Vim is generally a good thing, but sometimes you need to turn it off. This can be achieved with the `syntax` command.

    :syntax off

When you need that syntax highlighting back again, you can turn it right back on like so:

    :syntax on

See `:h syntax-on` for more details.

## Turning Off Search Highlighting

After performing a search for a common word, you end up with that word highlighted all over the place. To turn it off, I generally use the `set` command with `no` prepended to the `hlsearch` option -- as is convention in Vim.

It turns out though, that `nohlsearch` is a command in its own right. I can save a few characters by invoking:

    :nohlsearch

See `:h nohlsearch` for more details.

## Unloading a Buffer

My preferred workflow with vim involves working across as many buffers as I need in a single window. I open files in new buffers as needed and navigate between existing ones with a number of built-in vim features and plugin shortcuts. Eventually though, my list of buffers gets a bit crowded making it difficult to move around and keep everything straight. One method for dealing with this is occasionally unloading the buffers you no longer need. This can be accomplished with `:bd`.

To unload the current buffer:

    :bd

To unload some other buffer by buffer number, say buffer 10:

    :10bd

Caveats: unloading a buffer marks it as *unlisted* in the buffer list, meaning it won't appear in your normal view of the buffer list. It should also be noted that it does not remove the buffer from the jump list or the global mark list.

See `:h :bd` for more details.
## Use Active Window with BufExplorer

I often use [BufExplorer](https://github.com/jlanzarotta/bufexplorer) within long-running Vim sessions as a way of quickly jumping between the various buffers in my buffer list. After working with one buffer, I use `<leader>bs` to open BufExplorer, move my cursor to the next buffer of interest, and hit enter to open it in place of the current buffer. This is the default behavior at least.

With this setting toggled on, BufExplorer will open buffers in the _active window_. The _active window_ is the window that was active before BufExplorer was opened. If this setting is toggled off, BufExplorer doesn't bother finding the active window, it just opens the buffer up in place of itself in whatever split window was created for itself.

This setting can be toggled within the BufExplorer window by hitting `f`. It will toggle between `Locate Buffer` and `Don't Locate Buffer`. I prefer the default of `Locate Buffer`.

## Verbose Commits with Fugitive

Let's say you are using [fugitive.vim](https://github.com/tpope/vim-fugitive). You've staged some changes within the git index buffer using `:Ge:` and now you want to make a commit. From the git index buffer, you can hit `cvc` to pop open the commit message window in verbose mode. The verbose part means that all the staged changes are shown below as a reference for composing the commit message.

## View Commit History of a File

[Gitv](https://github.com/gregsexton/gitv) is an extension of the [Fugitive](https://github.com/tpope/vim-fugitive) plugin that allows you to view and step through the commit history of a file (among other things).

Open a file in Vim and enter the `:Gitv!` command to open a preview window listing the commits involving the current file. It will look something like this:

```
-- [plugin/fugitive.vim] --
*  (HEAD, r:origin/master, r:origin/HEAD, master) Provide g:fugitive_no_maps to disable key maps  4 weeks ago            Fedor Gusev    [0095769]
*  Support browsing with new netrw.vim                                                            2 weeks ago            Eli Young      [e8b9409]
*  Support for browsing with recent Vim                                                           4 weeks ago            Tim Pope       [eb8eb18]
*  s:Diff: use winnr with `<C-W>w` instead of `<C-W>p`                                            6 weeks ago            Daniel Hahler  [933f6a1]
*  (tag: t:v2.2) fugitive.vim 2.2                                                                 7 weeks ago            Tim Pope       [3471901]
*  Use `<nomodeline>` with Fugitive autocmds, and un-silent them                                  4 months ago           Daniel Hahler  [2c8461d]
*  Make configured_tree a caching global function                                                 4 months ago           John Whitley   [d3b98d9]
*  Fix instaweb support                                                                           4 months ago           Tim Pope       [5699f46]
*  Fix :Glog                                                                                      6 months ago           Tim Pope       [0374322]
```

You can skim over these commits and when one looks interesting, just hit `enter` when your cursor is over its respective line to view the file as it was at the time of that commit.

Restore your buffer to its original state by navigating to the top of the preview and hitting `enter` for the filename on the first line.

This tool is great for use in a team setting when you want to figure out what changes a file has undergone recently, especially when something about a file seems a little fishy. This is also great for individual and team use when you simply cannot remember why you changed a file or what it used to look like before that *clever* refactoring.

## Viewing Man Pages with man.vim

In [Quick Man Pages](quick-man-pages.md), I introduced `K` which shells out to the man page for the unix command under the cursor. It gets better though. Vim has a built-in plugin, `man.vim`, that you can enable which allows you to view man pages without shelling out.

Add the following to your `.vimrc` file

    runtime! ftplugin/man.vim
    " grep

Then save it and re-source the configuration with `:source %`.

With the `man.vim` plugin enabled, you can now move your cursor over the word `grep` and hit `<leader>K` which will pop open the man page for `grep` in a unnamed, split buffer.

Not only does this prevent context-switching when viewing a man page, but it also gives you the full power of vim over the content of the man page. You can search, you can yank text, or you can even pop open the man page for another command.

See `:h ft-man-plugin` for more details.## Vim Regex Word Boundaries

Today while writing a Vim regex to change every instance of `it` (ignoring larger matches like `itemized`), we stumbled upon Vim regex word boundary matching.

Given this file:

    foo
    foobar

The following Vim regex will match on both the first and second line:

    :%s/foo/baz/g

But with word boundaries, we'll only match (and change) the first line, because only the first `foo` is a standalone word:

    :%s/\<foo\>/baz/g

## Vim without the Extras

If you want to start up vim without loading all the usual plugins, you can supply the `--noplugin` flag

    $ vim --noplugin coffee.rb

You can take things even further by instead telling vim to open without loading any plugins or configuration files. That is, you can tell vim to skip all initializations.

    $ vim -u NONE coffee.rb

If you are used to lots of syntax highlighting, custom bindings, and other niceties, this may feel rather foreign.

## What is on the Runtime Path?

All of the plugins, syntax highlighting, language-specific indentation that extend the default behavior of Vim are on the runtime path. If something isn't on Vim's runtime path, then Vim won't know about and as a result will not load it at _runtime_. How do we see what is on the runtime path?

The `rtp` option is the shorthand for `runtimepath`. Calling `set` on either of these will show us the list of runtime paths, or at least some of them.

    :set rtp

This will generally be a truncated list if you have a lot of plugins. To be sure you are seeing all of them, use `echo` instead.

    :echo &rtp

See `:h rtp` for more details.

## Whole Line Auto-Completion

To get whole line auto-completion in Vim, you don't need a fancy plugin. It is built right in. There is a sub-mode of insert mode called *X mode* that allows you to do various kinds of special insertions. The `ctrl-x ctrl-l` binding corresponds to whole line completion. So, if you start typing a few characters and then (while still in insert mode) hit `ctrl-x ctrl-l` you will see a completed line that matches the initial characters you typed as well as a list of subsequent matches. You can cycle through the matches using `ctrl-n` and `ctrl-p` (going forward and backward, respectively).

The completion is done based on the configured completion sources. Generally, the completion sources will include the current buffer, other loaded and unloaded buffers, plus others. You can see which sources are configured with `:set complete?` and read more about the completion configuration at `:h 'complete'`.

## Wrap with Some Room

The [surround.vim](https://github.com/tpope/vim-surround) plugin allows you to wrap text objects with various surrounding characters (e.g. `( )`, `{ }`, `" "`). If you have a visual selection on `1 2 3 4 5` and type `S]` you will get:

    [1 2 3 4 5]

That works, but if you prefer a more readable version with some extra breathing room, you can make the visual selection and hit `S[` which will stick a space on either end:

    [ 1 2 3 4 5 ]

Now, if you already have some text wrapped in square braces, like the first example, and you want to convert it to the more spacious second example, you can do a *change surround* command followed by hitting the open square brace twice (that is, `cs[[`) which will convert `[1 2 3 4 5]` to `[ 1 2 3 4 5 ]`.

# Vim Plugins

## Amend Commits with Fugitive

Let's assume you are using [fugitive](https://github.com/tpope/vim-fugitive) for Git integration with Vim. You've made a commit that you want to amend. First, stage any changes that should be included in the amend with `:Gstatus` or
`:Ge:`. Then hit `ca` to open up the commit window for amending. Save and quit when you are finished.

Want to view the aggregate changes? Open the commit window for amending in verbose mode with `cva`.

See the [vim-fugitive docs](https://github.com/tpope/vim-fugitive/blob/master/doc/fugitive.txt) for more details.

## Blank Lines Above and Below

Generally when I want to add a line above or below the line that the cursor is on, I use `O` and `o`, respectively. This has a couple potential drawbacks. First and most prominent, the cursor is moved to the new line and left in insert mode. Usually, I'd like to remain in normal mode and stay on the current line. Second, these commands will emulate indentation and other
formatting rules. This is either exactly what you want or a bit of an annoyance.

The [`vim-unimpaired` plugin](https://github.com/tpope/vim-unimpaired) provides an alternative worth adding to your toolbelt. By hitting `[<space>` and `]<space>`, a new line will be opened above and below the current line, respectively. Additionally, it leaves you in normal mode, keeps the cursor on the current line, and moves the cursor to the first non-indented character. In the case of performing this command in the midst of a comment in a source code file, neither the indentation nor the comment character will be propagated onto the new line.

Hold on to `O`/`o` and `[<space>`/`]<space>` and know the difference. You'll likely need each of them from time to time.

## Build and Install a Go Program

With the [`vim-go`](https://github.com/fatih/vim-go) plugin, Vim gets all kinds of support for working with a Go project. Generally, with a Go project, you have to run `go build` to compile the project and if that is successful, you can run `go install` to put the executable binary on the `$GOPATH`.

This plugin allows you to tighten the feedback loop. You can build right within your Vim session using

    :GoBuild

which will alert you to any compilation errors.

You can then install the program using

    :GoInstall

Your program is now ready to run. It's worth noting that this plugin will also notify you about syntax errors
when you save, if they exist.

## Case-Aware Substitution with vim-abolish

Substitution in vim is, by default, case-sensitive. Adding the `i` `s-flag` makes it case-insensitive.
[`vim-abolish`](https://github.com/tpope/vim-abolish), on the other hand, lets you perform a case-insensitive substitution that preserves three case variants (foo, Foo, and FOO). Substitution with `vim-abolish` can be performed with `Subvert` or `S`.

For instance, `:%S/blog/article/g` will turn

    blog Blog bLOg BLOG

into

    article Article bLOg ARTICLE

Install `vim-abolish` and see `:h Subvert` for more details.

## Coercing Casing With vim-abolish

The [`vim-abolish`](https://github.com/tpope/vim-abolish) plugin provides a couple handy shortcuts for quickly coercing the casing of a variable.

For instance, if you have a variable in camel case and you want to change it snake case, you can navigate over the variable and hit `crs` (e.g. `myFavoriteVariable` -> `my_favorite_variable`).

Similarly, you can hit `crc` to change a variable to camel case. It even has support for mixed case (`crm`) and uppercase (`cru`).

## Creating Non-Existent Directories

When creating new files from within vim, using `:e`, you may find yourself creating that file in a directory that doesn't yet exist. Vim will tell you as much if you then try to save that file. To get around this, I have often shelled out with `:!mkdir %:h`. This is a bit awkward to type though.

The [`vim-eunuch`](https://github.com/tpope/vim-eunuch) plugin comes with a handy command for this. `:Mkdir` will create the parent directory for the current buffer. If you're in a situation where multiple levels of the buffer's directory don't exist, you can use `:Mkdir!` which will invoke `mkdir` with the `-p` flag.

## Deleting Buffers in BufExplorer

The [BufExplorer](https://github.com/jlanzarotta/bufexplorer) plugin makes it easy to browse and navigate to the various buffers open in a Vim session. It is based on your buffer list. After a bit of coding, your buffer list can start to get a bit out of control. There are surely going to be buffers that you want to close out, *delete* if you will.

Within the BufExplorer browser you can move your cursor onto a buffer and delete it.

To delete it by *unloading* the buffer (see `:h bd`), you can hit `d`.

To delete it by *wiping out* the buffer (see `:h bw`), you can hit `D`.

If you already have the plugin installed, see `:h bufexplorer` for more details.

## Interactive Buffer List

The `:ls` command is a great way to see what buffers you currently have open during a vim session. However, if you are trying to find and open a specific buffer it can be tedious to have to find it in the list and then enter a whole different command to move to it (e.g. `:b 10`).

The [`bufexplorer`](https://github.com/jlanzarotta/bufexplorer) plugin gives you quick access to an interactive buffer list. By using `<leader>bs` and `<leader>bv` you can open horizontally and vertically split windows, respectively, that allow you to navigate through and open specific buffers from your buffer list. This is a simple plugin you can add to your workflow that can make working with a lot of buffers a bit more efficient.

## Quick Quickfix List Navigation

There are lots of commands that will load up Vim's quickfix list with results that you'll want to traverse. For instance, if you use [Fugitive's `:Ggrep`](https://github.com/tpope/vim-fugitive/blob/master/doc/fugitive.txt#L94), it'll load up the quickfix list with line by line occurrences of the search term.

You can go forwards and backwards through this list using `:cnext` and `:cprevious`. Though this gets a bit tedious to type over and over, especially for long lists of results.

You can quickly navigate forwards and backwards through these results with two bindings provided by
[vim-unimpaired](https://github.com/tpope/vim-unimpaired). `]q` is mapped to `:cnext`, for going forwards, and `[q` is mapped to `:cprevious`, for going backwards.
