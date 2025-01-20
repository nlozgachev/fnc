# fnc

It is a pet-project to help me with my daily tasks.
Currently just helps to create Git [commits](https://www.conventionalcommits.org) and [branches](https://conventional-branch.github.io) in a conventional way.
Will think about adding more features in the future.

```
Usage:
  fnc <command>

Available commands:
  branch   Create a new branch
  commit   Create a new commit
  help     Show help
  version  Show the application version
```

## Recommended installation

You should setup your shell in order to use the executables.

Basic installation steps for UNIX-like systems:

1. Create `~/.bin` directory.

```bash
mkdir ~/.bin
```

2. Add it to `PATH` of your shell.

```bash
# Inside `.bashrc` for Bash or `.zshrc` for Zsh
PATH="$PATH:~/.bin"

# For Fish users
fish_add_path $HOME/.bin
```

3. Make scripts executable.

```bash
chmod u+x fnc
```

4. Create `~/.config/fnc/config.json` file. It is required for correct handling of branch prefixes during their creation. The example config can be found inside [`config.example.json`](./config.example.json).

5. Relaunch your shell and use the scripts.

## Examples

```
fnc branch

Select task type:
1. feature
2. bugfix
3. hotfix
4. release
5. chore
Enter choice (1-5): 1

Enter task ID (optional): 1234

Enter description (required): made something cool

Switched to a new branch 'feature/ABC-1234_made_something_cool'

Branch "feature/ABC-1234_made_something_cool" created from "develop"
```

```
fnc commit

Select the type of change you're committing:
1. fix
2. feat
3. chore
4. docs
5. style
6. refactor
7. perf
8. test
9. ci
10. build
Enter choice (1-10): 2

Enter the commit message (required): very important feature

Enter the commit body (optional): this will be in the next release!

[feature/ABC-1234_made_something_cool 8c32e2f] feat: very important feature (ABC-1234)

# other git output here

Commit successful.
```
