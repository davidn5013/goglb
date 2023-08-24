# goglb 

goglb lists all global variable and constans in golang module 

## Install

go install github.com/davidn5013/goglb@latest

## Usage

```bash
Go tool for listing global variables (and constans) in go module path
  -path string
        path to module (default ".")
  -varconst
        list global variabel and constrants
```

In module path run:

```bash
> goglb -varconst
```

```bash
Globals in build\build.go:
Variables:
Version
Time
User
RootPathPrefix
Constants:
Globals in cmd\gdu\app\app.go:
Variables:
Constants:
...
```
