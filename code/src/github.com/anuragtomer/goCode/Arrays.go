package main

import "fmt"

func _arrayUseCase() {
	grade1 := 97
	grade2 := 98
	grade3 := 79
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)
	grades := [3]int{97, 98, 79}          // [size]type{values}
	gradesAnother := [...]int{97, 98, 79} // create an array that has just enough size to hold these values
	var students [3]string
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Second grades: %v\n", gradesAnother)
	fmt.Printf("Students: %v\n", students)
	students[0] = "Anurag"
	fmt.Printf("%v\n", students)
	fmt.Printf("%v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	var identityMatrixAnother [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("%v\n", identityMatrix)
	fmt.Printf("%v\n", identityMatrixAnother)
	var identityMatrixThird [3][3]int
	identityMatrixThird[0] = [3]int{1, 0, 0}
	identityMatrixThird[1] = [3]int{0, 1, 0}
	identityMatrixThird[2] = [3]int{0, 0, 1}
	fmt.Printf("%v\n", identityMatrixThird)

	a := [...]int{1, 2, 3}
	b := a // Create a copy, deep copy
	b[1] = 5
	fmt.Printf("%v, %v\n", a, b)

	c := &a // Reference. Edit the same copy of a
	c[2] = 9
	fmt.Printf("%v, %v, %v", a, b, c)

}
func arrayUsecase() {
	// _arrayUseCase()
	slice()
}

func slice() {
	// a := []int{1, 2, 3} // Note no size, no ... These are slices, not array.
	// fmt.Printf("%v, %v, %v\n", a, len(a), cap(a))
	// b := a // Point to same slice as being pointed to by 'a'
	// b[1] = 27
	// fmt.Printf("%v, %v", a, b)

	// c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Can also be [...]. Then it would be an array. But subsequent declarations are making a slice out of this.
	// d := c[:]
	// e := c[3:]
	// f := c[:6]
	// g := c[3:6]
	// fmt.Printf("%v, %v, %v, %v, %v\n", c, d, e, f, g)
	// c[5] = 42
	// fmt.Printf("%v, %v, %v, %v, %v\n", c, d, e, f, g)

	// h := make([]int, 3, 100)
	// fmt.Println(h)
	// fmt.Printf("Length: %v\n", len(h))
	// fmt.Printf("Capacity: %v\n", cap(h))
	// j := []int{}
	// j = append(j, 1)
	// fmt.Print(j)
	// fmt.Printf("Length: %v\n", len(j))
	// fmt.Printf("Capacity: %v\n", cap(j))
	// j = append(j, 2, 3, 4, 5, 6, 7) // 2, 3, 4, etc are all appended to first param (j)
	// fmt.Print(j)
	// fmt.Printf("Length: %v\n", len(j))
	// fmt.Printf("Capacity: %v\n", cap(j))
	// // j = append(j, []int{0, 1, 2})    // This does not work because j is a slice of integers, but []int{} is a slice itself. append cannot add a slice to a slice of integers.
	// j = append(j, []int{0, 1, 2}...) // ... expands the slice to individual elements and then it can be appended.

	// Stack operations on slice
	a := []int{1, 2, 3, 4, 5}
	b := a[1:] // Remove from beginning
	fmt.Println(a)
	fmt.Println(b)
	b = a[:len(a)-1] // Remove from end
	fmt.Println(a)
	fmt.Println(b)
	b = append(a[:2], a[3:]...)
	fmt.Println(a) // This is weird. append would change the underlying array.
	fmt.Println(b)
}
