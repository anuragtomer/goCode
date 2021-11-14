package main

import (
	"fmt"
	"strconv"
)

func variableDeclarations() {
	var j int // Don't know what its value will be to begin with.
	j = 18
	fmt.Println(j)
	var k int = 19
	fmt.Println(k)
	i := 42
	fmt.Println(i)
	fmt.Println("Hello everyone")
	fmt.Println("Testing")
	n := 99
	fmt.Printf("%v, %T\n", n, n)
	m := "Another way"
	fmt.Printf("%v, %T\n", m, m)
	// Block level declaration
	// You can group variables logically in single block. But its just easier to write, nothing fancy for compiler
	var (
		actorName     string = "Anurag"
		companion     string = "Surabhi"
		enrollmentNum int    = 1
		class         int    = 5
	)
	fmt.Printf("%s %s %d %d", actorName, companion, enrollmentNum, class)
}

var i int = 42

func scopeResolution() {
	var i int = 99 // Shadowing global i
	fmt.Println(i) // Will print 99
}

func mustUseVariable() {
	i := 50
	j := 100
	fmt.Println(i)
	// Error that j is not used. Must use all declared variables
	fmt.Println(j)
}

func variableNamingConvention() {
	/*
		- Lower cased variables DECLARED GLOBALLY are scoped to this current package (main in our case)
		- Upper cased variables DECLARED GLOBALLY are scoped to outside world (global visibility).
		- Lower cased variables declared locally in a scope are visible only in that scope.
		- No private scope
	*/
}

func convertOneVariableTypeToAnother() {
	i := 42
	fmt.Printf("%v %T\n", i, i)
	var j float32
	j = float32(i) // Conversion function float32()
	fmt.Printf("%v %T\n", j, j)
	j = 42.5
	fmt.Printf("%v %T\n", j, j)
	i = int(j)
	fmt.Printf("%v %T\n", i, i)
	var k string
	k = string(i) // This doesn't really work. because string is a stream of bytes in golang. string looks for unicode corresponding for that num and change that to string. So wrong.
	fmt.Println(k)
	k = strconv.Itoa(i) // This works.
	fmt.Println(k)
}

func differentDataTypes() {
	var k bool = true
	fmt.Printf("%v, %T\n", k, k)
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	// All declared variable have default zero values [For bool, its false]

	/*
		Different types:
		bool

		int (Defaults to int32 but system dependent)
		int8
		int16
		int32
		int64
		uint
		uint8 [ or byte]
		uint16
		uint32
		NOTE: We are not allowed to do arithmetic between 2 different types.
		i.e. int + int8 is not allowed, even they both are ints

		float32
		float64 (Default)

		complex64 (Basically taking float32 + ifloat32, which makes it complex64)
		complex128
	*/
	/*
		Bitwise operators (Only works with integer data types, not floats):
		& - and
		| - or
		^ - exor (Either set, but not both)
		&^ - exnor (Both have same bits set/unset)
		<< bit shift left
		>> bit shift right
	*/
	f := 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Println(f)
	com := 1 + 2i
	fmt.Println(com)
	fmt.Printf("%v %v\n", real(com), imag(com))
	// Conversion from real to complex number
	var com2 complex64 = complex(5, 12)
	fmt.Println(com2)

	// String data type
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v\n", s[0])         // This prints unicode corresponding to 't'
	fmt.Printf("%v\n", string(s[0])) // This prints 't'
	// Strings are basically stream of unicode bytes. So when we tried to access s[i], it prints that unicode.
	// But because s[i] is just a byte, I can print it using string(s[i]).
	// Strings are immutable
	// s[0] = "a" // This is wrong. Can't manipulate. Also, "a" is stream of byte, can't assign it to s[i].
	// Arithmetic operations:
	s2 := "This is another"
	// Concatenation
	fmt.Println(s + " " + s2)
	// Byte slicing
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b) // This prints all ascii for all characters in s.

	// Rune -> alias for int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r) // Prints 97
	var r2 rune = 'a'
	fmt.Printf("%v, %T\n", r2, r2) // Same as above
	// To read rune from a reader, we need to call ReadRune(). This reads each rune one by one, sees what its size should be, and then returns back this character to the user.
}

func datatypes() {
	variableDeclarations()
	scopeResolution()
	mustUseVariable()
	variableNamingConvention()
	convertOneVariableTypeToAnother()
	differentDataTypes()
}
