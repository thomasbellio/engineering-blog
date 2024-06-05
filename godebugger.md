# Debugging With Go and Neovim

It has been a while since I have used a debugger for software development in fact the last time I used a debugger regularly was when I was a .Net developer using Visual Studio Pro. Since that time I have been getting by using print statements; but debuggers can be helpful especially when one is attempting to explore a code base that is unfamiliar. 

I have recently been attempting to do all of my development and overall work using neovim, because I want to be leet af ;).  In addition to being leet af, there is a value in using neovim because it really forces you to abandon the mouse in favor of the keyboard, which in my experience is absolutely more productive once you get past the learning curve. The learning curve is not just related to the keyboard shortcuts, but also related to configuring the editor.

One of the nice things about using a tool like Visual Studio or Intellij is that it abstracts away much of the technical details of how things are operating under the hood, but over the years I have learned the value of understanding more of the details especially when you are trying to build enterprise systems. There is, of course, a balance that must be struck between tinkering with tools and actually getting work done. I have historically been reluctant to use Vim, because I didn't want to spend a lot of time configuring my tools, but `neovim` really does make things a lot easier to configure and understand.

As I have been learning how to program in Go, I really wanted to improve my overall workflow and start from the ground up with a tool like neovim and part of that is understanding how to leverage tools like debuggers. I was recently able to `delve` into (don't worry you will get the joke later) the nitty gritty details of setting up debuggers in neovim and particularly for use with `go`.


In this pursuit I have learned about a new concept called  [The Debug Adapter Protocol](https://microsoft.github.io/debug-adapter-protocol/), the debug adapter protocol is an abstraction that facilitates the communication from your debug client which is basically your editor or IDE, in my case [neovim](https://neovim.io). Aside from the debug adapter protocol, which is really just an abstraction we need a debugger for our specific language, in this case we are focused on debugging `go`. Based on my reading and research the standard `go` debugger is called [Delve](https://github.com/go-delve/delve?tab=readme-ov-file) (now you get the joke). Delve is used to attach to or build `go` programs and enable debugging, but it also has an implementation of the Debug Adapter Protocol, so it can be ran as a DAP server.

In this post I am going to walk through:

1. The steps to use `Delve` to debug applications in `go`
2. How to integrate debugging into the neovim editor

You can see the full technical details of this implmentation in [the examples here](./examples/go-debugging/README.md). I hope whoever reads this finds it useful.

## Conclusion

Setting up tools like the debugger in neovim can be a bit of a struggle initially, but it is quite rewarding once you have a deeper undrestanding of the components involved. Although debuggers can often be overkill and dependending on the situation may actually lead one to be less productive they can play an important role in an engineer's workflow particularly when exploring an unknown code base.

