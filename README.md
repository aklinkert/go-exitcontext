# go-exitcontext

This package provides functions to create a context that listens for `SIG_` signals and cancels the context, used to abort programs by sending signals.


# Usage

If you already have a context and want to recycle that one:

```go
ctx := context.Background()
exitCtx := exitcontext.NewWithContext(ctx)
```

else just use the new method, which creates a `context.Background()` under the hood:

```go
exitCtx := exitcontext.New()
```

## License

    MIT License

    Copyright (c) 2019 Alexander Pinnecke
