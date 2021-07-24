package main

func main() {
	//constant 99999214748364899999999999999999999999999999999999 overflows int
	// var integer32 int = 99999214748364899999999999999999999999999999999999
	var integer32 int = 999992147483648
	println(integer32)

	// constant -10 overflows uint
	//var integer uint = -10
	var integer uint = 10
	println(integer)
}
