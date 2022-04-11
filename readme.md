# go-rproxy
Remember [py-rproxy](https://github.com/zeroxoneafour/py-rproxy2)? Of course you don't. Anyway, this is a continuation of that, but written in Go.

## usage
Just go to ex. `localhost:8000/https://<some website>`

**important** - Make sure, if you go to a root website (ex. https://google.com/) that you put the backslash at the end as to avoid a flaw in my parsing code.

## state as of now
Works fully on normal HTML, but not JavaScript. You can browse non-JS sites like Amazon and Old Reddit with full functionality, and use Google to a limited extent. I'm going to mostly abandon this approach in favor of client-side rewriting tools (such as a server that simply injects a script) and use a script like [this](https://github.com/titaniumnetwork-dev/Ultraviolet).

However, as this does work, it's the most accomplished thing I've probably ever done/will do in the ways of making proxies.
