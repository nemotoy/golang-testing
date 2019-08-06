package main

type Object struct {
	A string
	I int
	B bool
}

func main() {
}

func Sum(i, i2 int) int {
	return i + i2
}

func genObject(a string, i int, b bool) Object {
	return Object{a, i, b}
}
