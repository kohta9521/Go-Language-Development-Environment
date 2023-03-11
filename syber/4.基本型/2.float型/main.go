package main

import "fmt"

func main() {
	var f64 float64 = 2.4
	fmt.Println(f64)

	fl := 3.2
	fmt.Println(f64 + fl)
	fmt.Printf("%T, %T\n", f64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

}
