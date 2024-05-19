# GO PATH

This contains a basic application intended to demonstrate how to use the `GOPATH` for package management.

## Dependencies

* [go version < 1.11](https://go.dev/doc/manage-install)
* [nix package manager (optional)](https://nixos.org/download/)

For this example you will need a version of `go` before version 1.11.  The GOPATH was deprecated in version 1.11 in favor of [go modules](). You could install this version directly on your machine or in their virtual environment manager of choice, but installing it directly on the host machine may create conflicts if you have other versions of go installed, so it is recommended that you use an environment manager.

This example contains within it a `shell.nix` file which is used by the [nix package manager](https://nixos.org/). Nix is a tool for creating virtual environments independent of the global configurations. 

If you have nix installed, then you can run the following command within this directory:

```sh
$ nix-shell
```

This will install `go` version 1.9. To confirm it works you can run: 

```sh
$ go version
```

If everything worked as expected you should see an output similar to this:

```sh
go version go1.9 linux/amd64
```

Note that the go version installed in the nix-shell environment will only be available while in th nix-shell. To exit the nix shell, either just close the terminal or type exit to return to the standard terminal environment.

## Steps

Here I will walk through the steps for loading dependencies using the `GOPATH`. In the [main.go](./main.go) you will see that we are importing an example package located in [this repository](https://github.com/thomasbellio/learn-go-with-tests).

1. First we are going to set the GOPATH environment variable to the current working directory (this will assume you are in the same directory as this readme)

```sh
$ export GOPATH=$(pwd)
```

2. Now that we have set the GOPATH lets see what it looks like when we attempt to run the application in main.go. You'll notice an error on the console when you run the following command:

```sh
$ go run main.go

main.go:5:2: cannot find package "github.com/thomasbellio/learn-go-with-tests/iteration" in any of:
        /nix/store/bgcfs9vflm5fiis89jmd9g7f065q8bdg-go-1.9/share/go/src/github.com/thomasbellio/learn-go-with-tests/iteration (from $GOROOT)
        /path/to/this/repository/examples/dependency-management/go-path/src/github.com/thomasbellio/learn-go-with-tests/iteration (from $GOPATH)

```

We see this error because although we have set the `GOPATH`, we have not yet downloaded the dependency. In the next step we will use the `go get` command to download the package

3. We can use the `go get` command to install the dependency into the `GOPATH`


```sh
$ go get github.com/thomasbellio/learn-go-with-tests/iteration
```

The `go get` command you will download the repository into the `GOPATH` in a folder called `src` (technically you can have multiple paths in the GOPATH in this case it will download into the first path found in the GOPATH). The GOPATH will have the following structure:

```sh
.
└── github.com
    └── thomasbellio
        └── learn-go-with-tests
            ├── arrays
            │   ├── sum.go
            │   └── sum_test.go
            ├── go.mod
            ├── hello-world
            │   ├── hello.go
            │   └── hello_test.go
            ├── integers
            │   ├── adder.go
            │   └── adder_test.go
            ├── iteration
            │   ├── repeat.go
            │   └── repeat_test.go
            └── shell.nix

```

`go get` will also create a pkg directory with the following structure:

```sh
└── linux_amd64
    └── github.com
        └── thomasbellio
            └── learn-go-with-tests
                └── iteration.a

```

The iteration.a file is a precompiled binary of the iteration package, this allows go to build your project without needing to re-compile dependencies, this is helpful for accelerating build times, if you don't wish to pre-compile the binary you can pass the `-d` flag to the go get command.

Notice that there is a folder called `linux_amd64`, this is the pre-compiled binary that for your specific OS and architecture. This folder name is determined by the following environment variables:

```sh
GOOS="linux"
GOARCH="amd64"
```

Depending on what type of operating system you are running in this folder the name may be different. 

4. Now that we have installed the dependencies we should now be able to successfully run the application:

```sh
$ go run main.go
# This should be the output if everything went well
gogogogogogogogogogo
```

## Conclusion

This has been a quick primer on how to use the `GOPATH` for dependency management for versions of `go` before version `1.11`. While this is a deprecated approach to managing dependencies it is good to understand how it works so that if you ever encounter a `go` project that is written based on an earlier version you will know what to expect. 

In [another section](../go-modules/README.md) we will discuss the alternative and modern approach of using go modules.
