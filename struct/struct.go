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

func (e *Employee) resetName() {
	e.Name = "unknown"
}

type Color struct{ Red, Green, Blue int }

type Point struct {
	X, Y int
}

func (p *Point) reset() {
	p.X = 0
	p.Y = 0
}

type Circle struct {
	Point
	Radius int
}

type Rocket struct {
}

func (r *Rocket) Launch() {
	fmt.Println("launch")
}

func main() {

	var employee Employee
	employee.Address = "aaa"
	fmt.Println(employee)
	var address *string = &employee.Address
	*address = "bbb"
	fmt.Println(address)
	fmt.Println(employee)
	resetEmployeeName := employee.resetName
	resetEmployeeName() //equals to employee.resetName()
	fmt.Println(employee)

	employeePtr := &employee
	employeePtr.Address = "ccc" //equals to (*employeePtr).Address="ccc"
	fmt.Println(employee)

	employeePtr = new(Employee) //return a ptr
	employeePtr.Address = "ddd"
	fmt.Println(*employeePtr)

	color := Color{255, 255, 255}
	fmt.Println(color)
	color = Color{Red: 250, Green: 250}
	fmt.Println(color)

	circle := new(Circle)
	circle.X = 3
	circle.Y = 4
	circle.reset() //继承Point的方法

	fmt.Println(*circle)

	(&Point{1, 2}).reset() //compile error if Point{1,2}.reset()

	rocket := new(Rocket)
	(*Rocket).Launch(rocket)
	time.AfterFunc(1*time.Second, rocket.Launch)
	//equals to below
	time.AfterFunc(1*time.Second, func() { rocket.Launch() })
}
