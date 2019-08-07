package main

type Object struct {
	S string
	I int
	B bool
}

func main() {
}

func Sum(i, i2 int) int {
	return i + i2
}

func genObject(s string, i int, b bool) Object {
	return Object{s, i, b}
}
