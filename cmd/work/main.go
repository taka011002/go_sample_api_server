package main

import "fmt"

type user struct {
	name string
	human human
}

type human struct {
	name string
}

func (u user) create() {
	fmt.Println("created User")
}

func (u human) create() {
	fmt.Println("created Human")
}

func main()  {
	user := user{name: "taka"}

	fmt.Println(user)
	user.create()
	user.human.create()
}
