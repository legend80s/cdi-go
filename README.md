# cdi

> Search and **C**hange Current working **D**irectory **I**ntelligently.

Use `cd`

```sh
➜  topup-center git:(master) ✗ cd mb
cd: no such file or directory: mb
```

Use `cdi`

```sh
➜  topup-center git:(master) ✗ cdi mb
➜  mini-balance git:(master) ✗
```

😋

## Features

- Full name and abbr searching supported. *wont search node_modules*.
- Speed. History will be stored in a db file (`~/cdi-db-shortcuts.json`) for search speed.

## Usage

1 Add this shell function to your `.zshrc` because you cannot change the shell directory in golang process.

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

2 Then suppose we have the dir list in `~/workspace/legend80s`

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

### other cdi usage

1. **Full** match the basename:`cdi js2schema` equal to `cd ~/workspace/legend80s/js2schema`
2. **Abbr** match: `cdi ca` equal to `cd ~/workspace/legend80s/cli-aid`
3. **Contains** match: `cdi lint` equal to `cd ~/workspace/legend80s/commit-msg-linter`

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
