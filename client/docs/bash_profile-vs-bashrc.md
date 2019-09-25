---
title: bash_profile vs bashrc
lang: en-US
tags:
  - bash
---

# {{ $page.title }}

`bash` can be started in interactive mode or non-interactive mode. It can also act as a login shell or a non-login shell.

`bash` is started in interactive mode by your terminal emulator and can also be started in interactive mode like this:

```bash
bash
bash -i
bash -ic 'echo Hello!'
```

When you run a script through `bash` or if you start it with the `-c` option, it will run in non-interactive mode:

```bash
bash script.sh
bash -c 'echo Hello!'
```

`bash` is instructed to act as a login shell when you first log in to your machine or when you start `bash` with the `--login` or `-l` option.

`.bashrc` is sourced on every start in interactive mode when `bash` does not act as a login shell.

`.bash_profile` is only sourced when `bash` is started as an interactive login shell, or as a non-interactive shell with the `--login` option.

This means that `.bash_profile` is great for commands that should run only once and `.bashrc` for commands that should run every time you start a new shell.

For example, `PATH` customization should only happen once, since it is not an idempotent operation. Suppose something like this was in your `.bashrc`:

```bash
export PATH="$PATH:/addition"
```

Running these commands from a newly started interactive shell would produce the output below:

```bash
bash
bash
echo "$PATH"
/original_path:/addition:/addition:/addition
```

Setting `PATH` in `.bash_profile` alleviates this problem.
