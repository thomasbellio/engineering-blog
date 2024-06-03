# About 

This section contains a basic implementation of [fizz-buzz](https://leetcode.com/problems/fizz-buzz/), which is a basic program implementation. The purpose of this code is to demonstrate how to leverage the debugger. 

## Dependencies

* [go version 1.21.9](https://go.dev/doc/manage-install)
* [nix package manager (optional)](https://nixos.org/download/)
* [dlv](https://github.com/go-delve/delve/tree/master)

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

You will also need to [install dlv](https://github.com/go-delve/delve/tree/master/Documentation/installation). 

```sh
$ go install github.com/go-delve/delve/cmd/dlv@latest
```

Assuming everything went as expected you should be able to run:

```sh

$ dlv version
# This will be the output
Delve Debugger
Version: 1.22.1
Build: $Id: 0c3470054da6feac5f6dcf3e3e5144a64f7a9a48 $
```



