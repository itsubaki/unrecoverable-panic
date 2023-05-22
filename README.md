# unrecoverable-panic

```go
package main

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("unreachable: %v", err)
		}
	}()

	f()

	println("unreachable")
}

func f() {
	go panic("something went wrong")
	select {}
}
```

```shell
$ go run main.go 
panic: something went wrong

goroutine 17 [running]:
main.f.func1()
        /workspaces/unrecoverable-panic/main.go:16 +0x2a
created by main.f
        /workspaces/unrecoverable-panic/main.go:16 +0x4a
exit status 2
$
```

## Links

 * https://zenn.dev/matsuyoshi/articles/a492d339d21317
 * https://github.com/golang/sync/blob/a6666c150eb9762365996ad2f3a8c35b0818d7bb/singleflight/singleflight.go#L158
