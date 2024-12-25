package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	jack := User{Name: "jack Sparrow", Age: 20, Email: "jacksaprrow@gmail.com", Status: true}

	jack.GetStatus()
	jack.SetEmail("jack@mail.com")
	fmt.Printf("Name %v and Email is %v\n", jack.Name, jack.GetEmail())

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (email *User) SetEmail(e string) {
	email.Email = e
}
func (e User) GetEmail() string {
	// fmt.Printf("The %v Email is %v\n", e.Name, e.Email)
	return e.Email
}
