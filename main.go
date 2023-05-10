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