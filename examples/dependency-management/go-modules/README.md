# GO Modules

This contains a basic application intended to demonstrate how to use `go modules` for package management.

## Dependencies

* [go version 1.21.9](https://go.dev/doc/manage-install)
* [nix package manager (optional)](https://nixos.org/download/)

For this example you will need a version of `go`  greater than version 1.11, which is the version when `go` modules were introduced. This tutorial was written with `go` version 1.21.9, but it will probably work with any version post 1.11.  To install `go` you can install it directly on the host machine using the instructions linked from above or you can use a virtual environment manager.

This example contains within it a [shell.nix](./shell.nix) file which is used by the [nix package manager](https://nixos.org/). Nix is a tool for creating virtual environments independent of the global configurations. 

If you have nix installed, then you can run the following command within this directory:

```sh
$ nix-shell
```

This will install `go` version 1.21.9. To confirm it works you can run: 

```sh
$ go version
# If everything worked as expected you should see an output similar to this:
go version go1.21.9 linux/amd64
```

Note that the `go` version installed in the nix-shell environment will only be available while in th nix-shell. To exit the nix shell, either just close the terminal or type exit to return to the standard terminal environment.

## Steps

Here I will walk through the steps for managing package dependencies using `go` modules. In the [main.go](./main.go) you will see that we are importing an example package located in [this repository](https://github.com/thomasbellio/learn-go-with-tests).

1. First lets see what happens when we attempt to run the main.go file
```sh
$ go run main.go
# you will see this error output from below
main.go:5:2: no required module provides package github.com/thomasbellio/learn-go-with-tests/iteration: go.mod file not found in current directory or any parent directory; see 'go help modules'
```
2. To solve the error from above we first need to initialize the `go module` by running this command:
```sh
$  go mod init 'github.com/thomasbellio/engineering-blog/examples/go/dependency-management/go-modules'
# you will see output similar to this
go: creating new go.mod: module examples/dependency-management/go-modules
go: to add module requirements and sums:
        go mod tidy
```

The command from above will initialize our module by creating  a `go.mod` file in the current working directory. Its contents will look something like this


```go
# This is the name of the go module that we passed into the init command.
# Technically we could name the module anything we want, but for modules that you may want to load from other applications you can include the github repository url.
module github.com/thomasbellio/engineering-blog/examples/go/dependency-management/go-modules

go 1.21.9
```
The `go.mod` file includes the name of the module, along with the the version of go that this module relies upon. Next we are going to add the requirements using the `go tidy` command.

```sh
go mod tidy
```

The `go mod tidy` command will do a coupl of things. First it adds a new entry into the go.mod file adding all dependencies referenced in the project with the require directive. Now the `go.mod` file looks like this:

```sh
module github.com/thomasbellio/engineering-blog/examples/go/dependency-management/go-modules

go 1.22.2

toolchain go1.22.3

require github.com/thomasbellio/learn-go-with-tests v0.0.0-20240513230150-bdd12f7937b4

```
It also downloads the dependencies into the path set by the `GOMODCACHE` environment variable. The default location for GOMODCACHE is:

```sh
$HOME/go/pkg/mod/
```
You can update this default path using the `GOPATH` variable.

The `go mod tidy` command also adds a file called `go.sum`, it will looks something like this:

```sh
github.com/thomasbellio/learn-go-with-tests v0.0.0-20240513230150-bdd12f7937b4 h1:ti0WhMtl3P2x2xzTJZdtdhDeal/DVhIdPvyJAGPE1M0=
github.com/thomasbellio/learn-go-with-tests v0.0.0-20240513230150-bdd12f7937b4/go.mod h1:1zW/1kFT3vx9SM9b40y6cZE9udd235BMSnemj9dE4kY=
```

The `go.sum` file contains the hashes of the downloaded packages and modules, which ensures that the package hasn't been modified from when it was originally downloaded.

3. Finally now that we have the dependencies downloaded we can run our program:

```sh
$ go run main.go
# you will see an output that repeats the word 'go' 10 times
gogogogogogogogogogo
```

## Conclusion

This is just a primer on go modules there are more facets including vendoring and versioning, perhaps I will include a separate post and tutorial on those aspects of dependency management. 




