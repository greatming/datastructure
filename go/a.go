package main

func main() {
	println("=== begin ===")
	defer func() { // defer_0
		println("=== come in defer_0 ===")
	}()
	defer func() { // defer_1
		recover()
	}()
	defer func() { // defer_2
		panic("panic 2")
	}()
	panic("panic 1")
	println("=== end ===")
}
