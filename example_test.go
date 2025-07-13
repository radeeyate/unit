package unit

import "fmt"

func ExampleFeetInMeters() {
	ft := 1 * Foot
	fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
	// Output: 1 feet is 0.30479999999999996 meters
}

type MyUnit int

func ExampleFromOwnUnit() {
	n := MyUnit(2)
	ft := Length(n) * Foot
	fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
	// Output: 2 feet is 0.6095999999999999 meters
}

func ExampleConverter() {
	v, err := NewConverter(1).From("F").To("mF")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	v, err = NewConverter(1).From("kF").To("F")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	v, err = NewConverter(1).From("mF").To("uF")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.0f\n", v)

	// Output: 1000
	// 1000
	// 1000
}