---
title: Git How-Tos and Tips
lang: en-US
tags:
	- vcs
---

# {{ $page.title }}

This is more of a scratch pad and reference for me on various things I've found useful at some point and wanted to organize it all in one place.  Most are just composed of pieces copied and modified a little from where I've found them.

- [Gitignore Global](#gitignore-global)
- [Gitconfig](#gitconfig)
- [Basics](#basics)
- [Accessing a Lost Commit](#accessing-a-lost-commit)
- [Amend Author of Previous Commit](#amend-author-of-previous-commit)
- [Checking Commit Ancestry](#checking-commit-ancestry)
- [Checkout Old Version of a File](#checkout-old-version-of-a-file)
- [Checkout Previous Branch](#checkout-previous-branch)
- [Cherry Pick A Range Of Commits](#cherry-pick-a-range-of-commits)
- [Clean Out All Local Branches](#clean-out-all-local-branches)
- [Clean Up Old Remote Tracking References](#clean-up-old-remote-tracking-references)
- [Delete All Untracked Files](#delete-all-untracked-files)
- [Delete Remote Git Tags](#delete-remote-git-tags)
- [Determine the Hash Id for a Blob](#determine-the-hash-id-for-a-blob)
- [Diffing with Patience](#diffing-with-patience)
- [Dry Runs in Git](#dry-runs-in-git)
- [Excluding Files Locally](#excluding-files-locally)
- [Find The Initial Commit](#find-the-initial-commit)
- [Git Log since](#git-log-since)
- [Git Log With Authors](#git-log-with-authors)
- [Git Snapshot](#git-snapshot)
- [Grab A Single File From A Stash](#grab-a-single-file-from-a-stash)
- [Grep For a Pattern on Another Branch](#grep-for-a-pattern-on-another-branch)
- [Grep Over Commit Messages](#grep-over-commit-messages)
- [Ignore Changes to a Tracked File](#ignore-changes-to-a-tracked-file)
- [Intent to Add](#intent-to-add)
- [Interactively Unstage Changes](#interactively-unstage-changes)
- [Last Commit a File Appeared](#last-commit-a-file-appeared)
- [List All Files Changed Between Two Branches](#list-all-files-changed-between-two-branches)
- [List Different Commits Between Two Branches](#list-different-commits-between-two-branches)
- [List Filenames without the Diffs](#list-filenames-without-the-diffs)
- [List Most Git Commands](#list-most-git-commands)
- [List Untracked Files](#list-untracked-files)
- [Move The Latest Commit To A New Branch](#move-the-latest-commit-to-a-new-branch)
- [Reference a Commit Via Commit Message Pattern Matching](#reference-a-commit-via-commit-message-pattern-matching)
- [Rename a Remote](#rename-a-remote)
- [Renaming A Branch](#renaming-a-branch)
- [Resetting a Reset](#resetting-a-reset)
- [Show All Commits for a File Beyond Renaming](#show-all-commits-for-a-file-beyond-renaming)
- [Show the diffstat Summary of a Commit](#show-the-diffstat-summary-of-a-commit)
- [Single Key Presses in Interactive Mode](#single-key-presses-in-interactive-mode)
- [Staging Changes within Vim](#staging-changes-within-vim)
- [Staging Stashes Interactively](#staging-stashes-interactively)
- [Stashing Only Unstaged Changes](#stashing-only-unstaged-changes)
- [Stashing Untracked Files](#stashing-untracked-files)
- [The Alpha Commit](#the-alpha-commit)
- [Undo a Git Mistake](#undo-a-git-mistake)
- [Untrack a Directory of Files Without Deleting](#untrack-a-directory-of-files-without-deleting)
- [Untrack a File Without Deleting It](#untrack-a-file-without-deleting-it)
- [Update the URL of a Remote](#update-the-url-of-a-remote)
- [Verbose Commit Message](#verbose-commit-message)
- [Viewing a File on Another Branch](#viewing-a-file-on-another-branch)
- [What Changed?](#what-changed)
- [What is the Current Branch?](#what-is-the-current-branch)
- [Whitespace Warnings](#whitespace-warnings)


## Gitignore Global

You can have a global `.gitignore`file where many common entries that seem to be in almost all my `.gitignore` files.  Just create a `.gitignore_global`file in your home directory.  You also need to add an entry to your `.gitconfig`file which is show in the next section.

## Gitconfig

Here is a sample based on my real `.gitconfig` file but with dummy values for settings that are sensative.

```
[user]
	name = Jason Kumpf
	email = kumpf.jason@gmail.com
	signingkey = XXCCDDKKXXXXXXXX
[github]
	user = contd
	token = akdfgkagfalhdgfkhagfkajgsfkjag
[core]
	editor = vim
	excludesfile = /home/jason/.gitignore_global
	pager = diff-so-fancy | less --tabs=4 -RFX
[color]
	ui = true
[push]
	default = simple
[filter "lfs"]
	clean = git-lfs clean -- %f
	smudge = git-lfs smudge -- %f
	required = true
[gpg]
	program = gpg2
[alias]
	commend = commit --amend --no-edit
	it = !git init && git commit -m \"root\" --allow-empty
	shorty = status --short --branch
	grog = log --graph --abbrev-commit --decorate --all --format=format:\"%C(bold blue)%h%C(reset) - %C(bold cyan)%aD%C(dim white) - %an%C(reset) %C(bold green)(%ar)%C(reset)%C(bold yellow)%d%C(reset)%n %C(white)%s%C(reset)\"
```

## Basics

Here are some of the most common and basic commands shown in a typical workflow for setting up a project with existing files.  Here we have a directory named `notes`which will also be the repository name on Github.  It looks like the following:

```
├── general
│   └── some-general-note.md
├── home
│   └── shopping.md
├── README.md
├── todo
│   ├── 0001.md
│   ├── 0002.md
│   ├── 0003.md
│   ├── 0004.md
│   └── 0005.md
└── work
    ├── project1.md
    └── project2.md
```

It's always a good idea to have a `README.md` and `.gitignore`file for every repository especially if you are putting it on Github.  Now you have to create the repository on Github which will give you some basic instructions to add you files to it but here is what I use (only slightly different):

```bash
cd notes/
git init
git add .
git commit -m "Initial commit"
git remote add origin git@github.com:your-github-username/notes.git
git push origin master
```

If you get errors on the `git push` command its probably because you don't have ssh setup with your Github account.  See their documentation on that and you can test that its working with this command:

```bash
ssh -T git@github.com
```

Now that you have your files on Github and up to date you can start modifing, adding and removing files.  One thing to help it to remember to do the following whenever you add or remove a file or directory.

```bash
git add new_file.md
git rm old_file.md
git rm -r old_directory/
```

Also, when you make changes its good practice to use separate `git commit` commands for each so your comments are more specific to the committed file.


## Accessing a Lost Commit

If you have lost track of a recent commit (perhaps you did a reset), you can generally still get it back. Run `git reflog` and look through the output to see if you can find that commit. Note the sha value associated with that commit. Let's say it is `39e85b2`. You can peruse the details of that commit with `git show 39e85b2`.

From there, the utility belt that is git is at your disposal. For example, you can `cherry-pick` the commit or do a `rebase`.

## Amend Author of Previous Commit

The author of the previous commit can be amended with the following command

```bash
git commit --amend --author "Jason Kumpf <jkumpf@somedomain.com>"
```

## Checking Commit Ancestry

I have two commit shas and I want to know if the first is an ancestor of the second. Put another way, is this first commit somewhere in the history of this other commit.

Git's `merge-base` command combined with the `--is-ancestor` flag makes answering this question easy. Furthermore, because it is a plumbing command, it can be used in a script or sequence of commands as a switch based on the answer.

Here is an example of this command in action:

```bash
git merge-base --is-ancestor head~ head && echo 'yes, it is'
yes, it is
git merge-base --is-ancestor head~ head~~ && echo 'yes, it is'
```

In the first command, `head~` is clearly an ancestor of `head`, so the `echo` command is triggered. In the second, `head~` is not an ancestor of `head~~` so the return status of 1 short-circuits the rest of the command. Hence, no `echo`.

See `man git-merge-base` for more details.

## Checkout Old Version of a File

When you want to return to a past version of a file, you can reset to a past commit. When you don't want to abandon a bunch of other changes, this isn't going to cut it. Another option is to just checkout the particular file as it was at the time of a past commit.

If the sha of that past commit is `72f2675` and the file's name is `some_file.js`, then just use checkout like so:

```
git checkout 72f2675 some_file.js
```
## Checkout Previous Branch

Git makes it easy to checkout the last branch you were on.

```bash
git checkout -
```

This is shorthand for `git checkout @{-1}` which is a way of referring to the previous (or last) branch you were on. You can use this trick to easily bounce back and forth between `master` and a feature branch.

## Cherry Pick A Range Of Commits

Git's `cherry-pick` command allows you to specify a range of commits to be cherry picked onto the current branch. This can be done with the `A..B` style syntax -- where `A` is the older end of the range.

Consider a scenario with the following chain of commits: `A - B - C - D`.

```bash
git cherry-pick B..D
```

This will cherry pick commits `C` and `D` onto `HEAD`. This is because the lower-bound is exclusive. If you'd like to include `B` as well. Try the following:

```bash
git cherry-pick B^..D
```

See `man git-cherry-pick` for more details.


## Clean Out All Local Branches

Sometimes a project can get to a point where there are so many local branches that deleting them one by one is too tedious. This one-liner can help:

```bash
git branch --merged master | grep -v master | xargs git branch -d
```

This won't delete branches that are unmerged which saves you from doing something stupid, but can be annoying if you know what you are doing. If you are sure you want to wipe everything, just use `-D` like so:

```bash
git branch --merged master | grep -v master | xargs git branch -D
```

## Clean Up Old Remote Tracking References

After working on a Git-versioned project for a while, you may find that there are a bunch of references to remote branches in your local repository. You know those branches definitely don't exist on the remote server and you've removed the local branches, but you still have references to them lying around. You can reconcile this discrepancy with one command:

```bash
git fetch origin --prune
```

This will prune all those non-existent remote tracking references which is sure to clean up your git log (`git log --graph`).

## Delete All Untracked Files

Git provides a command explicitly intended for cleaning up (read: removing) untracked files from a local copy of a repository.

> git-clean - Remove untracked files from the working tree

Git does want you to be explicit though and requires you to use the `-f` flag to force it (unless otherwise configured).

Git also gives you fine-grained control with this command by defaulting to only deleting untracked files in the current directory. If you want directories of untracked files to be removed as well, you'll need the `-d` flag.

So if you have a local repository full of untracked files you'd like to get rid of, just:

```bash
git clean -f -d
```

or just:

```bash
git clean -fd
```

## Delete Remote Git Tags

Tagging releases with Git is a good idea. In case your tags get off track, here is how you delete a Git tag locally and on a remote:

```bash
git tag -d abc
git push origin :refs/tags/abc
To git@github.com:hashrocket/hr-til
 - [deleted]         abc
```

It gets trickier if you're using Semantic Versioning, which includes dots in the tag name. The above won't work for `v16.0.0`. This will:

```bash
git tag -d v16.0.0
git push origin :v16.0.0
To git@github.com:hashrocket/hr-til
 - [deleted]         v16.0.0
```

## Determine the Hash Id for a Blob

Git's `hash-object` command can be used to determine what hash id will be used by git when creating a blob in its internal file system.

```bash
echo 'Hello, world!' > hola
git hash-object hola
af5626b4a114abcb82d63db7c8082c3c4756e51b
```

When a commit happens, git will generate this digest (hash id) based on the contents of the file. The name and location of the file don't matter, just the contents. This is the magic of git. Anytime git needs to store a blob, it can quickly match against the hash id in order to avoid storing duplicate blobs.

Try it on your own machine, you don't even need to initialize a git repository to use `git hash-object`.

## Diffing with Patience

The default diff algorithm used by Git is pretty good, but it can get mislead by larger, complex changesets. The result is a noisier, misaligned diff output.

If you'd like a diff that is generally a bit cleaner and can afford a little slow down (you probably can), you can instead use the `patience` algorithm which is described as such:

> Patience Diff, instead, focuses its energy on the low-frequency high-content lines which serve as markers or signatures of important content in the text. It is still an LCS-based diff at its core, but with an important difference, as it only considers the longest common subsequence of the signature lines:

> Find all lines which occur exactly once on both sides, then do longest common subsequence on those lines, matching them up.

You can set this as the default algorithm by adding the following lines to your `~/.gitconfig` file:

```bash
[diff]
    algorithm = patience
```

or it can be set from the command line with:

```bash
git config --global diff.algorithm patience
```

## Dry Runs in Git

There are a few commands in git that allow you to do a *dry run*. That is, git will tell you the effects of executing a command without actually executing that command.

For instance, if you are clearing out untracked files, you can double check what files are going to be deleted with the *dry run* flag, like so:

```bash
git clean -fd --dry-run
Would remove tmp.txt
Would remove stuff/
```

Similarly, if you want to check in which files a commit is going to be incorporated, you can:

```bash
git commit --dry-run --short
M  README.md
A  new_file.rb
```

Try running `git commit --dry-run` (that is, without the `--short` flag).  Look familiar? That is the same output you are getting from `git status`.


## Excluding Files Locally

Excluding (read: ignoring) files that should not be tracked is generally done by listing such files in a tracked `.gitignore` file. Though it doesn't make sense to list all kinds of excluded files here. For files that you'd like to exclude that are temporary or are specific to your local environment, there is another option. These files can be added to the `.git/info/exclude` file as a way of ignoring them locally.

Add specific files or patterns as needed to that file and then save it. Relevant files will no longer show up as untracked files when you `git status`.

## Find The Initial Commit

By definition, the initial commit in a repository has no parents. You can exploit that fact and use `rev-list` to find the initial commit; a commit with no parents.

```bash
git rev-list --max-parents=0 HEAD
```

The `rev-list` command lists all commits in reverse chronological order. By restricting them to those with at most 0 parents, you are only going to get root commits. Generally, a repository will only have a single root commit, but it is possible for there to be more than one.

See `man git-rev-list` for more details.

## Git Log since

At the end of each day, I try to record what I did, to jog my memory during the next morning's standup. This is a helpful aid:

```bash
git log --since="24 hours ago"
```

I SSH into my work machine and run this in my project's root directory. Combined with an alias from the Hashrocket Dotmatrix, `glg` (`git log --graph --oneline --decorate --color --all`), I get a terse summary of the day's work, ready to be pasted into your note-taking or project management tool of choice:

```bash
glg --since="24 hours ago"
* 7191b92 (HEAD, origin/master, origin/HEAD, master) Good changes
* 3f4d61e More good changes
* ecd9dcd Even more
...
```

## Git Log With Authors

In my [never-ending quest](https://til.hashrocket.com/posts/32d01c979e-git-log-since) to better summarize my work at the end of the day using computers, I discovered today the Git `--author` flag. It works like this:

```bash
glg --since=midnight --author=dev+jwworth+mikechau@hashrocket.com
* 4ba91a8 (HEAD, origin/checkout, checkout) Add guard for manual entry of employee discount
* 3a4e4c9 Seed a coupon and code and auto-apply in preview
* cb1adee Add discount
...
```

The alias `glg` is discussed [here](https://til.hashrocket.com/posts/32d01c979e-git-log-since).

I use this when multiple developers or teams are committing throughout the day to the same repository, to disambiguate our work from others. Ready to paste into your billing software of choice.

## Git Snapshot

To save a snapshot of your current work in git, try this command:

```bash
git stash save "snapshot: $(date)" && git stash apply "stash@{0}"
```

This saves your current work in a timestamped stash, without removing it. In Hashrocket's dotmatrix this command is aliased to `git snapshot`.

## Grab A Single File From A Stash

In git, you can reference a commit SHA or branch to checkout differing versions of files.

```bash
git checkout d3d2e38 -- README.md
```

In the same way, you can snag the version of a file as it existed in a
stash. Just reference the relevant stash.

```bash
git checkout stash@{1} -- README.md
```

## Grep For a Pattern on Another Branch

Git has a built-in `grep` command that works essentially the same as the standard `grep` command that unix users are used to. The benefit of `git-grep` is that it is tightly integrated with Git. You can search for occurrences of a pattern on another branch. For example, if you have a feature branch, `my-feature`,  on which you'd like to search for occurrences of `user.last_name`, then your command would look like this:

```bash
git grep 'user\.last_name' my-feature
```

If there are matching results, they follow this format:

```
my-feature:app/views/users/show.html.erb:  <%= user.last_name %>
...
```

This formatting is handy because you can easily copy the branch and file directive for use with [`git-show`](viewing-a-file-on-another-branch.md).

See `man git-grep` for more details.

## Grep Over Commit Messages

The `git log` command supports a `--grep` flag that allows you to do a text search (using grep, obviously) over the commit messages for that repository. For the git user that writes descriptive commit messages, this can come in quite handy. In particular, this can be put to use in an environment where the standard process involves including ticket and bug numbers in the commit message. For example, finding bug `#123` can be accomplished with:

```bash
git log --grep="#123"
```

See `man git-log` for more details.

## Ignore Changes to a Tracked File

Files that should never be tracked are listed in your `.gitignore` file. But what about if you want to ignore some local changes to a tracked file?

You can tell git to assume the file is unchanged

```bash
git update-index --assume-unchanged <some-file>
```

Reversing the process can be done with the `--no-assume-unchanged` flag.

## Intent to Add

Git commands like `git diff` and `git add --patch` are awesome, but their little caveat is that they only work on files that are currently tracked in the repository. That means that after working on that new feature for 30 minutes, a `git diff` is only going to show you the changes to existing files and when you go to start making commits with `git add --patch` you will again be provided with only part of the picture.

The `git add` command has a flag, `-N`, that is described in the git man pages:

> Record only the fact that the path will be added later. An entry for the path is placed in the index with no content. This is useful for, among other things, showing the unstaged content of such files with git diff and committing them with git commit -a.

By adding untracked files with the `-N` flag, you are stating your *intent to add* these files as tracked files. Once git knows that these files are meant to be tracked, it will know to include them when doing things like computing the diff, performing an interactive add, etc.

## Interactively Unstage Changes

I often use `git add --patch` to interactively stage changes for a commit. Git takes me through changes to tracked files piece by piece to check if I want to stage them. This same interactive _staging_ of files can be used in reverse when removing changes from the index. Just add the `--patch` flag.

You can use it for a single file

```bash
git reset --patch README.md
```

or you can let it operate on the entire repository

```bash
git reset --patch
```

This is useful when you've staged part of a file for a commit and then realize that some of those changes shouldn't be committed.

## Last Commit a File Appeared

In my project, I have a `README.md` file that I haven't modified in a while. I'd like to take a look at the last commit that modified it. The `git log` command can be used here with few arguments to narrow it down.

```bash
git log -1 -- README.md
commit 6da76838549a43aa578604f8d0eee7f6dbf44168
Author: jbranchaud <jbranchaud@gmail.com>
Date:   Sun May 17 12:08:02 2015 -0500

    Add some documentation on configuring basic auth.
```

This same command will even work for files that have been deleted if you know the path and name of the file in question. For instance, I used to have an `ideas.md` file and I'd like to find the commit that removed it.

```bash
git log -1 -- docs/ideas.md
commit 0bb1d80ea8123dd12c305394e61ae27bdb706434
Author: jbranchaud <jbranchaud@gmail.com>
Date:   Sat May 16 14:53:57 2015 -0500

    Remove the ideas doc, it isn\'t needed anymore.
```

## List All Files Changed Between Two Branches

The `git-diff` command can help with finding all files that have changed between two branches. For instance, if you are at the `HEAD` of your current feature branch and you'd like to see the list of files that have changed since being in sync with the `master` branch, you'd formulate a command like the following:

```bash
git diff --name-only master
```

The `--name-only` flag is important because it cuts out the rest of the non-essential output.

You'll want to list the _older_ branch first and the current branch second as a way of showing what has changed from the original branch to get to the current branch.

Though the example shows the use of _branches_ any git reference can be used to see what has changed between two commits.

See `man git-diff` for more details.

## List Different Commits Between Two Branches

There are times when I want to get a sense of the difference between two branches. I don't want to look at the actual diff though, I just want to see what commits are on one and the other.

I can do just this by using the `git-log` command with a couple flags, most importantly the `--cherry-pick` flag.

To compare the `feature` branch against the `master` branch, I can run a command like the following:

```bash
git log --left-right --graph --cherry-pick --oneline feature...branch
```

This lists commits with the first line of their messages. It also includes either a `<` or `>` arrow at the front of each commit indicating whether the commit is on the left (`feature`) or right (`master`) branch, respectively.

**Note: you can use more than branches in a command like this. Any two references will work. You can just use two SHAs for instance.**

## List Filenames without the Diffs

The `git show` command will list all changes for a given reference including the diffs. With diffs included, this can get rather verbose at times. If you just want to see the list of files involved (excluding the diffs), you can use the `--name-only` flag. For instance,

```bash
git show HEAD --name-only
commit c563bafb511bb984c4466763db7e8937e7c3a509
Author: jbranchaud <jbranchaud@gmail.com>
Date:   Sat May 16 20:56:07 2015 -0500

    This is my sweet commit message

app/models/user.rb
README.md
spec/factories/user.rb
```

## List Most Git Commands

You can list most git commands by using the `-a` flag with `git-help`:

```
git help -a
```

## List Untracked Files

Git's `ls-files` command is a plumbing command that allows you to list different sets of files based on the state of the working directory. So, if you want to list all untracked files, you can do:

```bash
git ls-files --others
```

This includes files *ignored* by the `.gitignore` file. If you want to exclude the ignored files and only see *untracked* and *unignored* files, then you can do:

```bash
git ls-files --exclude-standard --others
```

There are some other flags worth checking at at `git help ls-files`.


## Move The Latest Commit To A New Branch

I sometimes find myself making a commit against the `master` branch that I intended to make on a new branch. To get this commit on a new branch, one possible approach is to do a reset, checkout a new branch, and then re-commit it. There is a better way.

```bash
git checkout -b my-new-branch
git checkout -
git reset --hard HEAD~
```

This makes better use of branches and avoids the need to redo a commit that has already been made.

Note: The example was against the `master` branch, but can work for any branch.

## Reference a Commit Via Commit Message Pattern Matching

Generally when referencing a commit, you'll use the SHA or a portion of the SHA. For example with `git-show`:

```bash
git show cd6a63d014
...
```

There are many ways to reference commits though. One way is via regex pattern matching on the commit message. For instance, if you recently had a commit with a typo and you had included *typo* in the commit message, then you could reference that commit like so:

```bash
git show :/typo
Author: Josh Branchaud
Date: Mon Dec 21 15:50:20 2015 -0600

    Fix a typo in the documentation
...
```

By using `:/` followed by some text, git will attempt to find the most recent commit whose commit message matches the text. As I alluded to, regex can be used in the text.

See `$ man gitrevisions` for more details and other ways to reference commits.

## Rename a Remote

If you just added a remote (`git remote add ...`) and messed up the name or just need to rename some existing remote, you can do so with the `rename` command.

First, let's see the remotes we have:

```bash
git remote -v
origin  https://github.com/jbranchaud/til.git (fetch)
origin  https://github.com/jbranchaud/til.git (push)
```

To then rename `origin` to `destination`, for example, we can issue the following command:

```bash
git remote rename origin destination
```

See `man git-remote` for more details.

## Renaming A Branch

The `-m` flag can be used with `git branch` to move/rename an existing branch. If you are already on the branch that you want to rename, all you need to do is provide the new branch name.

```bash
git branch -m <new-branch-name>
```

If you want to rename a branch other than the one you are currently on, you must specify both the existing (old) branch name and the new branch name.

```bash
git branch -m <old-branch-name> <new-branch-name>
```

## Resetting a Reset

Sometimes we run commands like `git reset --hard HEAD~` when we shouldn't have. We wish we could undo what we've done, but the commit we've *reset* is gone forever. Or is it?

When bad things happen, `git-reflog` can often lend a hand. Using `git-reflog`, we can find our way back to were we've been; to better times.

```bash
git reflog
00f77eb HEAD@{0}: reset: moving to HEAD~
9b2fb39 HEAD@{1}: commit: Add this set of important changes
...
```

We can see that `HEAD@{1}` references a time and place before we destroyed our last commit. Let's fix things by resetting to that.

```bash
git reset HEAD@{1}
```

Our lost commit is found.

Unfortunately, we cannot undo all the bad in the world. Any changes to
tracked files will be irreparably lost.

## Show All Commits for a File Beyond Renaming

By including `-- <filename>` with a `git log` command, we can list all the commits for a file. The following is an example of such a command with some formatting and file names.

```bash
git log --name-only --pretty=format:%H -- README.md
4e57c5d46637286731dc7fbb1e16330f1f3b2b7c
README.md

56955ff027f02b57212476e142a97ce2b7e60efe
README.md

5abdc5106529dd246450b381f621fa1b05808830
README.md
```

What we may not realize, though, is that we are missing out on a commit in this file's history. At one point, this file was renamed. The command above wasn't able to capture that.

Using the `--follow` flag with a file name, we can list all commits for a file beyond renaming.

```bash
git log --name-only --pretty=format:%H --follow README.md
4e57c5d46637286731dc7fbb1e16330f1f3b2b7c
README.md

56955ff027f02b57212476e142a97ce2b7e60efe
README.md

5abdc5106529dd246450b381f621fa1b05808830
README.md

ea885f458b0d525f673623f2440de9556954c0c9
README.rdoc
```

This command roped in a commit from when `README.md` used to be called `README.rdoc`. If you want to know about the *full* history of a file, this is the way to go.

## Show the diffstat Summary of a Commit

Use the `--stat` flag when running `git show` on a commit to see the `diffstat` summary of that commit. For instance, this is what I get for a recent commit to [delve](https://github.com/derekparker/delve):

```bash
git show 8a1f36a1ce --stat
commit 8a1f36a1ce71d708d1d82afbc2191de9aefba021
Author: Derek Parker <derek.parker@coreos.com>
Date:   Wed Jan 27 23:47:04 2016 -0800

    dlv: Flag to print stacktrace on trace subcommand

 cmd/dlv/main.go     | 45 ++++++++++-----------------------------------
 terminal/command.go |  7 +++++--
 2 files changed, 15 insertions(+), 37 deletions(-)
```

The following is a description of the `diffstat` program

> This  program  reads the output of diff and displays a histogram of the insertions, deletions, and modifications per-file.

See `man git-show` and `man diffstat` for more details.

## Single Key Presses in Interactive Mode

When staging changes in interactive mode (`git add -p`), you have a number of options associated with single keys. `y` is *yes*, `n` is *no*, `a` is *this and all remaining*, and so on.

By default, you have to press *enter* after making your selection. However, it can be configured for single key presses. Add the following to your git configuration file to enable single key presses for interactive mode:

```
[interactive]
    singlekey = true
```

## Staging Changes within Vim

I've always used git from the command line, but only recently have I started to leverage [fugitive.vim](https://github.com/tpope/vim-fugitive) to more quickly do some common git commands from within vim.

I mostly use *fugitive* to stage changes for committing. To stage entire files, you can view the repository status, `:Gstatus`, and then navigate up and down (`k` and `j`) tapping `-` next to the files to be staged (or unstaged).

I've started to use git's interactive mode for staging changes from the command line (`git add --patch`) more and more and recently wondered if the same thing can be accomplished with *fugitive*.

It turns out it's pretty simple to do so. Instead of tapping `-` next to a file you want to stage, you can tap `p` next to it and you will be immediately dropped into interactive mode for that file. Prepare the lines you want to stage (or, again, unstage) and save.


## Staging Stashes Interactively

The `-p` flag can be used with `git stash`, just as it is used with `git add`, for interactively staging a stash.

So, if you have changes in your working tree that you want to stash mixed with ones that you want to keep working with, then you can simply do:

```bash
git stash -p
```

Read more on [interactive staging](https://git-scm.com/book/en/v2/Git-Tools-Interactive-Staging).

## Stashing Only Unstaged Changes

If you have both staged and unstaged changes in your project, you can perform a stash on just the unstaged ones by using the `-k` flag. The staged changes will be left intact ready for a commit.

```bash
git status
On branch master
...
Changes to be committed:

    modified:   README.md

Changes not staged for commit:

    modified:   app/models/user.rb
```

```bash
git stash -k
Saved working directory and index state ...
```

```bash
git status
On branch master
...
Changes to be committed:

    modified:   README.md
```

## Stashing Untracked Files

Normally when stashing changes, using `git stash`, git is only going to stash changes to *tracked* files. If there are any new files in your project that aren't being tracked by git, they will just be left lying around.

You might not want these *untracked* files left behind, polluting your working copy. Perhaps these files change your projects functionality or affect the outcome of tests. You just want them out of the way, for the moment at least.

And so along comes the `-u` or `--untracked` flag.

```bash
touch new_file.js
git stash
  No local changes to stash
git stash -u
  Saved working directory and index state WIP ...
```

## The Alpha Commit

I like to read commit logs. Today I wanted to see the first commit on a project. Here's what I used:

```bash
git rev-list --max-parents=0 HEAD
```

*Show me the commits that led to* `HEAD` *in reverse chronological order; then limit that list to the commits with no parent.*

Here's a small modification, to show the entire commit rather than the SHA alone:

```
git show $(git rev-list --max-parents=0 HEAD)
```

## Undo a Git Mistake

`git reflog` is a record of your actions in Git. With this command, you can undo almost any Git mistake.

```bash
git reflog

4bd0090 HEAD@{0}: <bad place>
46bd839 HEAD@{1}: <bad place>
967e214 HEAD@{2}: <last good place>
46bd839 HEAD@{3}: <good place>
967e214 HEAD@{4}: <good place>

git reset --hard HEAD@{2}
```

## Untrack a Directory of Files Without Deleting

In [Untrack A File Without Deleting It](untrack-a-file-without-deleting-it.md), I explained how a specific file can be removed from tracking without actually deleting the file from the local file system. The same can be done for a directory of files that you don't want tracked. Just use the `-r` flag:

```bash
git rm --cached -r <directory>
```

## Untrack a File Without Deleting It

Generally when I invoke `git rm <filename>`, I do so with the intention of removing a file from the project entirely. `git-rm` does exactly that, removing the file both from the index and from the working tree.

If you want to untrack a file (remove it from the index), but still have it available locally (in the working tree), then you are going to want to use the `--cached` flag.

```bash
git rm --cached <filename>
```

If you do this, you may also consider adding that file to the `.gitignore` file.

## Update the URL of a Remote

I just changed the name of a Github repository. One of the implications of this is that the remote URL that my local git repository has on record is now out of date. I need to update it.

If I use `git-remote` with the `-v` flag. I can see what remotes I currently have.

```bash
git remote -v
origin  git@github.com:jbranchaud/pokemon.git (fetch)
origin  git@github.com:jbranchaud/pokemon.git (push)
```

Now, to update the URL for that remote, I can use `git remote set-url` specifying the name of the remote and the updated URL.

```bash
git remote set-url origin git@github.com:jbranchaud/pokemon_deluxe.git
```

If I check again, I can see it has been updated accordingly.

```bash
git remote -v
origin  git@github.com:jbranchaud/pokemon_deluxe.git (fetch)
origin  git@github.com:jbranchaud/pokemon_deluxe.git (push)
```

## Verbose Commit Message

Git allows you to display a *diff* of the staged changes in the commit message buffer. This gives you the opportunity to carefully craft your commit message in a way that accurately describes the changes being committed. To display the diff when committing:

```bash
git commit -v
```

## Viewing a File on Another Branch

Sometimes you want to view a file on another branch (without switching branches). That is, you want to view the version of that file as it exists on that branch. `git show` can help. If your branch is named `my_feature` and the file you want to see is `app/models/users.rb`, then your command should look like this:

```bash
git show my_feature:app/models/users.rb
```

You can even tab-complete the filename as you type it out.

See `man git-show` for more details.

## What Changed?

If you want to know what has changed at each commit in your Git history, then just ask `git whatchanged`.

```bash
git whatchanged

commit ddc929c03f5d629af6e725b690f1a4d2804bc2e5
Author: jbranchaud <jbranchaud@gmail.com>
Date:   Sun Feb 12 14:04:12 2017 -0600

    Add the source to the latest til

:100644 100644 f6e7638... 2b192e1... M  elixir/compute-md5-digest-of-a-string.md

commit 65ecb9f01876bb1a7c2530c0df888f45f5a11cbb
Author: jbranchaud <jbranchaud@gmail.com>
Date:   Sat Feb 11 18:34:25 2017 -0600

    Add Compute md5 Digest Of A String as an Elixir til

:100644 100644 5af3ca2... 7e4794f... M  README.md
:000000 100644 0000000... f6e7638... A  elixir/compute-md5-digest-of-a-string.md

...
```

This is an old command that is mostly equivalent to `git-log`. In fact, the man page for `git-whatchanged` says:

> New users are encouraged to use git-log(1) instead.

The difference is that `git-whatchanged` shows you the changed files in their raw format which can be useful if you know what you are looking for.

See `man git-whatchanged` for more details.

## What is the Current Branch?

This question can be answered with one of git's plumbing commands, `rev-parse`.

```bash
git rev-parse --abbrev-ref HEAD
```

The `--abbrev-ref` flag tells `git-rev-parse` to give us the short name for `HEAD` instead of the SHA.

## Whitespace Warnings

You can configure git to warn you about whitespace issues in a file when you diff it. This is handled by the `core.whitespace` configuration. Add the following to your `.gitconfig` file:

```bash
[core]
      whitespace = warn
```

By default, git will warn you of trailing whitespace at the end of a line as well as blank lines at the end of a file.
