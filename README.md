<h1 align="center">cd by golang <img width="6%" src="https://golang.org/lib/godoc/images/go-logo-blue.svg"></img></h1>
<p align="center"><img width="30%" src="https://golang.org/lib/godoc/images/footer-gopher.jpg"></img></p>

> `C`hange Current working `D`irectory Fast 🚀, `I`ntelligently 🐬.
>
> And Ergonomically 🦄 in the Aspect of Human Searching Preferences.
>
> `cd` 命令 golang 进阶版 `cdi`。符合人体工程学搜索习惯的 `cd` 命令

Use `cd`

```sh
$ cd mb
cd: no such file or directory: mb
```

😡

Use `cdi`

```sh
$ cdi mb
/Users/legend80s/workspace/mini-balance
➜  mini-balance git:(master) ✗
```

😋

- [🍬 Features](#-features)
- [Download](#download)
- [Usage](#usage)
- [More interesting usage 🤠](#more-interesting-usage-)
- [Development](#development)
- [Testing](#testing)
- [Build](#build)
- [Publish](#publish)
- [Version History](#version-history)
- [What does the `install.sh` do](#what-does-the-installsh-do)

<h2 align="center">🍬 Features</h2>

✔️ 🐬 Intelligent and ergonomical matching. Designed with an emphasis on ergonomics. Your search comfort and speed is a priority.

✔️ 🚀 Speed! Powered by Golang.

✔️ 🚀 Speed! Histories will be stored in a db file (`~/cdi-db-shortcuts.json`) for search speed.

✔️ 🚀 Speed! node_modules and other black hole directories wont be searched.

<h2 align="center">Download</h2>

To **install** `cdi`, you should run the install script:

```sh
cd ~ && git clone --depth 1 https://github.com/legend80s/cdi-go.git && sh ~/cdi-go/scripts/install.sh ~/cdi-go/cdi-v5 && source ~/.zshrc
```

To **update**, you should run the update script:

```sh
cd ~ && git clone --depth 1 https://github.com/legend80s/cdi-go.git && sh ~/cdi-go/scripts/install.sh ~/cdi-go/cdi-v5 update
```

To **uninstall** just remove cdi and codi functions in `~/.zshrc`.

*If the download takes too long, jump to FAQ [Install mannually](#install-manually).*

<h2 align="center">Usage</h2>

Suppose we have these directories in `~/workspace/`

```txt
cli-aid
commit-msg-linter
gallery-server
js2schema
```

Just use `cdi` instead of the builtin `cd` command

```sh
$ cdi ca
```

will `cd` into `~/workspace/cli-aid`

### Ergonomically

> Ergonomics or human engineering, science dealing with the application of information on physical and psychological characteristics to the design of devices and systems for human use.

Weight or priority from highest to lowest:

1. **Full** match basename: `cdi js2schema` equal to `cd ~/workspace/js2schema`
2. **Prefix** match: `cdi cli` equal to `cd ~/workspace/cli-aid`
3. **Abbr** match: `cdi ca` equal to `cd ~/workspace/cli-aid`
4. **Contains word** match: `cdi msg` equal to `cd ~/workspace/commit-msg-linter`

Suppose we have these directories in `~/workspace/`:

```txt
dir-lab
├── amuser-low-info
├── long
│   ├── ali
│   ├── ali-test
│   ├── alitest
│   ├── hello-ali-test
│   └── hello-alitest
└── long-long-long-long-long
    └── ali
```

`cdi ali` will match:

```
dir-lab
├── amuser-low-info ✅
├── long
│   ├── ali ✅
│   ├── ali-test ✅
│   ├── alitest ✅
│   ├── hello-ali-test ✅
│   └── hello-alitest ❌
└── long-long-long-long-long
    └── ali ✅
```

After sorted by priority then by length:

#### Full match

✅ dir-lab/long/ali

✅ dir-lab/long-long-long-long-long/ali

#### Prefix

✅ dir-lab/long/alitest

✅ dir-lab/long/ali-test

#### Abbr

✅ dir-lab/amuser-low-info

#### Contains word

✅ dir-lab/long/hello-ali-test

So the best match is `dir-lab/long/ali`.

### Advanced Usage

### 1. Set search dir

The default search dir is `~/workspace`, change it to `work`:

```sh
cdi set-root ~/work
```

### 2. List saved shortcuts

```sh
cdi stat
```

Outputs file content in `~/cdi-db-shortcuts.json`

### 3. Clear saved shortcuts

```sh
cdi clear
```

### 4. Force search the dir tree instead of from cache

```sh
--walk
```

Example:

```sh
cdi --walk balance
```

<h2 align="center">More interesting usage 🤠</h2>

### Make your VSCode `code` command much more intelligent and convenient!

You can

```sh
codi dir-you-want-to-open-in-vscode
```

instead of executing 2 commands

```sh
cd ~/the/long/long/and/hard-to-memorize-dir-that-you-want-to-open-in-vscode

code .
```

Put this in your ~/.zshrc:

```sh
# Intelligent `code` command `codi`
codi() {
  target=$($cdipath "$@")

  echo $target

  if [[ $target == *"no such dirname"* ]]; then
    # DO NOTHING WHEN NO DIRECTORY FOUND
  else
    code $(cdi-echo $1)
  fi
}
```

<h2 align="center">Development</h2>

```sh
go run main.go -walk {directory-to-cd}
```

<h2 align="center">Testing</h2>

```sh
go test ./utils
```

<h2 align="center">Build</h2>

```sh
go build -v -o cdi
```

<h2 align="center">Publish</h2>

```sh
git tag v1.x.x && gp && gp --tags
```

<h2 align="center">Version History</h2>

### [V1 ☞](https://github.com/legend80s/cdi-go/raw/master/cdi-v1)

Match against without priority

- full match
- abbr
- abbr contains
- basename letter contains

Then pick the shortest one.

### V2

> Priority or weight should be taken into account.

Match priority

1. full match
2. abbr
3. abbr contains
4. basename letter contains

Then pick the one with highest priority if priority equals then pick the shortest one.

### V3

Match priority

1. full match
2. prefix
3. abbr
4. basename **word** contains

Then pick the one with highest priority if priority equals then pick the shortest one.

### [V4 ☞](https://github.com/legend80s/cdi-go/raw/master/cdi-v4)

> Directory nested level should be taken into account.

Match priority

1. full match
2. prefix
3. abbr
4. basename word contains

Then conditionally sort by length or directory nested level.

  1. If the directories take the same priority , we pick the shortest one.
  2. Otherwise we pick the directory with least nested level, instead of the the one with higher priority but has more nested level greater than **1**.

#### Directory nested level sorting

> We choose the directory with least nested level when the counterparts's level is too deep. "too deep" means level diff > 1

For example:

```plaintext
/Users/name/workspace/6/paytm/mr # (priority 0, directory nested level 6)
/Users/3/MiniRecharge # (priority 2, directory nested level 3)
```

When we `cdi mr`, though the first path `paytm/mr` is matched `mr` by "full match" with highest priority but its nested level is too deep (6).

But the second path `/Users/3/MiniRecharge` matched by "abbr" though with much lower priority but has lower nested level 3.

`6 - 3 = 3 > 1`. So `cdi mr` will `cd` into the second path `/Users/3/MiniRecharge`.

In V2 or V3 it will `cd` into `/Users/name/workspace/6/paytm/mr` because V3 always pick the one with highest priority.

In V1 it will `cd` into `/Users/3/MiniRecharge` because V1 always pick the shortest path.

### [V5 ☞](https://github.com/legend80s/cdi-go/raw/master/cdi-v5)

> 返璞归真，简单即美

Match against without priority

- full match
- prefix
- abbr
- basename word contains

Then pick the **least nested level** directory.

## FAQ

### Install manually

1. Download the cdi cmd https://raw.githubusercontent.com/legend80s/cdi-go/master/cdi-v5.

2. Make it executable:

   ```sh
   chmod +x ~/path/to/downloaded/cdi && xattr -c ~/path/to/downloaded/cdi
   ```

3. Copy the script in [what-does-the-installsh-do](#what-does-the-installsh-do) to you `~/.zshrc`.

4. Make it alive:

   ```
   source ~/.zshrc
   ```

5. Give it a try:

   ```shell
   cdi whatever-directory-you-want-to-jump-into
   codi whatever-directory-you-want-to-open-in-vscode
   ```

### What does the `install.sh` do

It will add the shell functions to your `.zshrc` because you cannot change shell directory in golang process.

```sh
## --- cdi begin ---
cdipath="~/path/to/downloaded/cdi"

cdi() {
  target=$($cdipath -fallback "$@")

  echo $target
  cd $target
}

# Show debug info
cdi-echo() {
  target=$($cdipath "$@")

  echo target
}

# Intelligent `code` command `codi`
codi() {
  target=$($cdipath "$@")

  echo $target

  if [[ $target == *"no such dirname"* ]]; then
    # DO NOTHING WHEN NO DIRECTORY FOUND
  else
    code $(cdi-echo $1)
  fi
}

# Show cache
alias cdi-stat="$cdipath stat"
# Clear cache
alias cdi-stat-clear="$cdipath stat --clear && echo -n 'Clear cache success. ' && cdi-stat"
## --- cdi end ---
```
