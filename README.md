# go-repl

> Quickly prototype some Go code without having to setup a temporary project, a
> file or even open your editor

## Features

- remember previously entered lines
- multiline support by default
- import all go standard library packages and use them
- inline package imports, no need to keep em at the top of your input

## Installation

```shell
go install github.com/xnacly/gorepl@latest
```

## Usage

`gorepl` ships a repl and allows the execution of inputs via the `-c` cli flag.

### Cli flag

```bash
$ gorepl -c 'println("Hello World")'
Hello World
```

### Repl

> The repl assumes multiline mode by default, end your input with ';' to execute it

#### Hello world

```text
$ gorepl
go > println("Hello World");
Hello World
```

#### Using packages

```text
$ gorepl
go > import "os"
>>>> cache, err := os.UserCacheDir()
>>>> import "fmt"
>>>> if err != nil {
>>>>   fmt.Println("Failed to get cache dir", err)
>>>>   return
>>>> }
>>>> fmt.Println(cache)
>>>> ;
/home/teo/.cache
go >
```

## Why

I wanted a go repl to quickly prototype something or check if the behaviour of
slices still match the idea i have in my mind :^).
