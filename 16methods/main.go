package main

import "fmt"

func main() {

	junaid := User{
		Name:   "Junaid",
		Email:  "junaid@example.com",
		Status: true,
		Age:    30,
	}
	fmt.Println(junaid)
	fmt.Printf("this person name are given below %+v\n",junaid.Name)
	junaid.GetStatus()
	junaid.setEmail("k@changed.com")
	fmt.Println("The changed email is",junaid.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("The status is",u.Status)
}
// func (u User) setEmail(email string) User{
// 	u.Email=email
// 	return u
// }
func (u *User) setEmail(email string){
	u.Email=email
}