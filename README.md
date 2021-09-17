<h1 align="center">cd by golang</h1>
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
➜  mini-balance git:(master) ✗
```

😋

<h2 align="center">Features</h2>

✔️ 🐬 Intelligent and ergonomical matching. Designed with an emphasis on ergonomics. Your search comfort and speed is a priority.

✔️ 🚀 Speed! Powered by Golang.

✔️ 🚀 Speed!! Histories will be stored in a db file (`~/cdi-db-shortcuts.json`) for search speed.

✔️ 🚀 Speed!!! node_modules and other black hole directories wont be searched.

<h2 align="center">Download</h2>

[Download cdi exe](https://github.com/legend80s/cdi-go/raw/master/cdi) and make it executable:

```sh
chmod +x ~/where-cdi-cmd/cdi && xattr -c ./cdi
```

<h2 align="center">Usage</h2>

1 Add the shell functions to your `.zshrc` because you cannot change shell directory in golang process.

```sh
# cdi begin
cdi() {
  cd $(~/where-cdi-cmd/cdi -fallback "$@")
}

# show debug info
cdi-echo() {
  echo $(~/where-cdi-cmd/cdi "$@")
}

# show cache
alias cdi-stat="~/where-cdi-cmd/cdi stat"
# clear cache
alias cdi-stat-clear="~/where-cdi-cmd/cdi stat --clear"
# cdi end
```

2 Then suppose we have these directories in `~/workspace/`

```txt
cli-aid
commit-msg-linter
gallery-server
js2schema
```

3 just use `cdi` instead of the builtin `cd` command

```sh
$ cdi ca
```

will `cd` into `~/workspace/legend80s/cli-aid`

### Ergonomically

> Ergonomics or human engineering, science dealing with the application of information on physical and psychological characteristics to the design of devices and systems for human use.

Weight or priority from highest to lowest:

1. **Full** match basename: `cdi js2schema` equal to `cd ~/workspace/legend80s/js2schema`
2. **Prefix** match: `cdi cli` equal to `cd ~/workspace/legend80s/cli-aid`
3. **Abbr** match: `cdi ca` equal to `cd ~/workspace/legend80s/cli-aid`
4. **Contains word** match: `cdi msg` equal to `cd ~/workspace/legend80s/commit-msg-linter`

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
  target=$(~/where-cdi-cmd/cdi "$@")

  echo $target

  if [[ $target == *"no such dirname"* ]]; then
    # DO NOTHING WHEN NO DIRECTORY FOUND
  else
    code $(cdi-echo $1)
  fi
}
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
