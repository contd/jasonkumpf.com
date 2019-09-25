---
title: Today I Learned Git
lang: en-US
tags:
  - vcs
---

# {{ $page.title }}

## Use a file from another branch

Sometimes you just need one file from another branch. Sure you could `git cherry-pick` but then you're dealing with commits. That sort of thing gets sticky fast; don't go there.

The best way is to use your old pal `git checkout`. Just make sure you're on the branch you want to bring the file into and then checkout the file from its source branch. Here's the syntax.

```
git checkout my-awesome-source-branch the/path/to/yourfile.rad
```

That's it – enjoy.

## Split Up a Commit, Rewrite History

When working on a branch with multiple commits, you can "go back in time" and revise previous commits any way you please.

```bash
git rebase -i origin/master
```

This command will prompt you inside of your `$VISUAL` with a series of commit SHAs and commit titles

```bash
# ...
pick ed1ff41 Move templates
pick 274ac0e Move components & views
# ...
```

To split up `274ac0e`, replace `pick` with `edit` and save the buffer.

```bash
# ...
pick ed1ff41 Move templates
edit 274ac0e Move components & views
# ...
```

You are now detached from the `HEAD` of your branch and "back in time". To split up the current commit (`274ac0e`):

```bash
git reset HEAD~
```

This will unstage all files within the commit. Next, split up the commit any way you'd like.

```bash
git add app/views
git commit -m "Move views"
git add app/components
git commit -m "Move components"
```

When you're finished, continue the rebase.

```bash
git rebase --continue
```

If you introduced merge conflicts down the line, you'll have to resolve them. If all went well, your branch's history will be re-written.

```bash
git log origin/master..

# ...
ed1ff41 Move templates
7b84a01 Move views
274ac0e Move components
# ...
```

Keep in mind that you might have to force push your branch to `origin`, depending on whether or not your revised commits have been pushed.

```bash
git push origin my-branch-name --force
```

## Two ways of squashing commits

It is handy to squash down your commits before merging your PR with `my-new-cool-feature`. You can either squash them down by doing an interactive rebase like so:

```bash
git checkout my-new-cool-feature
git rebase -i master
```

This will open up your `$EDITOR` of choice and you are free to pick and choose which commits to squash together.

This might be tedious if you have a big number of commits to squash together. Very tedious. TIL that you can use `git merge` to squash your commits, all in one go.

```bash
git checkout master
git merge --squash my-cool-new-feature
```

This will leave all of your changes staged on `master`, ready to be committed as one.
