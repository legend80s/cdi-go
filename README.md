# cdi

> Search and **C**hange Current working **D**irectory **I**ntelligently.

## Features

- Full name and abbr searching supported. *wont search node_modules*.
- Speed. History will be stored in a db file (`~/cdi-db-shortcuts.json`) for search speed.

## Usage

1 Add this shell function to your `.zshrc` because you cannot change the shell directory in golang process.

```sh
to() {
  cd $(cdi "$1")
}
```

2 Then suppose we have the dir below

```txt
~/workspace/legend80s/

cli-aid
commit-msg-linter
gallery-server
js2schema
```

3

```sh
$ to ca
```

will cd into `~/workspace/legend80s/cli-aid`

### cd

cd ~/workspace/legend80s/js2schema

```sh
to js2schema
```

cd ~/workspace/legend80s/cli-aid

```sh
to ca
```

cd ~/workspace/legend80s/commit-msg-linter

```sh
to cml
```

### Advanced Usage

### 1. Set base search dir

The default base search dir is `~/workspace`, change it to `work`:

```sh
cdi set --base-dir ~/work
```

### 2. List saved shortcuts

```sh
cdi stat
```

utputs file content in `~/cdi-db-shortcuts.json`

### 3. force search the dir tree

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
