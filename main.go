package main

func main() {

}

// FindXYZ - X, Y, 5, 9, 15, 23, Z  - Please create a new function for finding X, Y, Z value
// X, Y, 5, 9, 15, 23, Z
//  \/ \/ \/ \/  \/  \/
//   0  2  4  6  8   10
// So, Z = 23 + 10 = 33
// Y = 5 - 2 = 3
// X = Y - 0 = 3 - 0 = 3
func FindXYZ() (int, int, int) {
	x := findX()
	y := findY()
	z := findZ()
	return x, y, z
}

func findX() int {
	return 3
}

func findY() int {
	return 3
}

func findZ() int {
	return 33
}

// FindB - A = 21
func FindB() int {
	return 23 - 21
}

// FindC - A = 21
func FindC() int {
	return -21 - 21
}
