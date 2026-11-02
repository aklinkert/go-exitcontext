# go-exitcontext

[![test](https://github.com/aklinkert/go-exitcontext/actions/workflows/test.yml/badge.svg)](https://github.com/aklinkert/go-exitcontext/actions/workflows/test.yml)

A Go package that returns a `context.Context` which is cancelled when the process receives `SIGINT` or `SIGTERM`, to shut programs down gracefully.

## Installation

```sh
go get github.com/aklinkert/go-exitcontext
```

## Usage

If you already have a context and want to recycle that one:

```go
ctx := context.Background()
exitCtx := exitcontext.NewWithContext(ctx)
```

else just use `New`, which creates a `context.Background()` under the hood:

```go
exitCtx := exitcontext.New()

<-exitCtx.Done() // SIGINT or SIGTERM received
```

The signal handler is released as soon as the context is done, whether a signal arrived or the parent context was cancelled.

For new code, the standard library's [`signal.NotifyContext`](https://pkg.go.dev/os/signal#NotifyContext) offers the same with an explicit stop function.

## License

MIT, see [LICENSE](LICENSE).
