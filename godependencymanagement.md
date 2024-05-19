## Introduction

Over the last several days I have been doing a deep dive into how `go` manages dependencies, I am relatively new to writing code in `go`, but I have been building, deploying, and supporting large scale systems across several different programming languages including PHP, Python, and C# for nearly 15 years. I wanted to start learning about dependency management early on in my journey with `go`, because to me it is one of the most important concepts for building systems that operate at any non-trivial scale.

When learning a new language most tutorials start with the basics of data structures and control flow. These things are, of course, important parts of any programming language, but the details of these constructs are, save for the ocassional syntactical oddities, very much the same in most languages. Control flow and data structures are the core building blocks of programming, but once you understand the basics in one language it is relatively easy to port over these ideas into new languages (I will call out that functional languages like Haskell or Erlang may be exceptions to this rule, but I don't have enough experience with these languages to speak confidently). To me one of the hardest parts of learning a new language is learning about the bigger eco-system. One of the biggest components of any language eco-system is dependency management. Understanding this part will make it much easier to build useful systems.

If you are new to programming then beginning with the basics of control flow and data structures is the ideal place to start, but if you already have experience programming I think it is more important to learn about the overall eco-system from the beginning. 

Before I get into the technical details of how dependency management works in `go`, I want to briefly discuss the significance of  dependency management more broadly and how `go's` system of dependency management may be different from other languages. 

## Dependency Management

All modern applications rely on 3rd party integrations to make an application work and run. This is for good reason, because if we didn't have the ability to leverage 3rd party implementations programmers would spend most of their time re-inventing implementations of common paradigms. Dependency management is a challenging problem because each dependency has its own version and each version of a dependency has its own dependencies and each of those dependencies have their own dependencies. Sometimes two dependencies may have dependencies on the same package but with different versions. Furthermore by its nature a dependency management system relies on pulling those dependencies from remote sources which may or may not be consistent or secure. Creating reproducible builds that produce the same artifact for a given revision can be challenging and present some key challenges related to security and stability. It all quickly becomes very complicated and evokes the famous 'turtles all the way down' aphorism.

Some examples of dependency management tools in some popular languages are:

* Node.js -> npm
* PHP -> composer
* Java -> maven
* C# -> nuget
* Python -> pip

One of the interesting things about the list above is that each of these languages have a dependency management tool that is separate and apart from the underlying programming language. 

This can have some advantages as the language itself may be agnostic to how dependencies and indeed environments are managed. Composer, for example, is a package management frameworkd written in PHP, but there is nothing about writing in PHP that expects or requires the use of Composer. This is also true for nodejs, which has as its most common dependency management tool `npm`, but there are other implementations like yarn that is used by many projects. Don't get me started on Python, there are a litany tools and approaches to managing python dependencies, but `pip` is the most common.

## Dependency Management in Go

When we compare `go` to other languages we start to see how it differs, because `go`, unlike many other programming languages, is very opinionated about how dependencies should be managed. As I have been learning `go` it quickly becomes obvious that it has 'convention over configuration' as one of its core design heuristics. Dependency management in `go` relies very heavily on certain conventions around the naming and structure of dependencies and unlike other languages these conventions are built directly into the language itself rather than relying on secondary tools. This is a pattern that I have noticed throughout the `go` eco-system, and not just in the case of dependency management. For example `go` only exports package function names written in [pascal case](https://en.wiktionary.org/wiki/Pascal_case#English) while [camelcase](https://en.wikipedia.org/wiki/Camel_case) functions are considered private.

Following a conventions approach, as `go` does, comes with some key advantages and one of those advantages is that it enforces consistency from module to module. A conventions based approach has some disadvantages also with the biggest disadvantage, in my opinion, being that it requires one to have knowledge of those conventions and sometimes those conventions can become so esoteric that it creates a barrier to learning and adoption. Overall though I think the consistency created by a conventions based approach outweighs the disadvantages.

One such convention is the convention of including the remote location of the dependency in the package name itself. For example assume that we wanted to import the compare library from google into our application. To do this we would add it to the `import` section of our application like this:

```go
package main

import (
    "fmt"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

Notice the name of the package `github.com/google/go-cmp/cmp` includes the url for github where the package resides. When we install the dependencies `go` knows that it has to retrieve the package from github.com, and it also knows how to download a `git` repository. When I first saw this convention it felt a bit weird, but it actually makes a lot of sense, this saves us from having to configure a packages remote location independently as we might have to do in a dependency manager like `pip` in the case of Python.

### GOPATH

In versions of `go` before `1.11` the location of dependencies was found in the `GOPATH`, which is very similar to how Python leverages `PYTHONPATH`, for a given workspace one would set the `GOPATH` and `go` would look for dependencies in that location on the local file system. Dependencies were not explicitly set anywhere in particular it would rely on the use of the `go get` command to install dependencies, and perhaps the dependencies might be set in a `make` file or some other build structure, to indicate which dependencies existed in the project.

I have created a small tutorial within this repository that walks through how to [use the GOPATH](./examples/dependency-management/go-path/README.md).

## Go Modules
Starting in `go` version 1.11, the concept of `go modules` was introduced and replaced the now deprecated GOPATH  approach to dependency management. For those who have used to tools like `composer` and `npm` this approach will seem more familiar, as it explicitly lists out the dependencies of the module in a file called: `go.mod`. The `go.mod` file might be comparable to the package.json file in `npm` where dependencies and the specific version of `go` is specified. 

```go

module github.com/thomasbellio/engineering-blog/examples/go/dependency-management/go-modules

go 1.22.2

toolchain go1.22.3

require ( 
    github.com/thomasbellio/learn-go-with-tests v0.0.0-20240513230150-bdd12f7937b4
)
```

With the `go` directive you specify the specific version of go that this module relies upon along with the version of the `go` toolchain. The `require` directive lets you specify a list of dependencie for the module. Overall this approach is more similar to other package management tools in other languages and makes the dependencies explicit within the module.I have created a small tutorial within this repository that walks through how to [use go modules](./examples/dependency-management/go-modules/README.md)


## Conclusion

Certainly with the introduction of `go` modules, dependency management in `go` should seem more fammiliar to those that have used package management tools in other languages. I have just touched on what I believe to be the most salient points of `go` dependency management, but look for other posts as I dive deeper into other dependency management concepts.



