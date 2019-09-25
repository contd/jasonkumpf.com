---
title: Nix How-Tos and Tips
lang: en-US
tags:
  - nix
---

# {{ $page.title }}

This is more of a scratchpad and reference for me on various things I've found useful at somepoint and wanted to organize it all in one place.  Most are just composed of pieces copied and modified a little from where I've found them.

- [Cat a File with Line Numbers](#cat-a-file-with-line-numbers)
- [Change Default Shell for a User](#change-default-shell-for-a-user)
- [Change to that New Directory](#change-to-that-new-directory)
- [Check if a Port is in Use](#check-if-a-port-is-in-use)
- [Check the Current Working Directory](#check-the-current-working-directory)
- [Clear the Screen](#clear-the-screen)
- [Command Line Length Limitations](#command-line-length-limitations)
- [Copying File Contents to System Paste Buffer](#copying-file-contents-to-system-paste-buffer)
- [Create a File Descriptor with Process Substitution](#create-a-file-descriptor-with-process-substitution)
- [Curling For Headers](#curling-for-headers)
- [Curling with Basic Auth Credentials](#curling-with-basic-auth-credentials)
- [Display Free Disk Space](#display-free-disk-space)
- [Do Not Overwrite Existing Files](#do-not-overwrite-existing-files)
- [Exclude a Directory with Find](#exclude-a-directory-with-find)
- [File Type Info with File](#file-type-info-with-file)
- [Find Newer Files](#find-newer-files)
- [Get the Unix Timestamp](#get-the-unix-timestamp)
- [Global Substitution on the Previous Command](#global-substitution-on-the-previous-command)
- [Grep for Files without a Match](#grep-for-files-without-a-match)
- [Grep for Multiple Patterns](#grep-for-multiple-patterns)
- [Hexdump a Compiled File](#hexdump-a-compiled-file)
- [Jump to the Ends of Your Shell History](#jump-to-the-ends-of-your-shell-history)
- [Kill Everything Running on a Certain Port](#kill-everything-running-on-a-certain-port)
- [Killing a Frozen SSH Session](#killing-a-frozen-ssh-session)
- [Last Argument of the Last Command](#last-argument-of-the-last-command)
- [List All Users](#list-all-users)
- [List Names of Files with Matches](#list-names-of-files-with-matches)
- [List of Sessions to a Machine](#list-of-sessions-to-a-machine)
- [List Parent pid with ps](#list-parent-pid-with-ps)
- [Map a Domain to localhost](#map-a-domain-to-localhost)
- [Only Show the Matches](#only-show-the-matches)
- [Open the Current Command in an Editor](#open-the-current-command-in-an-editor)
- [Partial String Matching in Bash Scripts](#partial-string-matching-in-bash-scripts)
- [Repeat Yourself](#repeat-yourself)
- [Saying Yes](#saying-yes)
- [Search History](#search-history)
- [Search man Page Descriptions](#search-man-page-descriptions)
- [Securely Remove Files](#securely-remove-files)
- [Show Disk Usage for the Current Directory](#show-disk-usage-for-the-current-directory)
- [Sort in Numerical Order](#sort-in-numerical-order)
- [SSH Escape Sequences](#ssh-escape-sequences)
- [SSH with a Specific Key](#ssh-with-a-specific-key)
- [SSH with Port Forwarding](#ssh-with-port-forwarding)
- [Undo Some Command Line Editing](#undo-some-command-line-editing)
- [View a Web Page in the Terminal](#view-a-web-page-in-the-terminal)
- [Watch the Difference](#watch-the-difference)
- [Watch This Run Repeatedly](#watch-this-run-repeatedly)
- [Where are the Binaries?](#where-are-the-binaries)

## Cat a File with Line Numbers

You can quickly view a file using `cat` but with the `-n` flag you can view that file with line numbers

```bash
cat -n Gemfile
 1  source 'https://rubygems.org'
 2
 3
 4  # Bundle edge Rails instead: gem 'rails', github: 'rails/rails'
 5  gem 'rails', '4.2.0'
 6  # Use postgresql as the database for Active Record
 7  gem 'pg'
```
## Change Default Shell for a User

You can change the default shell program for a particular unix user with the `chsh` command. Just tell it what shell program you want to use (e.g. `bash` or `zsh`) and which user the change is for:

```bash
[sudo] chsh -s /usr/bin/zsh username
```

This command needs to be invoked with root privileges.  This command updates the entry for that user in the `/etc/passwd` file.

## Change to that New Directory

The `$_` variable provided by bash is always set to the last argument of the previous command. One handy use of this is for changing directories into a newly created directory.  This command will leave you in your newly created directory, `new_dir`.

```bash
mkdir new_dir && cd $_
```

## Check if a Port is in Use

The `lsof` command is used to *list open files*. This includes listing network connections. This means I can check if a particular port is in use and what process is using that port. For instance, I can check if a node.js application is currently running on port 3000.

```
$ lsof -i TCP:3000
COMMAND   PID       USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
node    13821 jason   12u  IPv6 0xdf2e9fd346cc12b5      0t0  TCP localhost:hbci (LISTEN)
node    13821 jason   13u  IPv4 0xdf2e9fd33ca74d65      0t0  TCP localhost:hbci (LISTEN)
```

I can see that a node process (my node.js app) is using port 3000. The PID and a number of other details are included.

See more details with `man lsof`.

## Check the Current Working Directory

Use `pwd` to display the absolute path of the current working directory.

```bash
pwd
```
See `man pwd` for more details.

## Clear the Screen

If you type `clear` into your shell, the screen will be cleared. There is a handy keybinding though that will save you a few keystrokes. Just hit `ctrl-l` to achieve the same effect.

## Command Line Length Limitations

The other day I tried to run a `rm` command on the contents of a directory with a **LOT** of files.

```bash
rm images/*
```

Instead of deleting the contents of the directory, the following message was displayed:

```bash
/bin/rm: cannot execute [Argument list too long]
```

Bash wanted to expand the entire command before executing it. It was too long. But what is too long?

It turns out that we can figure out the max length of commands with the following command:

```bash
getconf ARG_MAX
```

## Copying File Contents to System Paste Buffer

If you need to copy and paste the contents of a file, the `pbcopy` command can be one of the best ways to accomplish this. Simply `cat` the file and pipe that into `pbcopy` to get the contents of the file into the system paste buffer.

```bash
cat some-file.txt | pbcopy
```

See `man pbcopy` for more details.

## Create a File Descriptor with Process Substitution

Process substitution can be used to create a file descriptor from the evaluation of a shell command. The syntax for process substitution is `<(LIST)` where `LIST` is one or more bash commands.

```bash
cat <(echo 'hello, world')
hello, world
```

This is particularly useful for commands that expect files, such as diff:

```bash
diff <(echo 'hello, world') <(echo 'hello, mars')
1c1
< hello, world
---
> hello, mars
```

## Curling For Headers

If you want to inspect the headers of a response from some endpoint, look no further than a quick `curl` command. By including the `-I` flag, `curl` will return just the headers.

For example, if you are developing a web app that is being locally served at `localhost:3000` and you'd like to see what the headers look like for a particular URL, you might try something like the following command:

```bash
curl -I localhost:3000/posts
```

## Curling with Basic Auth Credentials

I often use `curl` to take a quick look at the responses of particular endpoints. If I try to `curl` a URL that is secured with HTTP Basic Authentication, this is what the response looks like:

```bash
curl staging.example.com
HTTP Basic: Access denied.
```

I can give the credentials to `curl` so that it can plug them in as it makes the request using the `-u` (or `--user`) flag:

```bash
curl -u username:password staging.example.com
<html><body>...</body></html>
```

If I don't want the password showing up in my command-line history, I can just provide the username and `curl` will prompt me for my password:

```bash
curl -u username staging.example.com
Enter host password for user 'username':
<html><body>...</body></html>
```

See `man curl` for more details.

## Display Free Disk Space

The `df` utility is a handy way to display the free disk space available on on a specific file system or all mounted file systems.

Use `df` with the `-h` flag to display the disk space usage and availability in a human-readable format.  Here is the output from a linode box of mine:

```bash
df -h
Filesystem      Size  Used Avail Use% Mounted on
/dev/xvda        20G  3.8G   16G  20% /
none            4.0K     0  4.0K   0% /sys/fs/cgroup
devtmpfs        994M  4.0K  994M   1% /dev
none            200M  196K  199M   1% /run
none            5.0M     0  5.0M   0% /run/lock
none            996M     0  996M   0% /run/shm
none            100M     0  100M   0% /run/user
```

## Do Not Overwrite Existing Files

When using the `cp` command to copy files, you can use the `-n` flag to make sure that you do not overwrite existing files.

## Exclude a Directory with Find

Using `find` is a handy way to track down files that meet certain criteria. However, if there are directories full of irrelevant files, you may end up with a lot of noise. What you want to do is exclude or ignore such directories. For example, you probably don't want `find` to return results from the `.git` directory of your project.

Specific directories can be excluded by combining the `-not` and `-path` arguments.

For instance, to see all files modified within the last 10 days, but not including anything in the `.git` directory, run the following:

```bash
find . -type f -not -path './.git/*' -ctime -10
```

## File Type Info with File

Use the `file` utility to determine the type of a file:

```bash
file todo.md
todo.md: ASCII English text

file Hello.java
Hello.java: ASCII C++ program text

file Hello.class
Hello.class: compiled Java class data, version 52.0
```

The `Hello.java` file isn't exactly a C++ program, but close enough.

## Find Newer Files

Use the `-newer` flag with the name of a file to find files that have a newer modification date than the named file.  For instance,

```bash
find blog -name '*.md' -newer blog/first-post.md
```

will find all markdown files in the `blog` directory that have a modification date more recent than `blog/first-post.md`.

## Get the Unix Timestamp

To get the Unix timestamp from your shell, use the `time` command with the appropriate format:

```bash
date +%s
```

## Global Substitution on the Previous Command

Let's say we just executed the following command:

```bash
grep 'foo' foo.md
```

It gave us the information we were looking for and now we want to execute a similar command to find the occurrences of `bar` in `bar.md`. The `^` trick won't quite work here.

```bash
^foo^bar<tab>
grep 'bar' foo.md
```

What we need is a global replace of `foo` in our previous command. The `!!` command can help when we sprinkle in some `sed`-like syntax.

```bash
!!gs/foo/bar<tab>
grep 'bar' bar.md
```

For a short command like this, we haven't gained much. However, for large commands that span the length of the terminal, this can definitely save us a little trouble.

## Grep for Files without a Match

The `grep` command is generally used to find files whose contents match a pattern. With the `-L` (`--files-without-match`) flag, `grep` can be used to find files that don't match the given pattern.

For instance, to find files in the current directory that don't have `foobar` anywhere in their content, run:

```bash
grep -L "foobar" ./*
```

## Grep for Multiple Patterns

You can use the `-e` flag with the `grep` command to search for a pattern. Additionally, you can use multiple `-e` flags to search for multiple patterns. For instance, if you want to search for occurrences of `ruby` and `clojure` in a `README.md` file, use the following command:

```bash
grep -e ruby -e clojure README.md
```

See `man grep` for more details.

## Hexdump a Compiled File

The `hexdump` unix utility allows you to dump the contents of a compiled/executable file in a _readable_ hexadecimal format. Adding the `-C` flag includes a sidebar with a formatted version of that row of hexadecimal.

For example, a compiled _Hello World_ java program, `Hello.java`, will look something like this:

```bash
cat Hello.class | hexdump -C
00000000  ca fe ba be 00 00 00 34  00 1d 0a 00 06 00 0f 09  |.......4........|
00000010  00 10 00 11 08 00 12 0a  00 13 00 14 07 00 15 07  |................|
00000020  00 16 01 00 06 3c 69 6e  69 74 3e 01 00 03 28 29  |.....<init>...()|
00000030  56 01 00 04 43 6f 64 65  01 00 0f 4c 69 6e 65 4e  |V...Code...LineN|
00000040  75 6d 62 65 72 54 61 62  6c 65 01 00 04 6d 61 69  |umberTable...mai|
00000050  6e 01 00 16 28 5b 4c 6a  61 76 61 2f 6c 61 6e 67  |n...([Ljava/lang|
00000060  2f 53 74 72 69 6e 67 3b  29 56 01 00 0a 53 6f 75  |/String;)V...Sou|
00000070  72 63 65 46 69 6c 65 01  00 0a 48 65 6c 6c 6f 2e  |rceFile...Hello.|
00000080  6a 61 76 61 0c 00 07 00  08 07 00 17 0c 00 18 00  |java............|
00000090  19 01 00 0d 48 65 6c 6c  6f 2c 20 57 6f 72 6c 64  |....Hello, World|
000000a0  21 07 00 1a 0c 00 1b 00  1c 01 00 05 48 65 6c 6c  |!...........Hell|
000000b0  6f 01 00 10 6a 61 76 61  2f 6c 61 6e 67 2f 4f 62  |o...java/lang/Ob|
000000c0  6a 65 63 74 01 00 10 6a  61 76 61 2f 6c 61 6e 67  |ject...java/lang|
000000d0  2f 53 79 73 74 65 6d 01  00 03 6f 75 74 01 00 15  |/System...out...|
000000e0  4c 6a 61 76 61 2f 69 6f  2f 50 72 69 6e 74 53 74  |Ljava/io/PrintSt|
000000f0  72 65 61 6d 3b 01 00 13  6a 61 76 61 2f 69 6f 2f  |ream;...java/io/|
00000100  50 72 69 6e 74 53 74 72  65 61 6d 01 00 07 70 72  |PrintStream...pr|
00000110  69 6e 74 6c 6e 01 00 15  28 4c 6a 61 76 61 2f 6c  |intln...(Ljava/l|
00000120  61 6e 67 2f 53 74 72 69  6e 67 3b 29 56 00 20 00  |ang/String;)V. .|
00000130  05 00 06 00 00 00 00 00  02 00 00 00 07 00 08 00  |................|
00000140  01 00 09 00 00 00 1d 00  01 00 01 00 00 00 05 2a  |...............*|
00000150  b7 00 01 b1 00 00 00 01  00 0a 00 00 00 06 00 01  |................|
00000160  00 00 00 01 00 09 00 0b  00 0c 00 01 00 09 00 00  |................|
00000170  00 25 00 02 00 01 00 00  00 09 b2 00 02 12 03 b6  |.%..............|
00000180  00 04 b1 00 00 00 01 00  0a 00 00 00 0a 00 02 00  |................|
00000190  00 00 03 00 08 00 04 00  01 00 0d 00 00 00 02 00  |................|
000001a0  0e                                                |.|
000001a1
```

## Jump to the Ends of Your Shell History

There are all sorts of ways to do things in your shell environment without reaching for the arrow keys. For instance, if you want to move _up_ to the previous command, you can hit `Ctrl-p`. To move _down_ to the next command in your shell history, you can hit `Ctrl-n`.

But what if you want to move to the beginning and end of your entire shell history?

Find your meta key (probably the one labeled `alt`) and hit `META-<` and `META->` to move to the end and beginning of your shell history, respectively.

## Kill Everything Running on a Certain Port

You can quickly kill everything running on a certain port with the following command.

```bash
sudo kill `sudo lsof -t -i:3000`
```

This gets a list of pids for all the processes and then kills them.

## Killing a Frozen SSH Session

Whenever an SSH session freezes, I usually mash the keyboard in desperation and then kill the terminal session. This can be avoided though. SSH will listen for the following kill command:

```bash
~.<cr>
```

This will kill the frozen SSH session and leave you in the terminal where you were before you SSH\'d.

## Last Argument of the Last Command

You can use `!$` as a way to reference the last argument in the last command. This makes for an easy shortcut when you want to switch out commands for the same long file name. For instance, if you just ran `cat` on a file to see its contents

```bash
cat /Users/jbranchaud/.ssh/config
```

and now you want to edit that file. You can just pass `!$` to the `vim` command:

```bash
vim !$
```

Hit enter or tab to get the full command:

```bash
vim /Users/jbranchaud/.ssh/config
```

## List All Users

On unix-based systems, all system users are listed with other relevant information in the `/etc/passwd` file. You can output a quick list of the users by name with the following command:

```bash
cut -d: -f1 /etc/passwd
```

## List Names of Files with Matches

I often use `grep` and `ag` to search for patterns in a group or directory of files. Generally I am interested in looking at the matching lines themselves. However, sometimes I just want to know the set of files that have matches. Both `grep` and `ag` can be told to output nothing more than the names of the files with matches when given the `-l` flag.

This can come in particularly handy if you just want a list of files that can be piped (or copied) for use with another command. This eliminates all the extra noise.

## List of Sessions to a Machine

The `last` command is a handy way to find out who has been connecting to a machine and when.

> Last will list the sessions of specified users, ttys, and hosts, in reverse time order.  Each line of output contains the user name, the tty from which the session was conducted, any hostname, the start and stop times for the session, and the duration of the session.  If the session is still continuing or was cut short by a crash or shutdown, last will so indicate.

In particular, this can be useful for finding an IP address that you want to connect to.

See `man last` for more details.

## List Parent pid with ps

The `ps` command, which stands for `process status`, is a great way to find different processes running on a machine. Information like their `pid` (_process id_) is included. If you are tracking down a process to kill and find that that process is an unkillable zombie, then you may need to simultaneously kill the process' parent as well.

So, you'll need the parent `pid` as well. You can get both the `pid` and the parent `pid` of a process by including the `-f` flag with `ps`.

You may also want to include the `-e` flag to make sure that information about other users\' processes is included in the results.

## Map a Domain to localhost

Do you want your computer to treat a domain as `localhost`? You can map it as such in your `/etc/hosts` file. For example, if I have an web app that refers to itself with the `dev.app.com` domain, I can add the following line to my `/etc/hosts` file to make sure the domain resolves to `localhost`:

    127.0.0.1   dev.app.com

Now, if I pop open my browser and visit `dev.app.com:3000`, I will see whatever is being served to `localhost:3000`.

## Only Show the Matches

Tools like `grep`, `ack`, and `ag` make it easy to search for lines in a file that contain certain text and patterns. They all come with the `-o` flag which tells them to only show the part that matches.

This is particularly powerful when used with regex and piped into other programs.

## Open the Current Command in an Editor

If you are working with a complicated command in the terminal trying to get the arguments just right. Such as this `curl`:

```bash
curl https://api.stripe.com/v1/customers \
   -u sk_test_BQokikJOvBiI2HlWgH4olfQ2: \
   -d description="Customer for test@example.com" \
   -d source=tok_189fCz2eZvKYlo2CsGERUNIW
```

It can be tedious to move to and modify various parts of the command. However, by hitting `Ctrl-x Ctrl-e`, the contents of the command buffer will be opened into your default editor (i.e. `$EDITOR`). This will make editing the command a bit easier. Saving and quitting the editor will put the updated command in the command buffer, ready to run.

Hit `Ctrl-x Ctrl-e` with an empty command buffer if you want to start crafting a command from scratch or if you are pasting one in from somewhere.

## Partial String Matching in Bash Scripts

To compare two strings in a bash script, you will have a snippet of code similar to the following:

```bash
if [[ $(pwd) == "/path/to/current/directory" ]]
then
  echo "You are in that directory";
fi
```

You may only want to do a partial string match. For this, you can use the `*` wildcard symbol.

```bash
if [[ $(pwd) == *"directory"* ]]
then
  echo "You are in that directory";
fi
```

## Repeat Yourself

Use the `repeat` command to repeat some other command.  You can repeat a command any number of times like so:

```bash
repeat 5 say Hello World
```

## Saying Yes

Tired of being prompted for confirmation by command-line utilities? Wish you could blindly respond 'yes' to whatever it is they are bugging you about? The `yes` command is what you've been looking for.

```bash
yes | rm -r ~/some/dir
```

This will respond `y` as `rm` asks for confirmation on removing each and every file in that directory.

`yes` is just as good at saying *no*. Give it `no` as an argument and it will happily (and endlessly) print `no`.

```bash
yes no
```

## Search History

Often times there is a very specific command you have entered into your bash prompt that you need to run again. You don't want to have to type it againand stepping manually through your history may be suboptimal if you typed it quite a while ago. Fortunately, there is a simple history search feature that you can use in this kind of situation.

Hit `Ctrl+r` and then start typing a moderately specific search term. Your search history will be filtered by that term. Subsequent hitting of `Ctrl+r` will step forward through that filtered history. Once you find the command you are looking for, hit enter to execute it.

## Search man Page Descriptions

You can use the `apropos` command with a keyword argument to search for that words occurrence throughout all the man pages on your system. For instance, invoking `apropos whatis` will bring up a list of entries including the `apropos` command itself.

See `man apropos` for more details.

## Securely Remove Files

If you really want to make sure you have wiped a file from your hard drive, you are going to want to use `srm` instead of `rm`. The man page for `srm` gives the following description:

> srm  removes  each  specified file by overwriting, renaming, and truncating it before unlinking. This prevents other people from undeleting or recovering any information about the file from the command line.


## Show Disk Usage for the Current Directory

The `du` utility can be used to show disk usage for a particular directory or set of directories. When used without any arguments, it will show the disk usage for the current directory.

```bash
du
80      ./.git/hooks
8       ./.git/info
256     ./.git/logs/refs/heads
...
```

with the `-h` command we can see it all in a human-readable format

```bash
du -h
 40K    ./.git/hooks
4.0K    ./.git/info
128K    ./.git/logs/refs/heads
```

and to get an even clearer picture we can pipe that through `sort -nr`

```bash
du -h | sort -nr
412K    ./vim
352K    ./postgres
340K    ./.git/logs
216K    ./.git/logs/refs
184K    ./ruby
156K    ./unix
148K    ./git
...
```

This sorts it numerically in reverse order putting the largest stuff at the top.

## Sort in Numerical Order

By default, the `sort` command will sort things alphabetically. If you have numerical input though, you may want a numerical sort. This is what the `-n` flag is for.

If I have a directory of files with numbered names, sort doesn't quite do the job by itself.

```bash
ls | sort
1.txt
10.txt
11.txt
12.txt
2.txt
3.txt
4.txt
5.txt
```

with the `-n` flag, I get the sort order I am looking for.

```bash
ls | sort -n
1.txt
2.txt
3.txt
4.txt
5.txt
10.txt
11.txt
12.txt
```


## SSH Escape Sequences

In a previous post, I talked about an escape sequence for breaking out of an SSH session when the pipe has been broken. This isn't the only SSH escape sequence though. To see the others, hit `<Enter>~?`. This displays a help list with all the other escape sequences.

```bash
~?
Supported escape sequences:
 ~.   - terminate connection (and any multiplexed sessions)
 ~B   - send a BREAK to the remote system
 ~C   - open a command line
 ~R   - request rekey
 ~V/v - decrease/increase verbosity (LogLevel)
 ~^Z  - suspend ssh
 ~#   - list forwarded connections
 ~&   - background ssh (when waiting for connections to terminate)
 ~?   - this message
 ~~   - send the escape character by typing it twice
(Note that escapes are only recognized immediately after newline.)
```

## SSH with a Specific Key

When you SSH into another machine using public key authentication, the key pair from either `~/.ssh/id_dsa`, `~/.ssh/id_ecdsa`, or `~/.ssh/id_rsa` is used by default. This is generally what you want. But what if the target server is expecting to identify you with a different SSH key pair?

The `-i` option can be used with `ssh` to specify a different _identity file_ when the default isn\'t what you want.

## SSH with Port Forwarding

Use the `-L` flag with `ssh` to forward a connection to a remote server

```bash
ssh someserver -L3000:localhost:3000
```

## Undo Some Command Line Editing

When using some of the fancy command line editing shortcuts, such as `ctrl-u`, you may end up erroneously changing or deleting part of the current command. Retyping it may be a pain or impossible if you've forgotten exactly what was changed. Fortunately, bash's command line editing has undo built in. Just hit `ctrl-_` a couple times to get back to where you want to be.


## View a Web Page in the Terminal

There are a number of programs out there that allow you to view a web page from right within the terminal. These text-based browsers are great for viewing pages that make basic use of html without relying on JavaScript for fancy user interactions, such as the docs page for a language.

One such program is `link`. If you don't already have it, you can install it with something like `homebrew`. Then run `links` with any URL for a magnificent, ad-free experience:

```bash
links http://www.postgresql.org/docs/current/static/functions-string.html
```

## Watch the Difference

The `watch` command is a simple way to repeatedly run a particular command. I'll sometimes use it to monitor the response from some endpoint. `watch` can make monitoring responses even easier when the `-d` flag is employed.
This flag instructs `watch` to highlight the parts of the output that are *different* from the previous run of the command.

So if I run

```bash
watch -d curl -LIs localhost:3000
```

I can easily see if the http status of the request changes.

## Watch This Run Repeatedly

I usually reach for a quick bash for loop when I want to run a particular process a bunch of times in a row. The `watch` command is another way to run a process repeatedly.

```bash
watch rspec spec/some/test.rb
```

The default is 2 seconds in between subsequent executions of the command.  The period can be changed with the `-n` flag though:

```bash
watch -n 2 rspec spec/some/test.rb
```

## Where are the Binaries?

When I want to know where an executable is, I use `which` like so:

```bash
which rails
/Users/jbranchaud/.gem/ruby/2.1.4/bin/rails
```

That is the rails binary on my path that will be used if I enter a rails command.  However, with something like rails, there may be multiple versions on your path. If you want to know where all of them are, you can use `where`, like
so:

```bash
where rails
/Users/jbranchaud/.gem/ruby/2.1.4/bin/rails
/Users/jbranchaud/.rubies/2.1.4/bin/rails
/usr/bin/rails
```
