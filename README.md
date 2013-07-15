# Planet 0.0

## Hacking in GO:

1. create some folder in your user space, call it something like `go-projects`, or `go-libraries` and export it under the `GOPATH` environment variable.
For example, in my computer I have it exported like this:

```bash
export GOPATH=$HOME/projects/go-libraries
```

2. once you set GOPATH, everytime you install some library with `go
get url/to/library` it will be installed under GOPATH.  For more
information about how `Go` install and keep track of projects, click
[here](http://golang.org/doc/code.html#Organization) and learn how to
organize go code.

3. Install "planet"

```bash
go get github.com/gabrielfalcao/planet
```

## Running tests

```bash
cd /path/to/gopath/src/github.com/gabrielfalcao/planet
make test
```

## Running the binary

```bash
go run /path/to/gopath/src/github.com/gabrielfalcao/planet/src/main.go
```
