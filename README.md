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

* refactor `handler`s, need a few generic utils to reduce boilerplate
* add SQLlite repository demo

## Installing and Running

You will need GO installed. I highly recommend using [gobrew](<https://github.com/kevincobain2000/gobrew>)

Install using curl

``` zsh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/gobrew/master/git.io.sh | sh
# list remote versions available
gobrew ls-remote
# install the latest non-beta release (1.21.4 as of 2023-11-20)
gobrew install 1.21.4
```

You will also need MongoDB and Redis. On macOs I recommend the following:

``` zsh
# install MongoDB community edition ( for local testing )
brew tap mongodb/brew
brew update
brew install mongodb-community
brew services start mongodb-community
# install Redis
brew install redis
brew services start redis
```

Install dependencies

``` zsh
# install npm packages (tailwindcss, postcss, browser-sync, etc.)
# yes, HTMX is mostly Javascript-less, but some development packages are needed
npm install
```

Then run the following from a terminal window:

``` zsh
# run air, tailwind watcher concurrently
source ./dev.sh
```

The site should now be available at [<http://localhost:3000>](http://localhost:3000)

## Datastores

Both Redis and MonoDB-Community are in use. This shows how to use multiples data repositories and how to abstract the datastore mechanisms away from the web handlers.

A SQLLite example is planned.
