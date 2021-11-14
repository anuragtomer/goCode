package main

import (
	"fmt"
)

func constants() {
	const myConst int = 42 // Just as any other variable.
	const MyConst int = 99 // Global constant(starting with upper case). Accessible in other files in the same package
	fmt.Printf("%v\n", myConst)
	fmt.Printf("%v\n", MyConst)
	// const myConstCompileTime float64 = math.Sin(1.57) // Not allowed. math.Sin(1.57) will have result only at runtime, not at compile time.
	// Inferred constants
	const a = 42
	var b int16 = 27
	fmt.Printf("%v\n", a+b) // Works. Treats a as int16
	// Following doesn't work
	// const a1 int = 42
	// var b1 int16 = 27
	// fmt.Printf("%v\n", a1+b1) // Doesn't work. Mismatch between int and int16
	// enumeratedConstants()
	// writeOnly()
	FileSize()
}

// Enumerated Constants
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

// Enumerated Constants
const (
	a1 = iota // 0
	b1        // 1
	c1        // 2
)

func enumeratedConstants() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", c1)
}

const (
	_  = iota // Ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func writeOnly() {
	_ = 1 // Can dump to _, can't read it from.
}

func FileSize() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}

// 1:41:13
