package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID      int
	Name    string
	Address string
	Birth   time.Time
}

type Color struct{ Red, Green, Blue int }

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func main() {

	var employee Employee
	employee.Address = "aaa"
	fmt.Println(employee)
	var address *string = &employee.Address
	*address = "bbb"
	fmt.Println(address)
	fmt.Println(employee)

	employeePtr := &employee
	employeePtr.Address = "ccc" //equals to (*employeePtr).Address="ccc"
	fmt.Println(employee)

	color := Color{255, 255, 255}
	fmt.Println(color)
	color = Color{Red: 250, Green: 250}
	fmt.Println(color)

	circle := new(Circle)
	circle.X = 3
	circle.Y = 4
	fmt.Println(*circle)
}
