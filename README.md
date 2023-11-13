# GOponents

A technology demonstrator for

* [GO](https://go.dev) with auto-reloading via air
* [HTMX](https://htmx.org) with browser reloading via browser-sync
* [HyperScript](https://hyperscript.org) with built-in support for Server Sent Events
* [TailwindCSS](https://tailwindcss.com) with auto-reloading
* [daisyUI](https://daisyUI.com) nice, modern UI with themes, and more

One of the main principles of HTMX is HATEOAS (Hypermedia as the engine of application state).

Simply put, state is maintained on the server, not the client.

This is demonstrated here in several ways. Theme and (current) Module have their state maintained on the server. This allows the user to reload the page, withOUT losing state.

## TODOS

* refactor `template`s, with better `block`s and `define`s
* refactor `handler`s, need a few generic utils to reduce boilerplate
* move `routes` into the `module`, make the modules more transportable between projects
* add SQLlite repository demo
* add actual SSE example (perhaps Stock Prices?)

## Installing and Running

You will need GO installed. I highly recommend using [gobrew](<https://github.com/kevincobain2000/gobrew>)

Install using curl

``` zsh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/gobrew/master/git.io.sh | sh
```

Install dependencies

``` zsh
# install packages (tailwindcss, postcss, browser-sync, etc.)
npm install
```

Then run the following from a terminal window:

``` zsh
# run air, tailwind watcher and browser-sync concurrently
source ./dev.sh
```

## Datastores

Both Redis and MonoDB are in use. This shows how to use multiples data repositories and how to abstract the datastore mechanisms away from the web handlers.

Redis and MongoDB need to be installed, but that's not covered here.
