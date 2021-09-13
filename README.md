# cdi

> `C`hange Current working `D`irectory `I`ntelligently.

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

## Features

- Intelligent matching. *node_modules wont be searched*.
- Speed. Histories will be stored in a db file (`~/cdi-db-shortcuts.json`) for search speed.

## Download

[Download cdi exe](https://github.com/legend80s/cdi-go/raw/master/cdi) and make it executable:

```sh
chmod +x ~/where-cdi-cmd/cdi
```

## Usage

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

### `cdi` Match Priority from Highest to Lowest

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

## Testing

```sh
go test ./utils
```

## Build

```sh
go build -v -o cdi
```

## Publish

```sh
git tag v1.x.x && gp && gp --tags
```
