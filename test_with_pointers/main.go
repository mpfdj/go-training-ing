package main

import "fmt"

type User struct {
	username string
}

func (u *User) setUsername(new string) {
	//func (u User) setUsername(new string) {
	u.username = new
}

func main() {
	a := User{"first"}
	b := User{"second"}
	c := User{"three"}

	a.setUsername("test")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// automatically use a pointer because it's specified on setUsername
	c.setUsername("hello")
	fmt.Println(c)
}
