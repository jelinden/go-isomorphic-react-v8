React.js server and client side rendering with Go
=====

This experiment is based on [selfjs](https://github.com/nmerouze/selfjs) and
[w8worker](https://github.com/ry/v8worker), a Javascript interpreter for Go.
The goal for the experiment was to explore if reactjs server side rendering could be 
done with Go and to use the same code in the browser.

As a http server [echo](https://labstack.github.io/echo/) is used.

In rss.go we are fetching a rss feed (scheduled in main.go). Scheduling also renders the results 
and saves them in a global variable to be used later.

You can test it here: [isomorphic2.uutispuro.fi](http://isomorphic2.uutispuro.fi/)

To run:
```bash
go build && ./go-isomorphic-react-v8
```

To benchmark serverside rendering:
```bash
go test -bench=.
```

The result on my MacBook Air (1.4 GHz i5):
```bash
PASS
BenchmarkRender1               1        277493982 ns/op
ok      github.com/jelinden/go-isomorphic-react-v8 2.747s
```