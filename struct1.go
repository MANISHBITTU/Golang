package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type detail struct {
	person
	empid int
}

func main() {
	p := person{fname: "Manish", lname: "Sharma"}
	fmt.Println(p.fname, p.lname)
	d := detail{person: person{fname: "Manish", lname: "Sharma"}, empid: 25}
	fmt.Println(d)

}
