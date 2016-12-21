# Exercise

We did the exercise of reverse-engineering a binary again, same as on
[2016-10-19](../../../2016/2016-10-19/reverse-engineering/README.md).

We wrote a bit more of a `main.c` file, but still incomplete. We realized it is
gonna take some time, and personally I was lacking a bit of feedback from the
process... but seems to be part of the game.

Before stopping, we wrote a "hello world" in Go, just to see how it looks like
from `gdb`.

We compared the 64 and 32-bit binaries:

```
GOOS=linux GOARCH=386 go build -o go-hello main.go
GOOS=linux GOARCH=amd64 go build -o go-hello main.go
```
